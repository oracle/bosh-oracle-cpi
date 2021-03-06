package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// TCPOptions Optional object to specify ports for a TCP rule. If you specify TCP as the
// protocol but omit this object, then all ports are allowed.
//
// swagger:model TcpOptions
type TCPOptions struct {

	// An inclusive range of allowed destination ports. Use the same number for the min and max
	// to indicate a single port. Defaults to all ports if not specified.
	//
	DestinationPortRange *PortRange `json:"destinationPortRange,omitempty"`

	// An inclusive range of allowed source ports. Use the same number for the min and max to
	// indicate a single port. Defaults to all ports if not specified.
	//
	SourcePortRange *PortRange `json:"sourcePortRange,omitempty"`
}

// Validate validates this Tcp options
func (m *TCPOptions) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDestinationPortRange(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateSourcePortRange(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TCPOptions) validateDestinationPortRange(formats strfmt.Registry) error {

	if swag.IsZero(m.DestinationPortRange) { // not required
		return nil
	}

	if m.DestinationPortRange != nil {

		if err := m.DestinationPortRange.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("destinationPortRange")
			}
			return err
		}
	}

	return nil
}

func (m *TCPOptions) validateSourcePortRange(formats strfmt.Registry) error {

	if swag.IsZero(m.SourcePortRange) { // not required
		return nil
	}

	if m.SourcePortRange != nil {

		if err := m.SourcePortRange.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("sourcePortRange")
			}
			return err
		}
	}

	return nil
}
