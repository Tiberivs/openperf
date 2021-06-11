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

// RxFlow Counters for a flow of packets
//
// swagger:model RxFlow
type RxFlow struct {

	// Unique analyzer result identifier that observed this flow
	// Required: true
	AnalyzerResultID *string `json:"analyzer_result_id"`

	// counters
	// Required: true
	Counters *PacketAnalyzerFlowCounters `json:"counters"`

	// digests
	// Required: true
	Digests *PacketAnalyzerFlowDigests `json:"digests"`

	// Unique received flow identifier
	// Required: true
	ID *string `json:"id"`
}

// Validate validates this rx flow
func (m *RxFlow) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAnalyzerResultID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCounters(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDigests(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RxFlow) validateAnalyzerResultID(formats strfmt.Registry) error {

	if err := validate.Required("analyzer_result_id", "body", m.AnalyzerResultID); err != nil {
		return err
	}

	return nil
}

func (m *RxFlow) validateCounters(formats strfmt.Registry) error {

	if err := validate.Required("counters", "body", m.Counters); err != nil {
		return err
	}

	if m.Counters != nil {
		if err := m.Counters.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("counters")
			}
			return err
		}
	}

	return nil
}

func (m *RxFlow) validateDigests(formats strfmt.Registry) error {

	if err := validate.Required("digests", "body", m.Digests); err != nil {
		return err
	}

	if m.Digests != nil {
		if err := m.Digests.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("digests")
			}
			return err
		}
	}

	return nil
}

func (m *RxFlow) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this rx flow based on the context it is used
func (m *RxFlow) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCounters(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDigests(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RxFlow) contextValidateCounters(ctx context.Context, formats strfmt.Registry) error {

	if m.Counters != nil {
		if err := m.Counters.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("counters")
			}
			return err
		}
	}

	return nil
}

func (m *RxFlow) contextValidateDigests(ctx context.Context, formats strfmt.Registry) error {

	if m.Digests != nil {
		if err := m.Digests.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("digests")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RxFlow) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RxFlow) UnmarshalBinary(b []byte) error {
	var res RxFlow
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}