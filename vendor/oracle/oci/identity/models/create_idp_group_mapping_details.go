package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// CreateIdpGroupMappingDetails create idp group mapping details
// swagger:model CreateIdpGroupMappingDetails
type CreateIdpGroupMappingDetails struct {

	// The OCID of the IAM Service [group](#/en/identity/20160918/Group/)
	// you want to map to the IdP group.
	//
	// Required: true
	GroupID *string `json:"groupId"`

	// The name of the IdP group you want to map.
	// Required: true
	IdpGroupName *string `json:"idpGroupName"`
}

// Validate validates this create idp group mapping details
func (m *CreateIdpGroupMappingDetails) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateGroupID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateIdpGroupName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateIdpGroupMappingDetails) validateGroupID(formats strfmt.Registry) error {

	if err := validate.Required("groupId", "body", m.GroupID); err != nil {
		return err
	}

	return nil
}

func (m *CreateIdpGroupMappingDetails) validateIdpGroupName(formats strfmt.Registry) error {

	if err := validate.Required("idpGroupName", "body", m.IdpGroupName); err != nil {
		return err
	}

	return nil
}
