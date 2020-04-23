// Code generated by go-swagger; DO NOT EDIT.

package contacts

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

// NewGetAllContactsParams creates a new GetAllContactsParams object
// with the default values initialized.
func NewGetAllContactsParams() *GetAllContactsParams {

	return &GetAllContactsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetAllContactsParamsWithTimeout creates a new GetAllContactsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetAllContactsParamsWithTimeout(timeout time.Duration) *GetAllContactsParams {

	return &GetAllContactsParams{

		timeout: timeout,
	}
}

// NewGetAllContactsParamsWithContext creates a new GetAllContactsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetAllContactsParamsWithContext(ctx context.Context) *GetAllContactsParams {

	return &GetAllContactsParams{

		Context: ctx,
	}
}

// NewGetAllContactsParamsWithHTTPClient creates a new GetAllContactsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetAllContactsParamsWithHTTPClient(client *http.Client) *GetAllContactsParams {

	return &GetAllContactsParams{
		HTTPClient: client,
	}
}

/*GetAllContactsParams contains all the parameters to send to the API endpoint
for the get all contacts operation typically these are written to a http.Request
*/
type GetAllContactsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get all contacts params
func (o *GetAllContactsParams) WithTimeout(timeout time.Duration) *GetAllContactsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get all contacts params
func (o *GetAllContactsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get all contacts params
func (o *GetAllContactsParams) WithContext(ctx context.Context) *GetAllContactsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get all contacts params
func (o *GetAllContactsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get all contacts params
func (o *GetAllContactsParams) WithHTTPClient(client *http.Client) *GetAllContactsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get all contacts params
func (o *GetAllContactsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetAllContactsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}