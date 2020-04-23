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
)

// NewGetSmsCallerIdsParams creates a new GetSmsCallerIdsParams object
// with the default values initialized.
func NewGetSmsCallerIdsParams() *GetSmsCallerIdsParams {
	var ()
	return &GetSmsCallerIdsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetSmsCallerIdsParamsWithTimeout creates a new GetSmsCallerIdsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetSmsCallerIdsParamsWithTimeout(timeout time.Duration) *GetSmsCallerIdsParams {
	var ()
	return &GetSmsCallerIdsParams{

		timeout: timeout,
	}
}

// NewGetSmsCallerIdsParamsWithContext creates a new GetSmsCallerIdsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetSmsCallerIdsParamsWithContext(ctx context.Context) *GetSmsCallerIdsParams {
	var ()
	return &GetSmsCallerIdsParams{

		Context: ctx,
	}
}

// NewGetSmsCallerIdsParamsWithHTTPClient creates a new GetSmsCallerIdsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetSmsCallerIdsParamsWithHTTPClient(client *http.Client) *GetSmsCallerIdsParams {
	var ()
	return &GetSmsCallerIdsParams{
		HTTPClient: client,
	}
}

/*GetSmsCallerIdsParams contains all the parameters to send to the API endpoint
for the get sms caller ids operation typically these are written to a http.Request
*/
type GetSmsCallerIdsParams struct {

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

// WithTimeout adds the timeout to the get sms caller ids params
func (o *GetSmsCallerIdsParams) WithTimeout(timeout time.Duration) *GetSmsCallerIdsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get sms caller ids params
func (o *GetSmsCallerIdsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get sms caller ids params
func (o *GetSmsCallerIdsParams) WithContext(ctx context.Context) *GetSmsCallerIdsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get sms caller ids params
func (o *GetSmsCallerIdsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get sms caller ids params
func (o *GetSmsCallerIdsParams) WithHTTPClient(client *http.Client) *GetSmsCallerIdsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get sms caller ids params
func (o *GetSmsCallerIdsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSmsID adds the smsID to the get sms caller ids params
func (o *GetSmsCallerIdsParams) WithSmsID(smsID string) *GetSmsCallerIdsParams {
	o.SetSmsID(smsID)
	return o
}

// SetSmsID adds the smsId to the get sms caller ids params
func (o *GetSmsCallerIdsParams) SetSmsID(smsID string) {
	o.SmsID = smsID
}

// WithUserID adds the userID to the get sms caller ids params
func (o *GetSmsCallerIdsParams) WithUserID(userID string) *GetSmsCallerIdsParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the get sms caller ids params
func (o *GetSmsCallerIdsParams) SetUserID(userID string) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *GetSmsCallerIdsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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