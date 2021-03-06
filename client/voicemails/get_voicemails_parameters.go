// Code generated by go-swagger; DO NOT EDIT.

package voicemails

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

// NewGetVoicemailsParams creates a new GetVoicemailsParams object
// with the default values initialized.
func NewGetVoicemailsParams() *GetVoicemailsParams {

	return &GetVoicemailsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetVoicemailsParamsWithTimeout creates a new GetVoicemailsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetVoicemailsParamsWithTimeout(timeout time.Duration) *GetVoicemailsParams {

	return &GetVoicemailsParams{

		timeout: timeout,
	}
}

// NewGetVoicemailsParamsWithContext creates a new GetVoicemailsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetVoicemailsParamsWithContext(ctx context.Context) *GetVoicemailsParams {

	return &GetVoicemailsParams{

		Context: ctx,
	}
}

// NewGetVoicemailsParamsWithHTTPClient creates a new GetVoicemailsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetVoicemailsParamsWithHTTPClient(client *http.Client) *GetVoicemailsParams {

	return &GetVoicemailsParams{
		HTTPClient: client,
	}
}

/*GetVoicemailsParams contains all the parameters to send to the API endpoint
for the get voicemails operation typically these are written to a http.Request
*/
type GetVoicemailsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get voicemails params
func (o *GetVoicemailsParams) WithTimeout(timeout time.Duration) *GetVoicemailsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get voicemails params
func (o *GetVoicemailsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get voicemails params
func (o *GetVoicemailsParams) WithContext(ctx context.Context) *GetVoicemailsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get voicemails params
func (o *GetVoicemailsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get voicemails params
func (o *GetVoicemailsParams) WithHTTPClient(client *http.Client) *GetVoicemailsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get voicemails params
func (o *GetVoicemailsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetVoicemailsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
