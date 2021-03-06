package config

import (
	"encoding/json"
	"errors"
	fakesys "github.com/cloudfoundry/bosh-utils/system/fakes"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/oracle/bosh-oracle-cpi/registry"
	"os"
	"path/filepath"
	"testing"
)

var apiKeyFile = fakeAPIKeyPath()
var validOCIProperties = OCIProperties{
	Tenancy:       "Fake-tenancy-ocid",
	User:          "user-ocid",
	CompartmentID: "fake-compartment-id",
	APIKeyFile:    apiKeyFile,
	CpiKeyFile:    "/path/to/generated/configdir/cpikey.pem",
	CpiUser:       "fake-provisioned-user",
	AuthorizedKeys: AuthorizedKeys{
		Cpi:  "ssh-rsa pub key generated by create-env",
		User: "ssh-rsa user key",
	},
	Fingerprint: "fake-fingerprint",
}

var forwardSSHTunnel = SSHTunnel{
	User: "opc", LocalPort: 36868, Duration: "2m",
}
var validOCIPropertiesWithSSHTunnel = OCIProperties{
	Tenancy:       "Fake-tenancy-ocid",
	User:          "user-ocid",
	CompartmentID: "fake-compartment-id",
	APIKeyFile:    apiKeyFile,
	CpiKeyFile:    "/path/to/generated/configdir/cpikey.pem",
	CpiUser:       "fake-provisioned-user",
	AuthorizedKeys: AuthorizedKeys{
		Cpi: "ssh-rsa fake key",
	},
	Fingerprint: "fake-fingerprint",
	SSHTunnel:   forwardSSHTunnel,
}

var validAgentOptions = registry.AgentOptions{
	Mbus: "fake-mbus",
	Ntp:  []string{},
	Blobstore: registry.BlobstoreOptions{
		Provider: "fake-blobstore-type",
	},
}

var validRegistryOptions = registry.ClientOptions{
	Protocol: "http",
	Host:     "fake-host",
	Port:     5555,
	Username: "fake-username",
	Password: "fake-password",
}

var validConfig = Config{
	Cloud: Cloud{
		Plugin: "oracle",
		Properties: CPIProperties{
			OCI:      validOCIProperties,
			Agent:    validAgentOptions,
			Registry: validRegistryOptions,
		},
	},
}

var validConfigWithSSH = Config{
	Cloud: Cloud{
		Plugin: "oracle",
		Properties: CPIProperties{
			OCI:      validOCIPropertiesWithSSHTunnel,
			Agent:    validAgentOptions,
			Registry: validRegistryOptions,
		},
	},
}

var _ = Describe("NewConfigFromPath", func() {
	var (
		fs *fakesys.FakeFileSystem
	)

	BeforeEach(func() {
		fs = fakesys.NewFakeFileSystem()
	})

	It("returns error if config is empty", func() {
		_, err := NewConfigFromPath("", fs)
		Expect(err).To(HaveOccurred())
		Expect(err.Error()).To(ContainSubstring("Must provide a config file"))
	})

	It("returns error if config is not valid", func() {
		err := fs.WriteFileString("/config.json", "{}")
		Expect(err).ToNot(HaveOccurred())

		_, err = NewConfigFromPath("/config.json", fs)
		Expect(err).To(HaveOccurred())
		Expect(err.Error()).To(ContainSubstring("Validating config"))
	})

	It("returns error if file contains invalid json", func() {
		err := fs.WriteFileString("/config.json", "-")
		Expect(err).ToNot(HaveOccurred())

		_, err = NewConfigFromPath("/config.json", fs)
		Expect(err).To(HaveOccurred())
		Expect(err.Error()).To(ContainSubstring("Unmarshalling config"))
	})

	It("returns error if file cannot be read", func() {
		err := fs.WriteFileString("/config.json", "{}")
		Expect(err).ToNot(HaveOccurred())

		fs.ReadFileError = errors.New("fake-read-err")

		_, err = NewConfigFromPath("/config.json", fs)
		Expect(err).To(HaveOccurred())
		Expect(err.Error()).To(ContainSubstring("fake-read-err"))
	})
})

var _ = Describe("Validate", func() {
	var (
		config Config
	)

	Context("with valid cloud config", func() {
		BeforeEach(func() {
			config = validConfig
		})

		It("does not return an error", func() {
			err := config.Validate()
			Expect(err).ToNot(HaveOccurred())
		})
		Context(" with missing OCIProperties ", func() {
			BeforeEach(func() {
				config.Cloud.Properties.OCI = OCIProperties{}
			})
			It("returns invalid oci configuration error", func() {

				err := config.Validate()
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(ContainSubstring(errorMsg(invalidOCIConfiguration).String()))
			})
		})
		Context(" with missing agent options", func() {
			BeforeEach(func() {
				config.Cloud.Properties.Agent = registry.AgentOptions{}
			})

			It("returns invalid agent configuration error", func() {

				err := config.Validate()
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(ContainSubstring(errorMsg(invalidAgentConfiguration).String()))
			})
		})
		Context(" with missing client registry configuration", func() {
			BeforeEach(func() {
				config.Cloud.Properties.Registry = registry.ClientOptions{}
			})

			It("returns invalid registry client configuration error", func() {

				err := config.Validate()
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(ContainSubstring(errorMsg(invalidRegistryClientConfiguration).String()))
			})

		})
	})
})

var _ = Describe("Config", func() {
	var (
		config Config
	)

	Describe("Validate", func() {
		BeforeEach(func() {
			config = validConfigWithSSH
		})

		It("does not return error if all sections are valid", func() {
			err := config.Validate()
			Expect(err).ToNot(HaveOccurred())
		})
	})
})

var _ = Describe("DefaultOCIProperties", func() {
	var (
		ociProperties OCIProperties
	)

	Describe("DefaultSSHConfig", func() {
		BeforeEach(func() {
			ociProperties = validOCIProperties
		})

		It("checks default configuration uses private IP for ssh", func() {
			Expect(ociProperties.CpiSSHConfig().usePublicIP).To(BeFalse())
		})
	})
})

func TestConfig_Validate(t *testing.T) {
	b, err := json.MarshalIndent(validConfig, "", "")
	if err != nil {
		t.FailNow()
	}
	t.Logf("Marshalled struct %s\n", string(b))
}

func TestConfigWithSSHTunnel_Validate(t *testing.T) {
	b, err := json.MarshalIndent(validConfigWithSSH, "", "")
	if err != nil {
		t.FailNow()
	}
	t.Logf("Marshalled struct %s\n", string(b))
}

func assetsDir() string {
	dir, _ := os.Getwd()
	return filepath.Join(dir, "test/assets")
}

func fakeAPIKeyPath() string {
	return filepath.Join(assetsDir(), "fake_api_key.pem")
}
