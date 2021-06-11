// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PacketAnalyzer Packet analyzer; analyzers count received traffic and generate statistics
// using various analysis methods.
//
//
// swagger:model PacketAnalyzer
type PacketAnalyzer struct {

	// Indicates whether this object is currently analyzing packets or not.
	//
	// Required: true
	Active *bool `json:"active"`

	// config
	// Required: true
	Config *PacketAnalyzerConfig `json:"config"`

	// Unique analyzer identifier
	// Required: true
	ID *string `json:"id"`

	// Specifies the unique source of packets for this analyzer. This
	// id may refer to either a port or an interface.
	//
	// Required: true
	SourceID *string `json:"source_id"`
}

// Validate validates this packet analyzer
func (m *PacketAnalyzer) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateActive(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSourceID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PacketAnalyzer) validateActive(formats strfmt.Registry) error {

	if err := validate.Required("active", "body", m.Active); err != nil {
		return err
	}

	return nil
}

func (m *PacketAnalyzer) validateConfig(formats strfmt.Registry) error {

	if err := validate.Required("config", "body", m.Config); err != nil {
		return err
	}

	if m.Config != nil {
		if err := m.Config.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("config")
			}
			return err
		}
	}

	return nil
}

func (m *PacketAnalyzer) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *PacketAnalyzer) validateSourceID(formats strfmt.Registry) error {

	if err := validate.Required("source_id", "body", m.SourceID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this packet analyzer based on the context it is used
func (m *PacketAnalyzer) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PacketAnalyzer) contextValidateConfig(ctx context.Context, formats strfmt.Registry) error {

	if m.Config != nil {
		if err := m.Config.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("config")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PacketAnalyzer) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PacketAnalyzer) UnmarshalBinary(b []byte) error {
	var res PacketAnalyzer
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}