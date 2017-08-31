package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"bytes"
	"encoding/json"
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// DhcpSearchDomainOption DHCP option for specifying a search domain name for DNS queries. For more information, see
// [DNS in Your Virtual Cloud Network](/Content/Network/Concepts/dns.htm).
//
// swagger:model DhcpSearchDomainOption
type DhcpSearchDomainOption struct {

	// A single search domain name according to [RFC 952](https://tools.ietf.org/html/rfc952)
	// and [RFC 1123](https://tools.ietf.org/html/rfc1123). During a DNS query,
	// the OS will append this search domain name to the value being queried.
	//
	// If you set [DhcpDnsOption](#/en/iaas/20160918/DhcpDnsOption/) to `VcnLocalPlusInternet`,
	// and you assign a DNS label to the VCN during creation, the search domain name in the
	// VCN's default set of DHCP options is automatically set to the VCN domain
	// (e.g., `vcn1.oraclevcn.com`).
	//
	// If you don't want to use a search domain name, omit this option from the
	// set of DHCP options. Do not include this option with an empty list
	// of search domain names, or with an empty string as the value for any search
	// domain name.
	//
	// Required: true
	// Max Items: 1
	SearchDomainNames []string `json:"searchDomainNames"`
}

func (m *DhcpSearchDomainOption) Type() string {
	return "DhcpSearchDomainOption"
}
func (m *DhcpSearchDomainOption) SetType(val string) {

}

// UnmarshalJSON unmarshals this polymorphic type from a JSON structure
func (m *DhcpSearchDomainOption) UnmarshalJSON(raw []byte) error {
	var data struct {
		Type string `json:"type"`

		// A single search domain name according to [RFC 952](https://tools.ietf.org/html/rfc952)
		// and [RFC 1123](https://tools.ietf.org/html/rfc1123). During a DNS query,
		// the OS will append this search domain name to the value being queried.
		//
		// If you set [DhcpDnsOption](#/en/iaas/20160918/DhcpDnsOption/) to `VcnLocalPlusInternet`,
		// and you assign a DNS label to the VCN during creation, the search domain name in the
		// VCN's default set of DHCP options is automatically set to the VCN domain
		// (e.g., `vcn1.oraclevcn.com`).
		//
		// If you don't want to use a search domain name, omit this option from the
		// set of DHCP options. Do not include this option with an empty list
		// of search domain names, or with an empty string as the value for any search
		// domain name.
		//
		// Required: true
		// Max Items: 1
		SearchDomainNames []string `json:"searchDomainNames"`
	}

	buf := bytes.NewBuffer(raw)
	dec := json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&data); err != nil {
		return err
	}

	m.SearchDomainNames = data.SearchDomainNames

	return nil
}

// MarshalJSON marshals this polymorphic type to a JSON structure
func (m DhcpSearchDomainOption) MarshalJSON() ([]byte, error) {
	var data struct {
		Type string `json:"type"`

		// A single search domain name according to [RFC 952](https://tools.ietf.org/html/rfc952)
		// and [RFC 1123](https://tools.ietf.org/html/rfc1123). During a DNS query,
		// the OS will append this search domain name to the value being queried.
		//
		// If you set [DhcpDnsOption](#/en/iaas/20160918/DhcpDnsOption/) to `VcnLocalPlusInternet`,
		// and you assign a DNS label to the VCN during creation, the search domain name in the
		// VCN's default set of DHCP options is automatically set to the VCN domain
		// (e.g., `vcn1.oraclevcn.com`).
		//
		// If you don't want to use a search domain name, omit this option from the
		// set of DHCP options. Do not include this option with an empty list
		// of search domain names, or with an empty string as the value for any search
		// domain name.
		//
		// Required: true
		// Max Items: 1
		SearchDomainNames []string `json:"searchDomainNames"`
	}

	data.SearchDomainNames = m.SearchDomainNames
	data.Type = "DhcpSearchDomainOption"
	return json.Marshal(data)
}

// Validate validates this dhcp search domain option
func (m *DhcpSearchDomainOption) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSearchDomainNames(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DhcpSearchDomainOption) validateSearchDomainNames(formats strfmt.Registry) error {

	if err := validate.Required("searchDomainNames", "body", m.SearchDomainNames); err != nil {
		return err
	}

	iSearchDomainNamesSize := int64(len(m.SearchDomainNames))

	if err := validate.MaxItems("searchDomainNames", "body", iSearchDomainNamesSize, 1); err != nil {
		return err
	}

	for i := 0; i < len(m.SearchDomainNames); i++ {

		if err := validate.MinLength("searchDomainNames"+"."+strconv.Itoa(i), "body", string(m.SearchDomainNames[i]), 1); err != nil {
			return err
		}

		if err := validate.MaxLength("searchDomainNames"+"."+strconv.Itoa(i), "body", string(m.SearchDomainNames[i]), 251); err != nil {
			return err
		}

	}

	return nil
}
