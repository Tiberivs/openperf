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

// PacketAnalyzerConfig Packet analyzer configuration; the configuration specifies what packets
// are analyzed and what results are calculated.
//
//
// swagger:model PacketAnalyzerConfig
type PacketAnalyzerConfig struct {

	// Berkley Packet Filter (BPF) rules that matches input packets for this
	// analyzer to count. An empty rule, the default, matches all frames.
	//
	Filter string `json:"filter,omitempty"`

	// List of results to generate per flow for received packets.
	// Sequencing, latency, and jitter results require Spirent
	// signatures in the received packets. Pseudo Random Bit
	// Sequence (PRBS) results require packet payloads to contain
	// compatible PRBS data.
	//
	// Required: true
	// Unique: true
	FlowCounters []string `json:"flow_counters"`

	// List of result digests to generate per flow for received packets.
	// Sequence run length, latency, and jitter digests require Spirent
	// signatures in the received packets.
	//
	// Unique: true
	FlowDigests []string `json:"flow_digests"`

	// List of protocol counters to update per analyzer for received packets.
	//
	// Required: true
	// Unique: true
	ProtocolCounters []string `json:"protocol_counters"`
}

// Validate validates this packet analyzer config
func (m *PacketAnalyzerConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFlowCounters(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFlowDigests(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProtocolCounters(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var packetAnalyzerConfigFlowCountersItemsEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["advanced_sequencing","frame_count","frame_length","interarrival_time","jitter_ipdv","jitter_rfc","latency","prbs","header"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		packetAnalyzerConfigFlowCountersItemsEnum = append(packetAnalyzerConfigFlowCountersItemsEnum, v)
	}
}

func (m *PacketAnalyzerConfig) validateFlowCountersItemsEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, packetAnalyzerConfigFlowCountersItemsEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *PacketAnalyzerConfig) validateFlowCounters(formats strfmt.Registry) error {

	if err := validate.Required("flow_counters", "body", m.FlowCounters); err != nil {
		return err
	}

	if err := validate.UniqueItems("flow_counters", "body", m.FlowCounters); err != nil {
		return err
	}

	for i := 0; i < len(m.FlowCounters); i++ {

		// value enum
		if err := m.validateFlowCountersItemsEnum("flow_counters"+"."+strconv.Itoa(i), "body", m.FlowCounters[i]); err != nil {
			return err
		}

	}

	return nil
}

var packetAnalyzerConfigFlowDigestsItemsEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["frame_length","interarrival_time","jitter_ipdv","jitter_rfc","latency","sequence_run_length"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		packetAnalyzerConfigFlowDigestsItemsEnum = append(packetAnalyzerConfigFlowDigestsItemsEnum, v)
	}
}

func (m *PacketAnalyzerConfig) validateFlowDigestsItemsEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, packetAnalyzerConfigFlowDigestsItemsEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *PacketAnalyzerConfig) validateFlowDigests(formats strfmt.Registry) error {
	if swag.IsZero(m.FlowDigests) { // not required
		return nil
	}

	if err := validate.UniqueItems("flow_digests", "body", m.FlowDigests); err != nil {
		return err
	}

	for i := 0; i < len(m.FlowDigests); i++ {

		// value enum
		if err := m.validateFlowDigestsItemsEnum("flow_digests"+"."+strconv.Itoa(i), "body", m.FlowDigests[i]); err != nil {
			return err
		}

	}

	return nil
}

var packetAnalyzerConfigProtocolCountersItemsEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ethernet","ip","transport","tunnel","inner_ethernet","inner_ip","inner_transport"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		packetAnalyzerConfigProtocolCountersItemsEnum = append(packetAnalyzerConfigProtocolCountersItemsEnum, v)
	}
}

func (m *PacketAnalyzerConfig) validateProtocolCountersItemsEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, packetAnalyzerConfigProtocolCountersItemsEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *PacketAnalyzerConfig) validateProtocolCounters(formats strfmt.Registry) error {

	if err := validate.Required("protocol_counters", "body", m.ProtocolCounters); err != nil {
		return err
	}

	if err := validate.UniqueItems("protocol_counters", "body", m.ProtocolCounters); err != nil {
		return err
	}

	for i := 0; i < len(m.ProtocolCounters); i++ {

		// value enum
		if err := m.validateProtocolCountersItemsEnum("protocol_counters"+"."+strconv.Itoa(i), "body", m.ProtocolCounters[i]); err != nil {
			return err
		}

	}

	return nil
}

// ContextValidate validates this packet analyzer config based on context it is used
func (m *PacketAnalyzerConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PacketAnalyzerConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PacketAnalyzerConfig) UnmarshalBinary(b []byte) error {
	var res PacketAnalyzerConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
