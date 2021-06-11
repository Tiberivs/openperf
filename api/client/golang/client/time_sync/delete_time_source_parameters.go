// Code generated by go-swagger; DO NOT EDIT.

package time_sync

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

// NewDeleteTimeSourceParams creates a new DeleteTimeSourceParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteTimeSourceParams() *DeleteTimeSourceParams {
	return &DeleteTimeSourceParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteTimeSourceParamsWithTimeout creates a new DeleteTimeSourceParams object
// with the ability to set a timeout on a request.
func NewDeleteTimeSourceParamsWithTimeout(timeout time.Duration) *DeleteTimeSourceParams {
	return &DeleteTimeSourceParams{
		timeout: timeout,
	}
}

// NewDeleteTimeSourceParamsWithContext creates a new DeleteTimeSourceParams object
// with the ability to set a context for a request.
func NewDeleteTimeSourceParamsWithContext(ctx context.Context) *DeleteTimeSourceParams {
	return &DeleteTimeSourceParams{
		Context: ctx,
	}
}

// NewDeleteTimeSourceParamsWithHTTPClient creates a new DeleteTimeSourceParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteTimeSourceParamsWithHTTPClient(client *http.Client) *DeleteTimeSourceParams {
	return &DeleteTimeSourceParams{
		HTTPClient: client,
	}
}

/* DeleteTimeSourceParams contains all the parameters to send to the API endpoint
   for the delete time source operation.

   Typically these are written to a http.Request.
*/
type DeleteTimeSourceParams struct {

	/* ID.

	   Unique resource identifier

	   Format: string
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete time source params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteTimeSourceParams) WithDefaults() *DeleteTimeSourceParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete time source params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteTimeSourceParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete time source params
func (o *DeleteTimeSourceParams) WithTimeout(timeout time.Duration) *DeleteTimeSourceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete time source params
func (o *DeleteTimeSourceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete time source params
func (o *DeleteTimeSourceParams) WithContext(ctx context.Context) *DeleteTimeSourceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete time source params
func (o *DeleteTimeSourceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete time source params
func (o *DeleteTimeSourceParams) WithHTTPClient(client *http.Client) *DeleteTimeSourceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete time source params
func (o *DeleteTimeSourceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the delete time source params
func (o *DeleteTimeSourceParams) WithID(id string) *DeleteTimeSourceParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete time source params
func (o *DeleteTimeSourceParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteTimeSourceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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