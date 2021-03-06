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

// NewGetAddressesParams creates a new GetAddressesParams object
// with the default values initialized.
func NewGetAddressesParams() *GetAddressesParams {

	return &GetAddressesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetAddressesParamsWithTimeout creates a new GetAddressesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetAddressesParamsWithTimeout(timeout time.Duration) *GetAddressesParams {

	return &GetAddressesParams{

		timeout: timeout,
	}
}

// NewGetAddressesParamsWithContext creates a new GetAddressesParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetAddressesParamsWithContext(ctx context.Context) *GetAddressesParams {

	return &GetAddressesParams{

		Context: ctx,
	}
}

// NewGetAddressesParamsWithHTTPClient creates a new GetAddressesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetAddressesParamsWithHTTPClient(client *http.Client) *GetAddressesParams {

	return &GetAddressesParams{
		HTTPClient: client,
	}
}

/*GetAddressesParams contains all the parameters to send to the API endpoint
for the get addresses operation typically these are written to a http.Request
*/
type GetAddressesParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get addresses params
func (o *GetAddressesParams) WithTimeout(timeout time.Duration) *GetAddressesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get addresses params
func (o *GetAddressesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get addresses params
func (o *GetAddressesParams) WithContext(ctx context.Context) *GetAddressesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get addresses params
func (o *GetAddressesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get addresses params
func (o *GetAddressesParams) WithHTTPClient(client *http.Client) *GetAddressesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get addresses params
func (o *GetAddressesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetAddressesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
