// Code generated by go-swagger; DO NOT EDIT.

package learning

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

// NewGetPacketGeneratorLearningResultsParams creates a new GetPacketGeneratorLearningResultsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetPacketGeneratorLearningResultsParams() *GetPacketGeneratorLearningResultsParams {
	return &GetPacketGeneratorLearningResultsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetPacketGeneratorLearningResultsParamsWithTimeout creates a new GetPacketGeneratorLearningResultsParams object
// with the ability to set a timeout on a request.
func NewGetPacketGeneratorLearningResultsParamsWithTimeout(timeout time.Duration) *GetPacketGeneratorLearningResultsParams {
	return &GetPacketGeneratorLearningResultsParams{
		timeout: timeout,
	}
}

// NewGetPacketGeneratorLearningResultsParamsWithContext creates a new GetPacketGeneratorLearningResultsParams object
// with the ability to set a context for a request.
func NewGetPacketGeneratorLearningResultsParamsWithContext(ctx context.Context) *GetPacketGeneratorLearningResultsParams {
	return &GetPacketGeneratorLearningResultsParams{
		Context: ctx,
	}
}

// NewGetPacketGeneratorLearningResultsParamsWithHTTPClient creates a new GetPacketGeneratorLearningResultsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetPacketGeneratorLearningResultsParamsWithHTTPClient(client *http.Client) *GetPacketGeneratorLearningResultsParams {
	return &GetPacketGeneratorLearningResultsParams{
		HTTPClient: client,
	}
}

/* GetPacketGeneratorLearningResultsParams contains all the parameters to send to the API endpoint
   for the get packet generator learning results operation.

   Typically these are written to a http.Request.
*/
type GetPacketGeneratorLearningResultsParams struct {

	/* ID.

	   Unique resource identifier

	   Format: string
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get packet generator learning results params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetPacketGeneratorLearningResultsParams) WithDefaults() *GetPacketGeneratorLearningResultsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get packet generator learning results params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetPacketGeneratorLearningResultsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get packet generator learning results params
func (o *GetPacketGeneratorLearningResultsParams) WithTimeout(timeout time.Duration) *GetPacketGeneratorLearningResultsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get packet generator learning results params
func (o *GetPacketGeneratorLearningResultsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get packet generator learning results params
func (o *GetPacketGeneratorLearningResultsParams) WithContext(ctx context.Context) *GetPacketGeneratorLearningResultsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get packet generator learning results params
func (o *GetPacketGeneratorLearningResultsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get packet generator learning results params
func (o *GetPacketGeneratorLearningResultsParams) WithHTTPClient(client *http.Client) *GetPacketGeneratorLearningResultsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get packet generator learning results params
func (o *GetPacketGeneratorLearningResultsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get packet generator learning results params
func (o *GetPacketGeneratorLearningResultsParams) WithID(id string) *GetPacketGeneratorLearningResultsParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get packet generator learning results params
func (o *GetPacketGeneratorLearningResultsParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetPacketGeneratorLearningResultsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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