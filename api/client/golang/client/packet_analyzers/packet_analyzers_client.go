// Code generated by go-swagger; DO NOT EDIT.

package packet_analyzers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new packet analyzers API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for packet analyzers API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	BulkCreatePacketAnalyzers(params *BulkCreatePacketAnalyzersParams, opts ...ClientOption) (*BulkCreatePacketAnalyzersOK, error)

	BulkDeletePacketAnalyzers(params *BulkDeletePacketAnalyzersParams, opts ...ClientOption) (*BulkDeletePacketAnalyzersNoContent, error)

	BulkStartPacketAnalyzers(params *BulkStartPacketAnalyzersParams, opts ...ClientOption) (*BulkStartPacketAnalyzersOK, error)

	BulkStopPacketAnalyzers(params *BulkStopPacketAnalyzersParams, opts ...ClientOption) (*BulkStopPacketAnalyzersNoContent, error)

	CreatePacketAnalyzer(params *CreatePacketAnalyzerParams, opts ...ClientOption) (*CreatePacketAnalyzerCreated, error)

	DeletePacketAnalyzer(params *DeletePacketAnalyzerParams, opts ...ClientOption) (*DeletePacketAnalyzerNoContent, error)

	DeletePacketAnalyzerResult(params *DeletePacketAnalyzerResultParams, opts ...ClientOption) (*DeletePacketAnalyzerResultNoContent, error)

	DeletePacketAnalyzerResults(params *DeletePacketAnalyzerResultsParams, opts ...ClientOption) (*DeletePacketAnalyzerResultsNoContent, error)

	DeletePacketAnalyzers(params *DeletePacketAnalyzersParams, opts ...ClientOption) (*DeletePacketAnalyzersNoContent, error)

	GetPacketAnalyzer(params *GetPacketAnalyzerParams, opts ...ClientOption) (*GetPacketAnalyzerOK, error)

	GetPacketAnalyzerResult(params *GetPacketAnalyzerResultParams, opts ...ClientOption) (*GetPacketAnalyzerResultOK, error)

	GetRxFlow(params *GetRxFlowParams, opts ...ClientOption) (*GetRxFlowOK, error)

	ListPacketAnalyzerResults(params *ListPacketAnalyzerResultsParams, opts ...ClientOption) (*ListPacketAnalyzerResultsOK, error)

	ListPacketAnalyzers(params *ListPacketAnalyzersParams, opts ...ClientOption) (*ListPacketAnalyzersOK, error)

	ListRxFlows(params *ListRxFlowsParams, opts ...ClientOption) (*ListRxFlowsOK, error)

	ResetPacketAnalyzer(params *ResetPacketAnalyzerParams, opts ...ClientOption) (*ResetPacketAnalyzerCreated, error)

	StartPacketAnalyzer(params *StartPacketAnalyzerParams, opts ...ClientOption) (*StartPacketAnalyzerCreated, error)

	StopPacketAnalyzer(params *StopPacketAnalyzerParams, opts ...ClientOption) (*StopPacketAnalyzerNoContent, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  BulkCreatePacketAnalyzers bulks create packet analyzers

  Create multiple packet analyzers. Requests are processed in an
all-or-nothing manner, i.e. a single analyzer creation failure
causes all analyzer creations for this request to fail.

*/
func (a *Client) BulkCreatePacketAnalyzers(params *BulkCreatePacketAnalyzersParams, opts ...ClientOption) (*BulkCreatePacketAnalyzersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewBulkCreatePacketAnalyzersParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "BulkCreatePacketAnalyzers",
		Method:             "POST",
		PathPattern:        "/packet/analyzers/x/bulk-create",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &BulkCreatePacketAnalyzersReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*BulkCreatePacketAnalyzersOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for BulkCreatePacketAnalyzers: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  BulkDeletePacketAnalyzers bulks delete packet analyzers

  Delete multiple packet analyzers in a best-effort manner. Analyzers
can only be deleted when inactive. Active or Non-existant analyzer ids
do not cause errors. Idempotent.

*/
func (a *Client) BulkDeletePacketAnalyzers(params *BulkDeletePacketAnalyzersParams, opts ...ClientOption) (*BulkDeletePacketAnalyzersNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewBulkDeletePacketAnalyzersParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "BulkDeletePacketAnalyzers",
		Method:             "POST",
		PathPattern:        "/packet/analyzers/x/bulk-delete",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &BulkDeletePacketAnalyzersReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*BulkDeletePacketAnalyzersNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for BulkDeletePacketAnalyzers: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  BulkStartPacketAnalyzers bulks start packet analyzers

  Start multiple packet analyzers simultaneously
*/
func (a *Client) BulkStartPacketAnalyzers(params *BulkStartPacketAnalyzersParams, opts ...ClientOption) (*BulkStartPacketAnalyzersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewBulkStartPacketAnalyzersParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "BulkStartPacketAnalyzers",
		Method:             "POST",
		PathPattern:        "/packet/analyzers/x/bulk-start",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &BulkStartPacketAnalyzersReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*BulkStartPacketAnalyzersOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for BulkStartPacketAnalyzers: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  BulkStopPacketAnalyzers bulks stop packet analyzers

  Stop multiple packet analyzers simultaneously
*/
func (a *Client) BulkStopPacketAnalyzers(params *BulkStopPacketAnalyzersParams, opts ...ClientOption) (*BulkStopPacketAnalyzersNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewBulkStopPacketAnalyzersParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "BulkStopPacketAnalyzers",
		Method:             "POST",
		PathPattern:        "/packet/analyzers/x/bulk-stop",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &BulkStopPacketAnalyzersReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*BulkStopPacketAnalyzersNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for BulkStopPacketAnalyzers: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  CreatePacketAnalyzer creates a packet analyzer

  Create a new packet analyzer.
*/
func (a *Client) CreatePacketAnalyzer(params *CreatePacketAnalyzerParams, opts ...ClientOption) (*CreatePacketAnalyzerCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreatePacketAnalyzerParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "CreatePacketAnalyzer",
		Method:             "POST",
		PathPattern:        "/packet/analyzers",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &CreatePacketAnalyzerReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreatePacketAnalyzerCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for CreatePacketAnalyzer: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DeletePacketAnalyzer deletes a packet analyzer

  Delete a stopped packet analyzer by id. Also delete all results
created by this analyzer. Idempotent.

*/
func (a *Client) DeletePacketAnalyzer(params *DeletePacketAnalyzerParams, opts ...ClientOption) (*DeletePacketAnalyzerNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeletePacketAnalyzerParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "DeletePacketAnalyzer",
		Method:             "DELETE",
		PathPattern:        "/packet/analyzers/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeletePacketAnalyzerReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeletePacketAnalyzerNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for DeletePacketAnalyzer: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DeletePacketAnalyzerResult deletes a packet analyzer result

  Delete an inactive packet analyzer result. Also deletes all child
rx-flow objects. Idempotent.

*/
func (a *Client) DeletePacketAnalyzerResult(params *DeletePacketAnalyzerResultParams, opts ...ClientOption) (*DeletePacketAnalyzerResultNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeletePacketAnalyzerResultParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "DeletePacketAnalyzerResult",
		Method:             "DELETE",
		PathPattern:        "/packet/analyzer-results/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeletePacketAnalyzerResultReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeletePacketAnalyzerResultNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for DeletePacketAnalyzerResult: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DeletePacketAnalyzerResults deletes all analyzer results

  Delete all inactive packet analyzer results
*/
func (a *Client) DeletePacketAnalyzerResults(params *DeletePacketAnalyzerResultsParams, opts ...ClientOption) (*DeletePacketAnalyzerResultsNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeletePacketAnalyzerResultsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "DeletePacketAnalyzerResults",
		Method:             "DELETE",
		PathPattern:        "/packet/analyzer-results",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeletePacketAnalyzerResultsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeletePacketAnalyzerResultsNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for DeletePacketAnalyzerResults: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DeletePacketAnalyzers deletes all packet analyzers

  Delete all inactive packet analyzers and their results. Idempotent.

*/
func (a *Client) DeletePacketAnalyzers(params *DeletePacketAnalyzersParams, opts ...ClientOption) (*DeletePacketAnalyzersNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeletePacketAnalyzersParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "DeletePacketAnalyzers",
		Method:             "DELETE",
		PathPattern:        "/packet/analyzers",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeletePacketAnalyzersReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeletePacketAnalyzersNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for DeletePacketAnalyzers: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetPacketAnalyzer gets a packet analyzer

  Return a packet analyzer by id.
*/
func (a *Client) GetPacketAnalyzer(params *GetPacketAnalyzerParams, opts ...ClientOption) (*GetPacketAnalyzerOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPacketAnalyzerParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetPacketAnalyzer",
		Method:             "GET",
		PathPattern:        "/packet/analyzers/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetPacketAnalyzerReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetPacketAnalyzerOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetPacketAnalyzer: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetPacketAnalyzerResult gets a packet analyzer result

  Returns results from a packet analyzer by result id.
*/
func (a *Client) GetPacketAnalyzerResult(params *GetPacketAnalyzerResultParams, opts ...ClientOption) (*GetPacketAnalyzerResultOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPacketAnalyzerResultParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetPacketAnalyzerResult",
		Method:             "GET",
		PathPattern:        "/packet/analyzer-results/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetPacketAnalyzerResultReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetPacketAnalyzerResultOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetPacketAnalyzerResult: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetRxFlow gets packet flow counters for a single flow

  Returns packet flow counters by id.
*/
func (a *Client) GetRxFlow(params *GetRxFlowParams, opts ...ClientOption) (*GetRxFlowOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetRxFlowParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetRxFlow",
		Method:             "GET",
		PathPattern:        "/packet/rx-flows/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetRxFlowReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetRxFlowOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetRxFlow: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ListPacketAnalyzerResults lists analyzer results

  The `analyzer-results` endpoint returns all analyzer results created
by analyzer instances.

*/
func (a *Client) ListPacketAnalyzerResults(params *ListPacketAnalyzerResultsParams, opts ...ClientOption) (*ListPacketAnalyzerResultsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListPacketAnalyzerResultsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ListPacketAnalyzerResults",
		Method:             "GET",
		PathPattern:        "/packet/analyzer-results",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ListPacketAnalyzerResultsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListPacketAnalyzerResultsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ListPacketAnalyzerResults: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ListPacketAnalyzers lists packet analyzers

  The `analyzers` endpoint returns all packet analyzers that are
configured to collect and report port and flow statistics.

*/
func (a *Client) ListPacketAnalyzers(params *ListPacketAnalyzersParams, opts ...ClientOption) (*ListPacketAnalyzersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListPacketAnalyzersParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ListPacketAnalyzers",
		Method:             "GET",
		PathPattern:        "/packet/analyzers",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ListPacketAnalyzersReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListPacketAnalyzersOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ListPacketAnalyzers: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ListRxFlows lists received packet flows

  The `rx-flows` endpoint returns all packet flows that have been
received by analyzer instances.

*/
func (a *Client) ListRxFlows(params *ListRxFlowsParams, opts ...ClientOption) (*ListRxFlowsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListRxFlowsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ListRxFlows",
		Method:             "GET",
		PathPattern:        "/packet/rx-flows",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ListRxFlowsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListRxFlowsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ListRxFlows: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ResetPacketAnalyzer resets a running analyzer

  Used to generate a new result for a running analyzer. This
method effective resets all analyzer counters to zero. Note
that the new analyzer result will not contain any flow results
until packets are received after the reset event. Creates a
new analyzer result on success.

*/
func (a *Client) ResetPacketAnalyzer(params *ResetPacketAnalyzerParams, opts ...ClientOption) (*ResetPacketAnalyzerCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewResetPacketAnalyzerParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ResetPacketAnalyzer",
		Method:             "POST",
		PathPattern:        "/packet/analyzers/{id}/reset",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ResetPacketAnalyzerReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ResetPacketAnalyzerCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ResetPacketAnalyzer: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  StartPacketAnalyzer starts analyzing and collecting packet statistics

  Used to start a non-running analyzer. Creates a new analyzer
result on success.

*/
func (a *Client) StartPacketAnalyzer(params *StartPacketAnalyzerParams, opts ...ClientOption) (*StartPacketAnalyzerCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewStartPacketAnalyzerParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "StartPacketAnalyzer",
		Method:             "POST",
		PathPattern:        "/packet/analyzers/{id}/start",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &StartPacketAnalyzerReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*StartPacketAnalyzerCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for StartPacketAnalyzer: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  StopPacketAnalyzer stops analyzing and collecting packet statistics

  Use to halt a running analyzer. Idempotent.
*/
func (a *Client) StopPacketAnalyzer(params *StopPacketAnalyzerParams, opts ...ClientOption) (*StopPacketAnalyzerNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewStopPacketAnalyzerParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "StopPacketAnalyzer",
		Method:             "POST",
		PathPattern:        "/packet/analyzers/{id}/stop",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &StopPacketAnalyzerReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*StopPacketAnalyzerNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for StopPacketAnalyzer: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
