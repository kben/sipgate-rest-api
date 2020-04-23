// Code generated by go-swagger; DO NOT EDIT.

package sms

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

// NewAddCallerIDParams creates a new AddCallerIDParams object
// with the default values initialized.
func NewAddCallerIDParams() *AddCallerIDParams {
	var ()
	return &AddCallerIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAddCallerIDParamsWithTimeout creates a new AddCallerIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAddCallerIDParamsWithTimeout(timeout time.Duration) *AddCallerIDParams {
	var ()
	return &AddCallerIDParams{

		timeout: timeout,
	}
}

// NewAddCallerIDParamsWithContext creates a new AddCallerIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewAddCallerIDParamsWithContext(ctx context.Context) *AddCallerIDParams {
	var ()
	return &AddCallerIDParams{

		Context: ctx,
	}
}

// NewAddCallerIDParamsWithHTTPClient creates a new AddCallerIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAddCallerIDParamsWithHTTPClient(client *http.Client) *AddCallerIDParams {
	var ()
	return &AddCallerIDParams{
		HTTPClient: client,
	}
}

/*AddCallerIDParams contains all the parameters to send to the API endpoint
for the add caller Id operation typically these are written to a http.Request
*/
type AddCallerIDParams struct {

	/*Body*/
	Body *models.SmsCallerIDRequest
	/*SmsID
	  The unique short message service identifier

	*/
	SmsID string
	/*UserID
	  The unique user identifier

	*/
	UserID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the add caller Id params
func (o *AddCallerIDParams) WithTimeout(timeout time.Duration) *AddCallerIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add caller Id params
func (o *AddCallerIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add caller Id params
func (o *AddCallerIDParams) WithContext(ctx context.Context) *AddCallerIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add caller Id params
func (o *AddCallerIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add caller Id params
func (o *AddCallerIDParams) WithHTTPClient(client *http.Client) *AddCallerIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add caller Id params
func (o *AddCallerIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the add caller Id params
func (o *AddCallerIDParams) WithBody(body *models.SmsCallerIDRequest) *AddCallerIDParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the add caller Id params
func (o *AddCallerIDParams) SetBody(body *models.SmsCallerIDRequest) {
	o.Body = body
}

// WithSmsID adds the smsID to the add caller Id params
func (o *AddCallerIDParams) WithSmsID(smsID string) *AddCallerIDParams {
	o.SetSmsID(smsID)
	return o
}

// SetSmsID adds the smsId to the add caller Id params
func (o *AddCallerIDParams) SetSmsID(smsID string) {
	o.SmsID = smsID
}

// WithUserID adds the userID to the add caller Id params
func (o *AddCallerIDParams) WithUserID(userID string) *AddCallerIDParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the add caller Id params
func (o *AddCallerIDParams) SetUserID(userID string) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *AddCallerIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param smsId
	if err := r.SetPathParam("smsId", o.SmsID); err != nil {
		return err
	}

	// path param userId
	if err := r.SetPathParam("userId", o.UserID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
