// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// TrafficPacketTemplate Defines a set of packets for the packet generator
//
// swagger:model TrafficPacketTemplate
type TrafficPacketTemplate struct {

	// Specifies how modifiers from different protocols are combined in the
	// packet template
	//
	// Enum: [cartesian zip]
	ModifierTie *string `json:"modifier_tie,omitempty"`

	// Ordered list of packet protocols. Packets are constructed
	// by iterating over the sequence of protocols.
	//
	// Required: true
	// Min Items: 1
	Protocols []*TrafficProtocol `json:"protocols"`
}

// Validate validates this traffic packet template
func (m *TrafficPacketTemplate) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateModifierTie(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProtocols(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var trafficPacketTemplateTypeModifierTiePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["cartesian","zip"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		trafficPacketTemplateTypeModifierTiePropEnum = append(trafficPacketTemplateTypeModifierTiePropEnum, v)
	}
}

const (

	// TrafficPacketTemplateModifierTieCartesian captures enum value "cartesian"
	TrafficPacketTemplateModifierTieCartesian string = "cartesian"

	// TrafficPacketTemplateModifierTieZip captures enum value "zip"
	TrafficPacketTemplateModifierTieZip string = "zip"
)

// prop value enum
func (m *TrafficPacketTemplate) validateModifierTieEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, trafficPacketTemplateTypeModifierTiePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *TrafficPacketTemplate) validateModifierTie(formats strfmt.Registry) error {
	if swag.IsZero(m.ModifierTie) { // not required
		return nil
	}

	// value enum
	if err := m.validateModifierTieEnum("modifier_tie", "body", *m.ModifierTie); err != nil {
		return err
	}

	return nil
}

func (m *TrafficPacketTemplate) validateProtocols(formats strfmt.Registry) error {

	if err := validate.Required("protocols", "body", m.Protocols); err != nil {
		return err
	}

	iProtocolsSize := int64(len(m.Protocols))

	if err := validate.MinItems("protocols", "body", iProtocolsSize, 1); err != nil {
		return err
	}

	for i := 0; i < len(m.Protocols); i++ {
		if swag.IsZero(m.Protocols[i]) { // not required
			continue
		}

		if m.Protocols[i] != nil {
			if err := m.Protocols[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("protocols" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this traffic packet template based on the context it is used
func (m *TrafficPacketTemplate) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateProtocols(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TrafficPacketTemplate) contextValidateProtocols(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Protocols); i++ {

		if m.Protocols[i] != nil {
			if err := m.Protocols[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("protocols" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *TrafficPacketTemplate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TrafficPacketTemplate) UnmarshalBinary(b []byte) error {
	var res TrafficPacketTemplate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
