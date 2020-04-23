// Code generated by go-swagger; DO NOT EDIT.

package sessions

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

	"github.com/kben/sipgate-rest-api/models"
)

// NewSendFaxParams creates a new SendFaxParams object
// with the default values initialized.
func NewSendFaxParams() *SendFaxParams {
	var ()
	return &SendFaxParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSendFaxParamsWithTimeout creates a new SendFaxParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSendFaxParamsWithTimeout(timeout time.Duration) *SendFaxParams {
	var ()
	return &SendFaxParams{

		timeout: timeout,
	}
}

// NewSendFaxParamsWithContext creates a new SendFaxParams object
// with the default values initialized, and the ability to set a context for a request
func NewSendFaxParamsWithContext(ctx context.Context) *SendFaxParams {
	var ()
	return &SendFaxParams{

		Context: ctx,
	}
}

// NewSendFaxParamsWithHTTPClient creates a new SendFaxParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSendFaxParamsWithHTTPClient(client *http.Client) *SendFaxParams {
	var ()
	return &SendFaxParams{
		HTTPClient: client,
	}
}

/*SendFaxParams contains all the parameters to send to the API endpoint
for the send fax operation typically these are written to a http.Request
*/
type SendFaxParams struct {

	/*Body*/
	Body *models.SendFaxRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the send fax params
func (o *SendFaxParams) WithTimeout(timeout time.Duration) *SendFaxParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the send fax params
func (o *SendFaxParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the send fax params
func (o *SendFaxParams) WithContext(ctx context.Context) *SendFaxParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the send fax params
func (o *SendFaxParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the send fax params
func (o *SendFaxParams) WithHTTPClient(client *http.Client) *SendFaxParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the send fax params
func (o *SendFaxParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the send fax params
func (o *SendFaxParams) WithBody(body *models.SendFaxRequest) *SendFaxParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the send fax params
func (o *SendFaxParams) SetBody(body *models.SendFaxRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *SendFaxParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
