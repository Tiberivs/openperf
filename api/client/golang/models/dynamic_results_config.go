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
)

// DynamicResultsConfig Configurations of dynamic results
//
// swagger:model DynamicResultsConfig
type DynamicResultsConfig struct {

	// Array of T-Digest configurations
	Tdigests []*TDigestConfig `json:"tdigests"`

	// Array of Threshold configurations
	Thresholds []*ThresholdConfig `json:"thresholds"`
}

// Validate validates this dynamic results config
func (m *DynamicResultsConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTdigests(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateThresholds(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DynamicResultsConfig) validateTdigests(formats strfmt.Registry) error {
	if swag.IsZero(m.Tdigests) { // not required
		return nil
	}

	for i := 0; i < len(m.Tdigests); i++ {
		if swag.IsZero(m.Tdigests[i]) { // not required
			continue
		}

		if m.Tdigests[i] != nil {
			if err := m.Tdigests[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("tdigests" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *DynamicResultsConfig) validateThresholds(formats strfmt.Registry) error {
	if swag.IsZero(m.Thresholds) { // not required
		return nil
	}

	for i := 0; i < len(m.Thresholds); i++ {
		if swag.IsZero(m.Thresholds[i]) { // not required
			continue
		}

		if m.Thresholds[i] != nil {
			if err := m.Thresholds[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("thresholds" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this dynamic results config based on the context it is used
func (m *DynamicResultsConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateTdigests(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateThresholds(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DynamicResultsConfig) contextValidateTdigests(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Tdigests); i++ {

		if m.Tdigests[i] != nil {
			if err := m.Tdigests[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("tdigests" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *DynamicResultsConfig) contextValidateThresholds(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Thresholds); i++ {

		if m.Thresholds[i] != nil {
			if err := m.Thresholds[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("thresholds" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *DynamicResultsConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DynamicResultsConfig) UnmarshalBinary(b []byte) error {
	var res DynamicResultsConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}