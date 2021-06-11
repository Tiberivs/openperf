// Code generated by go-swagger; DO NOT EDIT.

package block_generator

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

// NewBulkDeleteBlockGeneratorsParams creates a new BulkDeleteBlockGeneratorsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewBulkDeleteBlockGeneratorsParams() *BulkDeleteBlockGeneratorsParams {
	return &BulkDeleteBlockGeneratorsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewBulkDeleteBlockGeneratorsParamsWithTimeout creates a new BulkDeleteBlockGeneratorsParams object
// with the ability to set a timeout on a request.
func NewBulkDeleteBlockGeneratorsParamsWithTimeout(timeout time.Duration) *BulkDeleteBlockGeneratorsParams {
	return &BulkDeleteBlockGeneratorsParams{
		timeout: timeout,
	}
}

// NewBulkDeleteBlockGeneratorsParamsWithContext creates a new BulkDeleteBlockGeneratorsParams object
// with the ability to set a context for a request.
func NewBulkDeleteBlockGeneratorsParamsWithContext(ctx context.Context) *BulkDeleteBlockGeneratorsParams {
	return &BulkDeleteBlockGeneratorsParams{
		Context: ctx,
	}
}

// NewBulkDeleteBlockGeneratorsParamsWithHTTPClient creates a new BulkDeleteBlockGeneratorsParams object
// with the ability to set a custom HTTPClient for a request.
func NewBulkDeleteBlockGeneratorsParamsWithHTTPClient(client *http.Client) *BulkDeleteBlockGeneratorsParams {
	return &BulkDeleteBlockGeneratorsParams{
		HTTPClient: client,
	}
}

/* BulkDeleteBlockGeneratorsParams contains all the parameters to send to the API endpoint
   for the bulk delete block generators operation.

   Typically these are written to a http.Request.
*/
type BulkDeleteBlockGeneratorsParams struct {

	/* Delete.

	   Bulk delete
	*/
	Delete BulkDeleteBlockGeneratorsBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the bulk delete block generators params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *BulkDeleteBlockGeneratorsParams) WithDefaults() *BulkDeleteBlockGeneratorsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the bulk delete block generators params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *BulkDeleteBlockGeneratorsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the bulk delete block generators params
func (o *BulkDeleteBlockGeneratorsParams) WithTimeout(timeout time.Duration) *BulkDeleteBlockGeneratorsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the bulk delete block generators params
func (o *BulkDeleteBlockGeneratorsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the bulk delete block generators params
func (o *BulkDeleteBlockGeneratorsParams) WithContext(ctx context.Context) *BulkDeleteBlockGeneratorsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the bulk delete block generators params
func (o *BulkDeleteBlockGeneratorsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the bulk delete block generators params
func (o *BulkDeleteBlockGeneratorsParams) WithHTTPClient(client *http.Client) *BulkDeleteBlockGeneratorsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the bulk delete block generators params
func (o *BulkDeleteBlockGeneratorsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDelete adds the delete to the bulk delete block generators params
func (o *BulkDeleteBlockGeneratorsParams) WithDelete(delete BulkDeleteBlockGeneratorsBody) *BulkDeleteBlockGeneratorsParams {
	o.SetDelete(delete)
	return o
}

// SetDelete adds the delete to the bulk delete block generators params
func (o *BulkDeleteBlockGeneratorsParams) SetDelete(delete BulkDeleteBlockGeneratorsBody) {
	o.Delete = delete
}

// WriteToRequest writes these params to a swagger request
func (o *BulkDeleteBlockGeneratorsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if err := r.SetBodyParam(o.Delete); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}