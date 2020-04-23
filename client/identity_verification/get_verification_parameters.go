// Code generated by go-swagger; DO NOT EDIT.

package identity_verification

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

// NewGetVerificationParams creates a new GetVerificationParams object
// with the default values initialized.
func NewGetVerificationParams() *GetVerificationParams {

	return &GetVerificationParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetVerificationParamsWithTimeout creates a new GetVerificationParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetVerificationParamsWithTimeout(timeout time.Duration) *GetVerificationParams {

	return &GetVerificationParams{

		timeout: timeout,
	}
}

// NewGetVerificationParamsWithContext creates a new GetVerificationParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetVerificationParamsWithContext(ctx context.Context) *GetVerificationParams {

	return &GetVerificationParams{

		Context: ctx,
	}
}

// NewGetVerificationParamsWithHTTPClient creates a new GetVerificationParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetVerificationParamsWithHTTPClient(client *http.Client) *GetVerificationParams {

	return &GetVerificationParams{
		HTTPClient: client,
	}
}

/*GetVerificationParams contains all the parameters to send to the API endpoint
for the get verification operation typically these are written to a http.Request
*/
type GetVerificationParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get verification params
func (o *GetVerificationParams) WithTimeout(timeout time.Duration) *GetVerificationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get verification params
func (o *GetVerificationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get verification params
func (o *GetVerificationParams) WithContext(ctx context.Context) *GetVerificationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get verification params
func (o *GetVerificationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get verification params
func (o *GetVerificationParams) WithHTTPClient(client *http.Client) *GetVerificationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get verification params
func (o *GetVerificationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetVerificationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
