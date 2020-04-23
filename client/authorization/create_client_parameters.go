// Code generated by go-swagger; DO NOT EDIT.

package authorization

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

// NewCreateClientParams creates a new CreateClientParams object
// with the default values initialized.
func NewCreateClientParams() *CreateClientParams {
	var ()
	return &CreateClientParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateClientParamsWithTimeout creates a new CreateClientParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateClientParamsWithTimeout(timeout time.Duration) *CreateClientParams {
	var ()
	return &CreateClientParams{

		timeout: timeout,
	}
}

// NewCreateClientParamsWithContext creates a new CreateClientParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateClientParamsWithContext(ctx context.Context) *CreateClientParams {
	var ()
	return &CreateClientParams{

		Context: ctx,
	}
}

// NewCreateClientParamsWithHTTPClient creates a new CreateClientParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateClientParamsWithHTTPClient(client *http.Client) *CreateClientParams {
	var ()
	return &CreateClientParams{
		HTTPClient: client,
	}
}

/*CreateClientParams contains all the parameters to send to the API endpoint
for the create client operation typically these are written to a http.Request
*/
type CreateClientParams struct {

	/*Client*/
	Client CreateClientBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create client params
func (o *CreateClientParams) WithTimeout(timeout time.Duration) *CreateClientParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create client params
func (o *CreateClientParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create client params
func (o *CreateClientParams) WithContext(ctx context.Context) *CreateClientParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create client params
func (o *CreateClientParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create client params
func (o *CreateClientParams) WithHTTPClient(client *http.Client) *CreateClientParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create client params
func (o *CreateClientParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClient adds the client to the create client params
func (o *CreateClientParams) WithClient(client CreateClientBody) *CreateClientParams {
	o.SetClient(client)
	return o
}

// SetClient adds the client to the create client params
func (o *CreateClientParams) SetClient(client CreateClientBody) {
	o.Client = client
}

// WriteToRequest writes these params to a swagger request
func (o *CreateClientParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if err := r.SetBodyParam(o.Client); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
