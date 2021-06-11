// Code generated by go-swagger; DO NOT EDIT.

package packet_generators

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

	"github.com/spirent/openperf/api/client/golang/models"
)

// NewCreatePacketGeneratorParams creates a new CreatePacketGeneratorParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreatePacketGeneratorParams() *CreatePacketGeneratorParams {
	return &CreatePacketGeneratorParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreatePacketGeneratorParamsWithTimeout creates a new CreatePacketGeneratorParams object
// with the ability to set a timeout on a request.
func NewCreatePacketGeneratorParamsWithTimeout(timeout time.Duration) *CreatePacketGeneratorParams {
	return &CreatePacketGeneratorParams{
		timeout: timeout,
	}
}

// NewCreatePacketGeneratorParamsWithContext creates a new CreatePacketGeneratorParams object
// with the ability to set a context for a request.
func NewCreatePacketGeneratorParamsWithContext(ctx context.Context) *CreatePacketGeneratorParams {
	return &CreatePacketGeneratorParams{
		Context: ctx,
	}
}

// NewCreatePacketGeneratorParamsWithHTTPClient creates a new CreatePacketGeneratorParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreatePacketGeneratorParamsWithHTTPClient(client *http.Client) *CreatePacketGeneratorParams {
	return &CreatePacketGeneratorParams{
		HTTPClient: client,
	}
}

/* CreatePacketGeneratorParams contains all the parameters to send to the API endpoint
   for the create packet generator operation.

   Typically these are written to a http.Request.
*/
type CreatePacketGeneratorParams struct {

	/* Generator.

	   New packet generator
	*/
	Generator *models.PacketGenerator

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create packet generator params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreatePacketGeneratorParams) WithDefaults() *CreatePacketGeneratorParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create packet generator params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreatePacketGeneratorParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create packet generator params
func (o *CreatePacketGeneratorParams) WithTimeout(timeout time.Duration) *CreatePacketGeneratorParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create packet generator params
func (o *CreatePacketGeneratorParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create packet generator params
func (o *CreatePacketGeneratorParams) WithContext(ctx context.Context) *CreatePacketGeneratorParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create packet generator params
func (o *CreatePacketGeneratorParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create packet generator params
func (o *CreatePacketGeneratorParams) WithHTTPClient(client *http.Client) *CreatePacketGeneratorParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create packet generator params
func (o *CreatePacketGeneratorParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithGenerator adds the generator to the create packet generator params
func (o *CreatePacketGeneratorParams) WithGenerator(generator *models.PacketGenerator) *CreatePacketGeneratorParams {
	o.SetGenerator(generator)
	return o
}

// SetGenerator adds the generator to the create packet generator params
func (o *CreatePacketGeneratorParams) SetGenerator(generator *models.PacketGenerator) {
	o.Generator = generator
}

// WriteToRequest writes these params to a swagger request
func (o *CreatePacketGeneratorParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Generator != nil {
		if err := r.SetBodyParam(o.Generator); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}