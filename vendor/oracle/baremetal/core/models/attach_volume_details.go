package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/validate"
)

// AttachVolumeDetails attach volume details
// swagger:discriminator AttachVolumeDetails type
type AttachVolumeDetails interface {
	runtime.Validatable

	// A user-friendly name. Does not have to be unique, and it cannot be changed.
	//
	// Max Length: 255
	// Min Length: 1
	DisplayName() string
	SetDisplayName(string)

	// The OCID of the instance.
	// Required: true
	// Max Length: 255
	// Min Length: 1
	InstanceID() *string
	SetInstanceID(*string)

	// The type of volume. The only supported value is "iscsi".
	// Required: true
	// Max Length: 255
	// Min Length: 1
	Type() string
	SetType(string)

	// The OCID of the volume.
	// Required: true
	// Max Length: 255
	// Min Length: 1
	VolumeID() *string
	SetVolumeID(*string)
}

// UnmarshalAttachVolumeDetailsSlice unmarshals polymorphic slices of AttachVolumeDetails
func UnmarshalAttachVolumeDetailsSlice(reader io.Reader, consumer runtime.Consumer) ([]AttachVolumeDetails, error) {
	var elements []json.RawMessage
	if err := consumer.Consume(reader, &elements); err != nil {
		return nil, err
	}

	var result []AttachVolumeDetails
	for _, element := range elements {
		obj, err := unmarshalAttachVolumeDetails(element, consumer)
		if err != nil {
			return nil, err
		}
		result = append(result, obj)
	}
	return result, nil
}

// UnmarshalAttachVolumeDetails unmarshals polymorphic AttachVolumeDetails
func UnmarshalAttachVolumeDetails(reader io.Reader, consumer runtime.Consumer) (AttachVolumeDetails, error) {
	// we need to read this twice, so first into a buffer
	data, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	return unmarshalAttachVolumeDetails(data, consumer)
}

func unmarshalAttachVolumeDetails(data []byte, consumer runtime.Consumer) (AttachVolumeDetails, error) {
	buf := bytes.NewBuffer(data)
	buf2 := bytes.NewBuffer(data)

	// the first time this is read is to fetch the value of the type property.
	var getType struct {
		Type string `json:"type"`
	}
	if err := consumer.Consume(buf, &getType); err != nil {
		return nil, err
	}

	if err := validate.RequiredString("type", "body", getType.Type); err != nil {
		return nil, err
	}

	// The value of type is used to determine which type to create and unmarshal the data into
	switch getType.Type {
	case "iscsi":
		var result AttachIScsiVolumeDetails
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil

	}
	return nil, errors.New(422, "invalid type value: %q", getType.Type)

}
