// Code generated by go-swagger; DO NOT EDIT.

package packet_analyzers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewListRxFlowsParams creates a new ListRxFlowsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListRxFlowsParams() *ListRxFlowsParams {
	return &ListRxFlowsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListRxFlowsParamsWithTimeout creates a new ListRxFlowsParams object
// with the ability to set a timeout on a request.
func NewListRxFlowsParamsWithTimeout(timeout time.Duration) *ListRxFlowsParams {
	return &ListRxFlowsParams{
		timeout: timeout,
	}
}

// NewListRxFlowsParamsWithContext creates a new ListRxFlowsParams object
// with the ability to set a context for a request.
func NewListRxFlowsParamsWithContext(ctx context.Context) *ListRxFlowsParams {
	return &ListRxFlowsParams{
		Context: ctx,
	}
}

// NewListRxFlowsParamsWithHTTPClient creates a new ListRxFlowsParams object
// with the ability to set a custom HTTPClient for a request.
func NewListRxFlowsParamsWithHTTPClient(client *http.Client) *ListRxFlowsParams {
	return &ListRxFlowsParams{
		HTTPClient: client,
	}
}

/* ListRxFlowsParams contains all the parameters to send to the API endpoint
   for the list rx flows operation.

   Typically these are written to a http.Request.
*/
type ListRxFlowsParams struct {

	/* AnalyzerID.

	   Filter by receive analyzer id
	*/
	AnalyzerID *string

	/* SourceID.

	   Filter by receive port or interface id
	*/
	SourceID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list rx flows params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListRxFlowsParams) WithDefaults() *ListRxFlowsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list rx flows params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListRxFlowsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list rx flows params
func (o *ListRxFlowsParams) WithTimeout(timeout time.Duration) *ListRxFlowsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list rx flows params
func (o *ListRxFlowsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list rx flows params
func (o *ListRxFlowsParams) WithContext(ctx context.Context) *ListRxFlowsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list rx flows params
func (o *ListRxFlowsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list rx flows params
func (o *ListRxFlowsParams) WithHTTPClient(client *http.Client) *ListRxFlowsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list rx flows params
func (o *ListRxFlowsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAnalyzerID adds the analyzerID to the list rx flows params
func (o *ListRxFlowsParams) WithAnalyzerID(analyzerID *string) *ListRxFlowsParams {
	o.SetAnalyzerID(analyzerID)
	return o
}

// SetAnalyzerID adds the analyzerId to the list rx flows params
func (o *ListRxFlowsParams) SetAnalyzerID(analyzerID *string) {
	o.AnalyzerID = analyzerID
}

// WithSourceID adds the sourceID to the list rx flows params
func (o *ListRxFlowsParams) WithSourceID(sourceID *string) *ListRxFlowsParams {
	o.SetSourceID(sourceID)
	return o
}

// SetSourceID adds the sourceId to the list rx flows params
func (o *ListRxFlowsParams) SetSourceID(sourceID *string) {
	o.SourceID = sourceID
}

// WriteToRequest writes these params to a swagger request
func (o *ListRxFlowsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.AnalyzerID != nil {

		// query param analyzer_id
		var qrAnalyzerID string

		if o.AnalyzerID != nil {
			qrAnalyzerID = *o.AnalyzerID
		}
		qAnalyzerID := qrAnalyzerID
		if qAnalyzerID != "" {

			if err := r.SetQueryParam("analyzer_id", qAnalyzerID); err != nil {
				return err
			}
		}
	}

	if o.SourceID != nil {

		// query param source_id
		var qrSourceID string

		if o.SourceID != nil {
			qrSourceID = *o.SourceID
		}
		qSourceID := qrSourceID
		if qSourceID != "" {

			if err := r.SetQueryParam("source_id", qSourceID); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}