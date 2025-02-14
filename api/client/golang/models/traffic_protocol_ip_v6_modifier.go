// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// TrafficProtocolIPV6Modifier Specifies how to modify an IPv6 address
//
// swagger:model trafficProtocolIpv6Modifier
type TrafficProtocolIPV6Modifier struct {

	// List of IPv6 addresses
	// Min Items: 1
	List []string `json:"list"`

	// sequence
	Sequence *TrafficProtocolIPV6ModifierSequence `json:"sequence,omitempty"`
}

// Validate validates this traffic protocol Ipv6 modifier
func (m *TrafficProtocolIPV6Modifier) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateList(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSequence(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TrafficProtocolIPV6Modifier) validateList(formats strfmt.Registry) error {
	if swag.IsZero(m.List) { // not required
		return nil
	}

	iListSize := int64(len(m.List))

	if err := validate.MinItems("list", "body", iListSize, 1); err != nil {
		return err
	}

	for i := 0; i < len(m.List); i++ {

		if err := validate.Pattern("list"+"."+strconv.Itoa(i), "body", m.List[i], `^((::[0-9a-fA-F]{1,4})|([0-9a-fA-F]{1,4}::)|(([0-9a-fA-F]{1,4}:){1,6}:[0-9a-fA-F])|(([0-9a-fA-F]{1,4}:){7}[0-9a-fA-F]{1,4}))$`); err != nil {
			return err
		}

	}

	return nil
}

func (m *TrafficProtocolIPV6Modifier) validateSequence(formats strfmt.Registry) error {
	if swag.IsZero(m.Sequence) { // not required
		return nil
	}

	if m.Sequence != nil {
		if err := m.Sequence.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("sequence")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this traffic protocol Ipv6 modifier based on the context it is used
func (m *TrafficProtocolIPV6Modifier) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSequence(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TrafficProtocolIPV6Modifier) contextValidateSequence(ctx context.Context, formats strfmt.Registry) error {

	if m.Sequence != nil {
		if err := m.Sequence.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("sequence")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TrafficProtocolIPV6Modifier) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TrafficProtocolIPV6Modifier) UnmarshalBinary(b []byte) error {
	var res TrafficProtocolIPV6Modifier
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// TrafficProtocolIPV6ModifierSequence Specifies a sequence of IPv6 addresses
//
// swagger:model TrafficProtocolIPV6ModifierSequence
type TrafficProtocolIPV6ModifierSequence struct {

	// The number of addresses in the sequence
	// Required: true
	// Minimum: 1
	Count *int32 `json:"count"`

	// List of addresses in the sequence to skip
	Skip []string `json:"skip"`

	// First IPv6 address in the sequence
	// Required: true
	// Pattern: ^((::[0-9a-fA-F]{1,4})|([0-9a-fA-F]{1,4}::)|(([0-9a-fA-F]{1,4}:){1,6}:[0-9a-fA-F])|(([0-9a-fA-F]{1,4}:){7}[0-9a-fA-F]{1,4}))$
	Start *string `json:"start"`

	// Last IPv6 address in the sequence
	// Pattern: ^((::[0-9a-fA-F]{1,4})|([0-9a-fA-F]{1,4}::)|(([0-9a-fA-F]{1,4}:){1,6}:[0-9a-fA-F])|(([0-9a-fA-F]{1,4}:){7}[0-9a-fA-F]{1,4}))$
	Stop string `json:"stop,omitempty"`
}

// Validate validates this traffic protocol IP v6 modifier sequence
func (m *TrafficProtocolIPV6ModifierSequence) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSkip(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStart(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStop(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TrafficProtocolIPV6ModifierSequence) validateCount(formats strfmt.Registry) error {

	if err := validate.Required("sequence"+"."+"count", "body", m.Count); err != nil {
		return err
	}

	if err := validate.MinimumInt("sequence"+"."+"count", "body", int64(*m.Count), 1, false); err != nil {
		return err
	}

	return nil
}

func (m *TrafficProtocolIPV6ModifierSequence) validateSkip(formats strfmt.Registry) error {
	if swag.IsZero(m.Skip) { // not required
		return nil
	}

	for i := 0; i < len(m.Skip); i++ {

		if err := validate.Pattern("sequence"+"."+"skip"+"."+strconv.Itoa(i), "body", m.Skip[i], `^((::[0-9a-fA-F]{1,4})|([0-9a-fA-F]{1,4}::)|(([0-9a-fA-F]{1,4}:){1,6}:[0-9a-fA-F])|(([0-9a-fA-F]{1,4}:){7}[0-9a-fA-F]{1,4}))$`); err != nil {
			return err
		}

	}

	return nil
}

func (m *TrafficProtocolIPV6ModifierSequence) validateStart(formats strfmt.Registry) error {

	if err := validate.Required("sequence"+"."+"start", "body", m.Start); err != nil {
		return err
	}

	if err := validate.Pattern("sequence"+"."+"start", "body", *m.Start, `^((::[0-9a-fA-F]{1,4})|([0-9a-fA-F]{1,4}::)|(([0-9a-fA-F]{1,4}:){1,6}:[0-9a-fA-F])|(([0-9a-fA-F]{1,4}:){7}[0-9a-fA-F]{1,4}))$`); err != nil {
		return err
	}

	return nil
}

func (m *TrafficProtocolIPV6ModifierSequence) validateStop(formats strfmt.Registry) error {
	if swag.IsZero(m.Stop) { // not required
		return nil
	}

	if err := validate.Pattern("sequence"+"."+"stop", "body", m.Stop, `^((::[0-9a-fA-F]{1,4})|([0-9a-fA-F]{1,4}::)|(([0-9a-fA-F]{1,4}:){1,6}:[0-9a-fA-F])|(([0-9a-fA-F]{1,4}:){7}[0-9a-fA-F]{1,4}))$`); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this traffic protocol IP v6 modifier sequence based on context it is used
func (m *TrafficProtocolIPV6ModifierSequence) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TrafficProtocolIPV6ModifierSequence) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TrafficProtocolIPV6ModifierSequence) UnmarshalBinary(b []byte) error {
	var res TrafficProtocolIPV6ModifierSequence
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
