package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// UpdateSecurityListDetails update security list details
// swagger:model UpdateSecurityListDetails
type UpdateSecurityListDetails struct {

	// A user-friendly name. Does not have to be unique, and it's changeable.
	// Max Length: 255
	// Min Length: 1
	DisplayName string `json:"displayName,omitempty"`

	// Rules for allowing egress IP packets.
	EgressSecurityRules []*EgressSecurityRule `json:"egressSecurityRules"`

	// Rules for allowing ingress IP packets.
	IngressSecurityRules []*IngressSecurityRule `json:"ingressSecurityRules"`
}

// Validate validates this update security list details
func (m *UpdateSecurityListDetails) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDisplayName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateEgressSecurityRules(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateIngressSecurityRules(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UpdateSecurityListDetails) validateDisplayName(formats strfmt.Registry) error {

	if swag.IsZero(m.DisplayName) { // not required
		return nil
	}

	if err := validate.MinLength("displayName", "body", string(m.DisplayName), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("displayName", "body", string(m.DisplayName), 255); err != nil {
		return err
	}

	return nil
}

func (m *UpdateSecurityListDetails) validateEgressSecurityRules(formats strfmt.Registry) error {

	if swag.IsZero(m.EgressSecurityRules) { // not required
		return nil
	}

	for i := 0; i < len(m.EgressSecurityRules); i++ {

		if swag.IsZero(m.EgressSecurityRules[i]) { // not required
			continue
		}

		if m.EgressSecurityRules[i] != nil {

			if err := m.EgressSecurityRules[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("egressSecurityRules" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *UpdateSecurityListDetails) validateIngressSecurityRules(formats strfmt.Registry) error {

	if swag.IsZero(m.IngressSecurityRules) { // not required
		return nil
	}

	for i := 0; i < len(m.IngressSecurityRules); i++ {

		if swag.IsZero(m.IngressSecurityRules[i]) { // not required
			continue
		}

		if m.IngressSecurityRules[i] != nil {

			if err := m.IngressSecurityRules[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ingressSecurityRules" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}
