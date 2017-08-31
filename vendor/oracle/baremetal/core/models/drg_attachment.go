package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// DrgAttachment A link between a DRG and VCN. For more information, see
// [Overview of the Networking Service](/Content/Network/Concepts/overview.htm).
//
// swagger:model DrgAttachment
type DrgAttachment struct {

	// The OCID of the compartment containing the DRG attachment.
	// Required: true
	// Max Length: 255
	// Min Length: 1
	CompartmentID *string `json:"compartmentId"`

	// A user-friendly name. Does not have to be unique, and it's changeable.
	// Max Length: 255
	// Min Length: 1
	DisplayName string `json:"displayName,omitempty"`

	// The OCID of the DRG.
	// Required: true
	// Max Length: 255
	// Min Length: 1
	DrgID *string `json:"drgId"`

	// The DRG attachment's Oracle ID (OCID).
	// Required: true
	// Max Length: 255
	// Min Length: 1
	ID *string `json:"id"`

	// The DRG attachment's current state.
	// Required: true
	LifecycleState *string `json:"lifecycleState"`

	// The date and time the DRG attachment was created, in the format defined by RFC3339.
	//
	// Example: `2016-08-25T21:10:29.600Z`
	//
	TimeCreated strfmt.DateTime `json:"timeCreated,omitempty"`

	// The OCID of the VCN.
	// Required: true
	// Max Length: 255
	// Min Length: 1
	VcnID *string `json:"vcnId"`
}

// Validate validates this drg attachment
func (m *DrgAttachment) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCompartmentID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateDisplayName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateDrgID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateLifecycleState(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateVcnID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DrgAttachment) validateCompartmentID(formats strfmt.Registry) error {

	if err := validate.Required("compartmentId", "body", m.CompartmentID); err != nil {
		return err
	}

	if err := validate.MinLength("compartmentId", "body", string(*m.CompartmentID), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("compartmentId", "body", string(*m.CompartmentID), 255); err != nil {
		return err
	}

	return nil
}

func (m *DrgAttachment) validateDisplayName(formats strfmt.Registry) error {

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

func (m *DrgAttachment) validateDrgID(formats strfmt.Registry) error {

	if err := validate.Required("drgId", "body", m.DrgID); err != nil {
		return err
	}

	if err := validate.MinLength("drgId", "body", string(*m.DrgID), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("drgId", "body", string(*m.DrgID), 255); err != nil {
		return err
	}

	return nil
}

func (m *DrgAttachment) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	if err := validate.MinLength("id", "body", string(*m.ID), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("id", "body", string(*m.ID), 255); err != nil {
		return err
	}

	return nil
}

var drgAttachmentTypeLifecycleStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ATTACHING","ATTACHED","DETACHING","DETACHED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		drgAttachmentTypeLifecycleStatePropEnum = append(drgAttachmentTypeLifecycleStatePropEnum, v)
	}
}

const (
	// DrgAttachmentLifecycleStateATTACHING captures enum value "ATTACHING"
	DrgAttachmentLifecycleStateATTACHING string = "ATTACHING"
	// DrgAttachmentLifecycleStateATTACHED captures enum value "ATTACHED"
	DrgAttachmentLifecycleStateATTACHED string = "ATTACHED"
	// DrgAttachmentLifecycleStateDETACHING captures enum value "DETACHING"
	DrgAttachmentLifecycleStateDETACHING string = "DETACHING"
	// DrgAttachmentLifecycleStateDETACHED captures enum value "DETACHED"
	DrgAttachmentLifecycleStateDETACHED string = "DETACHED"
)

// prop value enum
func (m *DrgAttachment) validateLifecycleStateEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, drgAttachmentTypeLifecycleStatePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *DrgAttachment) validateLifecycleState(formats strfmt.Registry) error {

	if err := validate.Required("lifecycleState", "body", m.LifecycleState); err != nil {
		return err
	}

	// value enum
	if err := m.validateLifecycleStateEnum("lifecycleState", "body", *m.LifecycleState); err != nil {
		return err
	}

	return nil
}

func (m *DrgAttachment) validateVcnID(formats strfmt.Registry) error {

	if err := validate.Required("vcnId", "body", m.VcnID); err != nil {
		return err
	}

	if err := validate.MinLength("vcnId", "body", string(*m.VcnID), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("vcnId", "body", string(*m.VcnID), 255); err != nil {
		return err
	}

	return nil
}
