package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"bytes"
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// IScsiVolumeAttachment An ISCSI volume attachment.
// swagger:model IScsiVolumeAttachment
type IScsiVolumeAttachment struct {
	availabilityDomainField *string

	compartmentIdField *string

	displayNameField string

	idField *string

	instanceIdField *string

	lifecycleStateField *string

	timeCreatedField *strfmt.DateTime

	volumeIdField *string

	// The Challenge-Handshake-Authentication-Protocol (CHAP) secret valid for the associated CHAP user name.
	// (Also called the "CHAP password".)
	//
	// Example: `d6866c0d-298b-48ba-95af-309b4faux45e`
	//
	ChapSecret string `json:"chapSecret,omitempty"`

	// The volume's system-generated Challenge-Handshake-Authentication-Protocol (CHAP) user name.
	//
	// Example: `ocid1.volume.oc1.phx.abyhqljrgvttnlx73nmrwfaux7kcvzfs3s66izvxf2h4lgvyndsdsnoiwr5q`
	//
	ChapUsername string `json:"chapUsername,omitempty"`

	// The volume's iSCSI IP address.
	//
	// Example: `169.254.0.2`
	//
	// Required: true
	// Max Length: 15
	// Min Length: 7
	IPV4 *string `json:"ipv4"`

	// The target volume's iSCSI Qualified Name in the format defined by RFC 3720.
	//
	// Example: `iqn.2015-12.us.oracle.com:456b0391-17b8-4122-bbf1-f85fc0bb97d9`
	//
	// Required: true
	// Min Length: 1
	Iqn *string `json:"iqn"`

	// The volume's iSCSI port.
	//
	// Example: `3260`
	//
	// Required: true
	// Maximum: 65535
	// Minimum: 1
	Port *int64 `json:"port"`
}

func (m *IScsiVolumeAttachment) AttachmentType() string {
	return DiscriminatorTypeValues["IScsiVolumeAttachment"]

}
func (m *IScsiVolumeAttachment) SetAttachmentType(val string) {

}

func (m *IScsiVolumeAttachment) AvailabilityDomain() *string {
	return m.availabilityDomainField
}
func (m *IScsiVolumeAttachment) SetAvailabilityDomain(val *string) {
	m.availabilityDomainField = val
}

func (m *IScsiVolumeAttachment) CompartmentID() *string {
	return m.compartmentIdField
}
func (m *IScsiVolumeAttachment) SetCompartmentID(val *string) {
	m.compartmentIdField = val
}

func (m *IScsiVolumeAttachment) DisplayName() string {
	return m.displayNameField
}
func (m *IScsiVolumeAttachment) SetDisplayName(val string) {
	m.displayNameField = val
}

func (m *IScsiVolumeAttachment) ID() *string {
	return m.idField
}
func (m *IScsiVolumeAttachment) SetID(val *string) {
	m.idField = val
}

func (m *IScsiVolumeAttachment) InstanceID() *string {
	return m.instanceIdField
}
func (m *IScsiVolumeAttachment) SetInstanceID(val *string) {
	m.instanceIdField = val
}

func (m *IScsiVolumeAttachment) LifecycleState() *string {
	return m.lifecycleStateField
}
func (m *IScsiVolumeAttachment) SetLifecycleState(val *string) {
	m.lifecycleStateField = val
}

func (m *IScsiVolumeAttachment) TimeCreated() *strfmt.DateTime {
	return m.timeCreatedField
}
func (m *IScsiVolumeAttachment) SetTimeCreated(val *strfmt.DateTime) {
	m.timeCreatedField = val
}

func (m *IScsiVolumeAttachment) VolumeID() *string {
	return m.volumeIdField
}
func (m *IScsiVolumeAttachment) SetVolumeID(val *string) {
	m.volumeIdField = val
}

// UnmarshalJSON unmarshals this polymorphic type from a JSON structure
func (m *IScsiVolumeAttachment) UnmarshalJSON(raw []byte) error {
	var data struct {
		AttachmentType string `json:"attachmentType"`

		AvailabilityDomain *string `json:"availabilityDomain"`

		CompartmentID *string `json:"compartmentId"`

		DisplayName string `json:"displayName,omitempty"`

		ID *string `json:"id"`

		InstanceID *string `json:"instanceId"`

		LifecycleState *string `json:"lifecycleState"`

		TimeCreated *strfmt.DateTime `json:"timeCreated"`

		VolumeID *string `json:"volumeId"`

		// The Challenge-Handshake-Authentication-Protocol (CHAP) secret valid for the associated CHAP user name.
		// (Also called the "CHAP password".)
		//
		// Example: `d6866c0d-298b-48ba-95af-309b4faux45e`
		//
		ChapSecret string `json:"chapSecret,omitempty"`

		// The volume's system-generated Challenge-Handshake-Authentication-Protocol (CHAP) user name.
		//
		// Example: `ocid1.volume.oc1.phx.abyhqljrgvttnlx73nmrwfaux7kcvzfs3s66izvxf2h4lgvyndsdsnoiwr5q`
		//
		ChapUsername string `json:"chapUsername,omitempty"`

		// The volume's iSCSI IP address.
		//
		// Example: `169.254.0.2`
		//
		// Required: true
		// Max Length: 15
		// Min Length: 7
		IPV4 *string `json:"ipv4"`

		// The target volume's iSCSI Qualified Name in the format defined by RFC 3720.
		//
		// Example: `iqn.2015-12.us.oracle.com:456b0391-17b8-4122-bbf1-f85fc0bb97d9`
		//
		// Required: true
		// Min Length: 1
		Iqn *string `json:"iqn"`

		// The volume's iSCSI port.
		//
		// Example: `3260`
		//
		// Required: true
		// Maximum: 65535
		// Minimum: 1
		Port *int64 `json:"port"`
	}

	buf := bytes.NewBuffer(raw)
	dec := json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&data); err != nil {
		return err
	}

	m.availabilityDomainField = data.AvailabilityDomain
	m.compartmentIdField = data.CompartmentID
	m.displayNameField = data.DisplayName
	m.idField = data.ID
	m.instanceIdField = data.InstanceID
	m.lifecycleStateField = data.LifecycleState
	m.timeCreatedField = data.TimeCreated
	m.volumeIdField = data.VolumeID
	m.ChapSecret = data.ChapSecret
	m.ChapUsername = data.ChapUsername
	m.IPV4 = data.IPV4
	m.Iqn = data.Iqn
	m.Port = data.Port

	return nil
}

// MarshalJSON marshals this polymorphic type to a JSON structure
func (m IScsiVolumeAttachment) MarshalJSON() ([]byte, error) {
	var data struct {
		AttachmentType string `json:"attachmentType"`

		AvailabilityDomain *string `json:"availabilityDomain"`

		CompartmentID *string `json:"compartmentId"`

		DisplayName string `json:"displayName,omitempty"`

		ID *string `json:"id"`

		InstanceID *string `json:"instanceId"`

		LifecycleState *string `json:"lifecycleState"`

		TimeCreated *strfmt.DateTime `json:"timeCreated"`

		VolumeID *string `json:"volumeId"`

		// The Challenge-Handshake-Authentication-Protocol (CHAP) secret valid for the associated CHAP user name.
		// (Also called the "CHAP password".)
		//
		// Example: `d6866c0d-298b-48ba-95af-309b4faux45e`
		//
		ChapSecret string `json:"chapSecret,omitempty"`

		// The volume's system-generated Challenge-Handshake-Authentication-Protocol (CHAP) user name.
		//
		// Example: `ocid1.volume.oc1.phx.abyhqljrgvttnlx73nmrwfaux7kcvzfs3s66izvxf2h4lgvyndsdsnoiwr5q`
		//
		ChapUsername string `json:"chapUsername,omitempty"`

		// The volume's iSCSI IP address.
		//
		// Example: `169.254.0.2`
		//
		// Required: true
		// Max Length: 15
		// Min Length: 7
		IPV4 *string `json:"ipv4"`

		// The target volume's iSCSI Qualified Name in the format defined by RFC 3720.
		//
		// Example: `iqn.2015-12.us.oracle.com:456b0391-17b8-4122-bbf1-f85fc0bb97d9`
		//
		// Required: true
		// Min Length: 1
		Iqn *string `json:"iqn"`

		// The volume's iSCSI port.
		//
		// Example: `3260`
		//
		// Required: true
		// Maximum: 65535
		// Minimum: 1
		Port *int64 `json:"port"`
	}

	data.AvailabilityDomain = m.availabilityDomainField
	data.CompartmentID = m.compartmentIdField
	data.DisplayName = m.displayNameField
	data.ID = m.idField
	data.InstanceID = m.instanceIdField
	data.LifecycleState = m.lifecycleStateField
	data.TimeCreated = m.timeCreatedField
	data.VolumeID = m.volumeIdField
	data.ChapSecret = m.ChapSecret
	data.ChapUsername = m.ChapUsername
	data.IPV4 = m.IPV4
	data.Iqn = m.Iqn
	data.Port = m.Port
	data.AttachmentType = DiscriminatorTypeValues["IScsiVolumeAttachment"]
	return json.Marshal(data)
}

// Validate validates this i scsi volume attachment
func (m *IScsiVolumeAttachment) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAvailabilityDomain(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCompartmentID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDisplayName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInstanceID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLifecycleState(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTimeCreated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVolumeID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIPV4(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIqn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePort(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IScsiVolumeAttachment) validateAvailabilityDomain(formats strfmt.Registry) error {

	if err := validate.Required("availabilityDomain", "body", m.AvailabilityDomain()); err != nil {
		return err
	}

	if err := validate.MinLength("availabilityDomain", "body", string(*m.AvailabilityDomain()), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("availabilityDomain", "body", string(*m.AvailabilityDomain()), 255); err != nil {
		return err
	}

	return nil
}

func (m *IScsiVolumeAttachment) validateCompartmentID(formats strfmt.Registry) error {

	if err := validate.Required("compartmentId", "body", m.CompartmentID()); err != nil {
		return err
	}

	if err := validate.MinLength("compartmentId", "body", string(*m.CompartmentID()), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("compartmentId", "body", string(*m.CompartmentID()), 255); err != nil {
		return err
	}

	return nil
}

func (m *IScsiVolumeAttachment) validateDisplayName(formats strfmt.Registry) error {

	if swag.IsZero(m.DisplayName()) { // not required
		return nil
	}

	if err := validate.MinLength("displayName", "body", string(m.DisplayName()), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("displayName", "body", string(m.DisplayName()), 255); err != nil {
		return err
	}

	return nil
}

func (m *IScsiVolumeAttachment) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID()); err != nil {
		return err
	}

	if err := validate.MinLength("id", "body", string(*m.ID()), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("id", "body", string(*m.ID()), 255); err != nil {
		return err
	}

	return nil
}

func (m *IScsiVolumeAttachment) validateInstanceID(formats strfmt.Registry) error {

	if err := validate.Required("instanceId", "body", m.InstanceID()); err != nil {
		return err
	}

	if err := validate.MinLength("instanceId", "body", string(*m.InstanceID()), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("instanceId", "body", string(*m.InstanceID()), 255); err != nil {
		return err
	}

	return nil
}

var iScsiVolumeAttachmentTypeLifecycleStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ATTACHING","ATTACHED","DETACHING","DETACHED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		iScsiVolumeAttachmentTypeLifecycleStatePropEnum = append(iScsiVolumeAttachmentTypeLifecycleStatePropEnum, v)
	}
}

// property enum
func (m *IScsiVolumeAttachment) validateLifecycleStateEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, iScsiVolumeAttachmentTypeLifecycleStatePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *IScsiVolumeAttachment) validateLifecycleState(formats strfmt.Registry) error {

	if err := validate.Required("lifecycleState", "body", m.LifecycleState()); err != nil {
		return err
	}

	// value enum
	if err := m.validateLifecycleStateEnum("lifecycleState", "body", *m.LifecycleState()); err != nil {
		return err
	}

	return nil
}

func (m *IScsiVolumeAttachment) validateTimeCreated(formats strfmt.Registry) error {

	if err := validate.Required("timeCreated", "body", m.TimeCreated()); err != nil {
		return err
	}

	return nil
}

func (m *IScsiVolumeAttachment) validateVolumeID(formats strfmt.Registry) error {

	if err := validate.Required("volumeId", "body", m.VolumeID()); err != nil {
		return err
	}

	if err := validate.MinLength("volumeId", "body", string(*m.VolumeID()), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("volumeId", "body", string(*m.VolumeID()), 255); err != nil {
		return err
	}

	return nil
}

func (m *IScsiVolumeAttachment) validateIPV4(formats strfmt.Registry) error {

	if err := validate.Required("ipv4", "body", m.IPV4); err != nil {
		return err
	}

	if err := validate.MinLength("ipv4", "body", string(*m.IPV4), 7); err != nil {
		return err
	}

	if err := validate.MaxLength("ipv4", "body", string(*m.IPV4), 15); err != nil {
		return err
	}

	return nil
}

func (m *IScsiVolumeAttachment) validateIqn(formats strfmt.Registry) error {

	if err := validate.Required("iqn", "body", m.Iqn); err != nil {
		return err
	}

	if err := validate.MinLength("iqn", "body", string(*m.Iqn), 1); err != nil {
		return err
	}

	return nil
}

func (m *IScsiVolumeAttachment) validatePort(formats strfmt.Registry) error {

	if err := validate.Required("port", "body", m.Port); err != nil {
		return err
	}

	if err := validate.MinimumInt("port", "body", int64(*m.Port), 1, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("port", "body", int64(*m.Port), 65535, false); err != nil {
		return err
	}

	return nil
}
