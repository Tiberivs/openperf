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

// StackMemoryStats Stack memory statistics
//
// swagger:model StackMemoryStats
type StackMemoryStats struct {

	// Available (bytes free)
	// Required: true
	Available *int64 `json:"available"`

	// Errors (count)
	// Required: true
	Errors *int64 `json:"errors"`

	// Illegal operations (count)
	// Required: true
	Illegal *int64 `json:"illegal"`

	// Maximum (bytes allocated)
	// Required: true
	Max *int64 `json:"max"`

	// Pool name
	// Required: true
	Name *string `json:"name"`

	// Used (bytes allocated)
	// Required: true
	Used *int64 `json:"used"`
}

// Validate validates this stack memory stats
func (m *StackMemoryStats) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAvailable(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateErrors(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIllegal(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMax(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUsed(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *StackMemoryStats) validateAvailable(formats strfmt.Registry) error {

	if err := validate.Required("available", "body", m.Available); err != nil {
		return err
	}

	return nil
}

func (m *StackMemoryStats) validateErrors(formats strfmt.Registry) error {

	if err := validate.Required("errors", "body", m.Errors); err != nil {
		return err
	}

	return nil
}

func (m *StackMemoryStats) validateIllegal(formats strfmt.Registry) error {

	if err := validate.Required("illegal", "body", m.Illegal); err != nil {
		return err
	}

	return nil
}

func (m *StackMemoryStats) validateMax(formats strfmt.Registry) error {

	if err := validate.Required("max", "body", m.Max); err != nil {
		return err
	}

	return nil
}

func (m *StackMemoryStats) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *StackMemoryStats) validateUsed(formats strfmt.Registry) error {

	if err := validate.Required("used", "body", m.Used); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this stack memory stats based on context it is used
func (m *StackMemoryStats) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *StackMemoryStats) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StackMemoryStats) UnmarshalBinary(b []byte) error {
	var res StackMemoryStats
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
