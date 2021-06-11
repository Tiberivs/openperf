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

// NewGetPacketCaptureResultParams creates a new GetPacketCaptureResultParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetPacketCaptureResultParams() *GetPacketCaptureResultParams {
	return &GetPacketCaptureResultParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetPacketCaptureResultParamsWithTimeout creates a new GetPacketCaptureResultParams object
// with the ability to set a timeout on a request.
func NewGetPacketCaptureResultParamsWithTimeout(timeout time.Duration) *GetPacketCaptureResultParams {
	return &GetPacketCaptureResultParams{
		timeout: timeout,
	}
}

// NewGetPacketCaptureResultParamsWithContext creates a new GetPacketCaptureResultParams object
// with the ability to set a context for a request.
func NewGetPacketCaptureResultParamsWithContext(ctx context.Context) *GetPacketCaptureResultParams {
	return &GetPacketCaptureResultParams{
		Context: ctx,
	}
}

// NewGetPacketCaptureResultParamsWithHTTPClient creates a new GetPacketCaptureResultParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetPacketCaptureResultParamsWithHTTPClient(client *http.Client) *GetPacketCaptureResultParams {
	return &GetPacketCaptureResultParams{
		HTTPClient: client,
	}
}

/* GetPacketCaptureResultParams contains all the parameters to send to the API endpoint
   for the get packet capture result operation.

   Typically these are written to a http.Request.
*/
type GetPacketCaptureResultParams struct {

	/* ID.

	   Unique resource identifier

	   Format: string
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get packet capture result params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetPacketCaptureResultParams) WithDefaults() *GetPacketCaptureResultParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get packet capture result params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetPacketCaptureResultParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get packet capture result params
func (o *GetPacketCaptureResultParams) WithTimeout(timeout time.Duration) *GetPacketCaptureResultParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get packet capture result params
func (o *GetPacketCaptureResultParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get packet capture result params
func (o *GetPacketCaptureResultParams) WithContext(ctx context.Context) *GetPacketCaptureResultParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get packet capture result params
func (o *GetPacketCaptureResultParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get packet capture result params
func (o *GetPacketCaptureResultParams) WithHTTPClient(client *http.Client) *GetPacketCaptureResultParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get packet capture result params
func (o *GetPacketCaptureResultParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get packet capture result params
func (o *GetPacketCaptureResultParams) WithID(id string) *GetPacketCaptureResultParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get packet capture result params
func (o *GetPacketCaptureResultParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetPacketCaptureResultParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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