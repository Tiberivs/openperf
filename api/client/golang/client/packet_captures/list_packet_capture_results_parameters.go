// Code generated by go-swagger; DO NOT EDIT.

package packet_captures

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

// NewListPacketCaptureResultsParams creates a new ListPacketCaptureResultsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListPacketCaptureResultsParams() *ListPacketCaptureResultsParams {
	return &ListPacketCaptureResultsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListPacketCaptureResultsParamsWithTimeout creates a new ListPacketCaptureResultsParams object
// with the ability to set a timeout on a request.
func NewListPacketCaptureResultsParamsWithTimeout(timeout time.Duration) *ListPacketCaptureResultsParams {
	return &ListPacketCaptureResultsParams{
		timeout: timeout,
	}
}

// NewListPacketCaptureResultsParamsWithContext creates a new ListPacketCaptureResultsParams object
// with the ability to set a context for a request.
func NewListPacketCaptureResultsParamsWithContext(ctx context.Context) *ListPacketCaptureResultsParams {
	return &ListPacketCaptureResultsParams{
		Context: ctx,
	}
}

// NewListPacketCaptureResultsParamsWithHTTPClient creates a new ListPacketCaptureResultsParams object
// with the ability to set a custom HTTPClient for a request.
func NewListPacketCaptureResultsParamsWithHTTPClient(client *http.Client) *ListPacketCaptureResultsParams {
	return &ListPacketCaptureResultsParams{
		HTTPClient: client,
	}
}

/* ListPacketCaptureResultsParams contains all the parameters to send to the API endpoint
   for the list packet capture results operation.

   Typically these are written to a http.Request.
*/
type ListPacketCaptureResultsParams struct {

	/* CaptureID.

	   Filter by capture id
	*/
	CaptureID *string

	/* SourceID.

	   Filter by receive port or interface id
	*/
	SourceID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list packet capture results params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListPacketCaptureResultsParams) WithDefaults() *ListPacketCaptureResultsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list packet capture results params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListPacketCaptureResultsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list packet capture results params
func (o *ListPacketCaptureResultsParams) WithTimeout(timeout time.Duration) *ListPacketCaptureResultsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list packet capture results params
func (o *ListPacketCaptureResultsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list packet capture results params
func (o *ListPacketCaptureResultsParams) WithContext(ctx context.Context) *ListPacketCaptureResultsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list packet capture results params
func (o *ListPacketCaptureResultsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list packet capture results params
func (o *ListPacketCaptureResultsParams) WithHTTPClient(client *http.Client) *ListPacketCaptureResultsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list packet capture results params
func (o *ListPacketCaptureResultsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCaptureID adds the captureID to the list packet capture results params
func (o *ListPacketCaptureResultsParams) WithCaptureID(captureID *string) *ListPacketCaptureResultsParams {
	o.SetCaptureID(captureID)
	return o
}

// SetCaptureID adds the captureId to the list packet capture results params
func (o *ListPacketCaptureResultsParams) SetCaptureID(captureID *string) {
	o.CaptureID = captureID
}

// WithSourceID adds the sourceID to the list packet capture results params
func (o *ListPacketCaptureResultsParams) WithSourceID(sourceID *string) *ListPacketCaptureResultsParams {
	o.SetSourceID(sourceID)
	return o
}

// SetSourceID adds the sourceId to the list packet capture results params
func (o *ListPacketCaptureResultsParams) SetSourceID(sourceID *string) {
	o.SourceID = sourceID
}

// WriteToRequest writes these params to a swagger request
func (o *ListPacketCaptureResultsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.CaptureID != nil {

		// query param capture_id
		var qrCaptureID string

		if o.CaptureID != nil {
			qrCaptureID = *o.CaptureID
		}
		qCaptureID := qrCaptureID
		if qCaptureID != "" {

			if err := r.SetQueryParam("capture_id", qCaptureID); err != nil {
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
