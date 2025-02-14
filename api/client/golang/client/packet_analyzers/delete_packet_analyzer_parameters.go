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

// NewDeletePacketAnalyzerParams creates a new DeletePacketAnalyzerParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeletePacketAnalyzerParams() *DeletePacketAnalyzerParams {
	return &DeletePacketAnalyzerParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeletePacketAnalyzerParamsWithTimeout creates a new DeletePacketAnalyzerParams object
// with the ability to set a timeout on a request.
func NewDeletePacketAnalyzerParamsWithTimeout(timeout time.Duration) *DeletePacketAnalyzerParams {
	return &DeletePacketAnalyzerParams{
		timeout: timeout,
	}
}

// NewDeletePacketAnalyzerParamsWithContext creates a new DeletePacketAnalyzerParams object
// with the ability to set a context for a request.
func NewDeletePacketAnalyzerParamsWithContext(ctx context.Context) *DeletePacketAnalyzerParams {
	return &DeletePacketAnalyzerParams{
		Context: ctx,
	}
}

// NewDeletePacketAnalyzerParamsWithHTTPClient creates a new DeletePacketAnalyzerParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeletePacketAnalyzerParamsWithHTTPClient(client *http.Client) *DeletePacketAnalyzerParams {
	return &DeletePacketAnalyzerParams{
		HTTPClient: client,
	}
}

/* DeletePacketAnalyzerParams contains all the parameters to send to the API endpoint
   for the delete packet analyzer operation.

   Typically these are written to a http.Request.
*/
type DeletePacketAnalyzerParams struct {

	/* ID.

	   Unique resource identifier

	   Format: string
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete packet analyzer params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeletePacketAnalyzerParams) WithDefaults() *DeletePacketAnalyzerParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete packet analyzer params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeletePacketAnalyzerParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete packet analyzer params
func (o *DeletePacketAnalyzerParams) WithTimeout(timeout time.Duration) *DeletePacketAnalyzerParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete packet analyzer params
func (o *DeletePacketAnalyzerParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete packet analyzer params
func (o *DeletePacketAnalyzerParams) WithContext(ctx context.Context) *DeletePacketAnalyzerParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete packet analyzer params
func (o *DeletePacketAnalyzerParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete packet analyzer params
func (o *DeletePacketAnalyzerParams) WithHTTPClient(client *http.Client) *DeletePacketAnalyzerParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete packet analyzer params
func (o *DeletePacketAnalyzerParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the delete packet analyzer params
func (o *DeletePacketAnalyzerParams) WithID(id string) *DeletePacketAnalyzerParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete packet analyzer params
func (o *DeletePacketAnalyzerParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeletePacketAnalyzerParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
