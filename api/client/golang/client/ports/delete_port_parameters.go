// Code generated by go-swagger; DO NOT EDIT.

package ports

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

// NewDeletePortParams creates a new DeletePortParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeletePortParams() *DeletePortParams {
	return &DeletePortParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeletePortParamsWithTimeout creates a new DeletePortParams object
// with the ability to set a timeout on a request.
func NewDeletePortParamsWithTimeout(timeout time.Duration) *DeletePortParams {
	return &DeletePortParams{
		timeout: timeout,
	}
}

// NewDeletePortParamsWithContext creates a new DeletePortParams object
// with the ability to set a context for a request.
func NewDeletePortParamsWithContext(ctx context.Context) *DeletePortParams {
	return &DeletePortParams{
		Context: ctx,
	}
}

// NewDeletePortParamsWithHTTPClient creates a new DeletePortParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeletePortParamsWithHTTPClient(client *http.Client) *DeletePortParams {
	return &DeletePortParams{
		HTTPClient: client,
	}
}

/* DeletePortParams contains all the parameters to send to the API endpoint
   for the delete port operation.

   Typically these are written to a http.Request.
*/
type DeletePortParams struct {

	/* ID.

	   Unique resource identifier

	   Format: string
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete port params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeletePortParams) WithDefaults() *DeletePortParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete port params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeletePortParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete port params
func (o *DeletePortParams) WithTimeout(timeout time.Duration) *DeletePortParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete port params
func (o *DeletePortParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete port params
func (o *DeletePortParams) WithContext(ctx context.Context) *DeletePortParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete port params
func (o *DeletePortParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete port params
func (o *DeletePortParams) WithHTTPClient(client *http.Client) *DeletePortParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete port params
func (o *DeletePortParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the delete port params
func (o *DeletePortParams) WithID(id string) *DeletePortParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete port params
func (o *DeletePortParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeletePortParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
