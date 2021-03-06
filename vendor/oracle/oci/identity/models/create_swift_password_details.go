package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// CreateSwiftPasswordDetails create swift password details
// swagger:model CreateSwiftPasswordDetails
type CreateSwiftPasswordDetails struct {

	// The description you assign to the Swift password during creation. Does not have to be unique, and it's changeable.
	//
	// Required: true
	// Max Length: 400
	// Min Length: 1
	Description *string `json:"description"`
}

// Validate validates this create swift password details
func (m *CreateSwiftPasswordDetails) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDescription(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateSwiftPasswordDetails) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("description", "body", m.Description); err != nil {
		return err
	}

	if err := validate.MinLength("description", "body", string(*m.Description), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("description", "body", string(*m.Description), 400); err != nil {
		return err
	}

	return nil
}
