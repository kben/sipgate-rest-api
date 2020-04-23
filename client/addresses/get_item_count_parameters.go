// Code generated by go-swagger; DO NOT EDIT.

package addresses

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

// NewGetItemCountParams creates a new GetItemCountParams object
// with the default values initialized.
func NewGetItemCountParams() *GetItemCountParams {

	return &GetItemCountParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetItemCountParamsWithTimeout creates a new GetItemCountParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetItemCountParamsWithTimeout(timeout time.Duration) *GetItemCountParams {

	return &GetItemCountParams{

		timeout: timeout,
	}
}

// NewGetItemCountParamsWithContext creates a new GetItemCountParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetItemCountParamsWithContext(ctx context.Context) *GetItemCountParams {

	return &GetItemCountParams{

		Context: ctx,
	}
}

// NewGetItemCountParamsWithHTTPClient creates a new GetItemCountParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetItemCountParamsWithHTTPClient(client *http.Client) *GetItemCountParams {

	return &GetItemCountParams{
		HTTPClient: client,
	}
}

/*GetItemCountParams contains all the parameters to send to the API endpoint
for the get item count operation typically these are written to a http.Request
*/
type GetItemCountParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get item count params
func (o *GetItemCountParams) WithTimeout(timeout time.Duration) *GetItemCountParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get item count params
func (o *GetItemCountParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get item count params
func (o *GetItemCountParams) WithContext(ctx context.Context) *GetItemCountParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get item count params
func (o *GetItemCountParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get item count params
func (o *GetItemCountParams) WithHTTPClient(client *http.Client) *GetItemCountParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get item count params
func (o *GetItemCountParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetItemCountParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
