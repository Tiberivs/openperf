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

// PacketProtocolIPV4 Describes an IPv4 header
//
// swagger:model PacketProtocolIpv4
type PacketProtocolIPV4 struct {

	// IPv4 checksum
	// Maximum: 65535
	// Minimum: 0
	Checksum *int32 `json:"checksum,omitempty"`

	// IPv4 destination address
	// Pattern: ^((25[0-5]|2[0-4][0-9]|[01]?[1-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[1-9]?[0-9])$
	Destination string `json:"destination,omitempty"`

	// IPv4 dscp
	// Maximum: 63
	// Minimum: 0
	Dscp *int32 `json:"dscp,omitempty"`

	// IPv4 ecn
	// Enum: [non_ect ect_0 ect_1 ce]
	Ecn string `json:"ecn,omitempty"`

	// IPv4 flags
	Flags []string `json:"flags"`

	// IPv4 fragment offset
	// Maximum: 65528
	// Minimum: 0
	// Multiple Of: 8
	FragmentOffset *int32 `json:"fragment_offset,omitempty"`

	// IPv4 header length
	// Maximum: 60
	// Minimum: 0
	// Multiple Of: 4
	HeaderLength *int32 `json:"header_length,omitempty"`

	// IPv4 identification
	// Maximum: 65535
	// Minimum: 0
	Identification *int32 `json:"identification,omitempty"`

	// IPv4 protocol
	// Maximum: 255
	// Minimum: 0
	Protocol *int32 `json:"protocol,omitempty"`

	// IPv4 source address
	// Pattern: ^((25[0-5]|2[0-4][0-9]|[01]?[1-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[1-9]?[0-9])$
	Source string `json:"source,omitempty"`

	// IPv4 time to live
	// Maximum: 255
	// Minimum: 0
	TimeToLive *int32 `json:"time_to_live,omitempty"`

	// IPv4 total length
	// Maximum: 65535
	// Minimum: 0
	TotalLength *int32 `json:"total_length,omitempty"`

	// IPv4 version
	// Maximum: 15
	// Minimum: 0
	Version *int32 `json:"version,omitempty"`
}

// Validate validates this packet protocol Ipv4
func (m *PacketProtocolIPV4) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateChecksum(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDestination(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDscp(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEcn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFlags(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFragmentOffset(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHeaderLength(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIdentification(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProtocol(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSource(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTimeToLive(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTotalLength(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVersion(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PacketProtocolIPV4) validateChecksum(formats strfmt.Registry) error {
	if swag.IsZero(m.Checksum) { // not required
		return nil
	}

	if err := validate.MinimumInt("checksum", "body", int64(*m.Checksum), 0, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("checksum", "body", int64(*m.Checksum), 65535, false); err != nil {
		return err
	}

	return nil
}

func (m *PacketProtocolIPV4) validateDestination(formats strfmt.Registry) error {
	if swag.IsZero(m.Destination) { // not required
		return nil
	}

	if err := validate.Pattern("destination", "body", m.Destination, `^((25[0-5]|2[0-4][0-9]|[01]?[1-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[1-9]?[0-9])$`); err != nil {
		return err
	}

	return nil
}

func (m *PacketProtocolIPV4) validateDscp(formats strfmt.Registry) error {
	if swag.IsZero(m.Dscp) { // not required
		return nil
	}

	if err := validate.MinimumInt("dscp", "body", int64(*m.Dscp), 0, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("dscp", "body", int64(*m.Dscp), 63, false); err != nil {
		return err
	}

	return nil
}

var packetProtocolIpv4TypeEcnPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["non_ect","ect_0","ect_1","ce"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		packetProtocolIpv4TypeEcnPropEnum = append(packetProtocolIpv4TypeEcnPropEnum, v)
	}
}

const (

	// PacketProtocolIPV4EcnNonEct captures enum value "non_ect"
	PacketProtocolIPV4EcnNonEct string = "non_ect"

	// PacketProtocolIPV4EcnEct0 captures enum value "ect_0"
	PacketProtocolIPV4EcnEct0 string = "ect_0"

	// PacketProtocolIPV4EcnEct1 captures enum value "ect_1"
	PacketProtocolIPV4EcnEct1 string = "ect_1"

	// PacketProtocolIPV4EcnCe captures enum value "ce"
	PacketProtocolIPV4EcnCe string = "ce"
)

// prop value enum
func (m *PacketProtocolIPV4) validateEcnEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, packetProtocolIpv4TypeEcnPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *PacketProtocolIPV4) validateEcn(formats strfmt.Registry) error {
	if swag.IsZero(m.Ecn) { // not required
		return nil
	}

	// value enum
	if err := m.validateEcnEnum("ecn", "body", m.Ecn); err != nil {
		return err
	}

	return nil
}

var packetProtocolIpv4FlagsItemsEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["evil_bit","dont_fragment","more_fragments"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		packetProtocolIpv4FlagsItemsEnum = append(packetProtocolIpv4FlagsItemsEnum, v)
	}
}

func (m *PacketProtocolIPV4) validateFlagsItemsEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, packetProtocolIpv4FlagsItemsEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *PacketProtocolIPV4) validateFlags(formats strfmt.Registry) error {
	if swag.IsZero(m.Flags) { // not required
		return nil
	}

	for i := 0; i < len(m.Flags); i++ {

		// value enum
		if err := m.validateFlagsItemsEnum("flags"+"."+strconv.Itoa(i), "body", m.Flags[i]); err != nil {
			return err
		}

	}

	return nil
}

func (m *PacketProtocolIPV4) validateFragmentOffset(formats strfmt.Registry) error {
	if swag.IsZero(m.FragmentOffset) { // not required
		return nil
	}

	if err := validate.MinimumInt("fragment_offset", "body", int64(*m.FragmentOffset), 0, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("fragment_offset", "body", int64(*m.FragmentOffset), 65528, false); err != nil {
		return err
	}

	if err := validate.MultipleOfInt("fragment_offset", "body", int64(*m.FragmentOffset), 8); err != nil {
		return err
	}

	return nil
}

func (m *PacketProtocolIPV4) validateHeaderLength(formats strfmt.Registry) error {
	if swag.IsZero(m.HeaderLength) { // not required
		return nil
	}

	if err := validate.MinimumInt("header_length", "body", int64(*m.HeaderLength), 0, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("header_length", "body", int64(*m.HeaderLength), 60, false); err != nil {
		return err
	}

	if err := validate.MultipleOfInt("header_length", "body", int64(*m.HeaderLength), 4); err != nil {
		return err
	}

	return nil
}

func (m *PacketProtocolIPV4) validateIdentification(formats strfmt.Registry) error {
	if swag.IsZero(m.Identification) { // not required
		return nil
	}

	if err := validate.MinimumInt("identification", "body", int64(*m.Identification), 0, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("identification", "body", int64(*m.Identification), 65535, false); err != nil {
		return err
	}

	return nil
}

func (m *PacketProtocolIPV4) validateProtocol(formats strfmt.Registry) error {
	if swag.IsZero(m.Protocol) { // not required
		return nil
	}

	if err := validate.MinimumInt("protocol", "body", int64(*m.Protocol), 0, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("protocol", "body", int64(*m.Protocol), 255, false); err != nil {
		return err
	}

	return nil
}

func (m *PacketProtocolIPV4) validateSource(formats strfmt.Registry) error {
	if swag.IsZero(m.Source) { // not required
		return nil
	}

	if err := validate.Pattern("source", "body", m.Source, `^((25[0-5]|2[0-4][0-9]|[01]?[1-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[1-9]?[0-9])$`); err != nil {
		return err
	}

	return nil
}

func (m *PacketProtocolIPV4) validateTimeToLive(formats strfmt.Registry) error {
	if swag.IsZero(m.TimeToLive) { // not required
		return nil
	}

	if err := validate.MinimumInt("time_to_live", "body", int64(*m.TimeToLive), 0, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("time_to_live", "body", int64(*m.TimeToLive), 255, false); err != nil {
		return err
	}

	return nil
}

func (m *PacketProtocolIPV4) validateTotalLength(formats strfmt.Registry) error {
	if swag.IsZero(m.TotalLength) { // not required
		return nil
	}

	if err := validate.MinimumInt("total_length", "body", int64(*m.TotalLength), 0, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("total_length", "body", int64(*m.TotalLength), 65535, false); err != nil {
		return err
	}

	return nil
}

func (m *PacketProtocolIPV4) validateVersion(formats strfmt.Registry) error {
	if swag.IsZero(m.Version) { // not required
		return nil
	}

	if err := validate.MinimumInt("version", "body", int64(*m.Version), 0, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("version", "body", int64(*m.Version), 15, false); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this packet protocol Ipv4 based on context it is used
func (m *PacketProtocolIPV4) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PacketProtocolIPV4) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PacketProtocolIPV4) UnmarshalBinary(b []byte) error {
	var res PacketProtocolIPV4
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}