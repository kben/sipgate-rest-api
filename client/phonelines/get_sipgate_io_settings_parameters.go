// Code generated by go-swagger; DO NOT EDIT.

package phonelines

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

// NewGetSipgateIoSettingsParams creates a new GetSipgateIoSettingsParams object
// with the default values initialized.
func NewGetSipgateIoSettingsParams() *GetSipgateIoSettingsParams {
	var ()
	return &GetSipgateIoSettingsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetSipgateIoSettingsParamsWithTimeout creates a new GetSipgateIoSettingsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetSipgateIoSettingsParamsWithTimeout(timeout time.Duration) *GetSipgateIoSettingsParams {
	var ()
	return &GetSipgateIoSettingsParams{

		timeout: timeout,
	}
}

// NewGetSipgateIoSettingsParamsWithContext creates a new GetSipgateIoSettingsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetSipgateIoSettingsParamsWithContext(ctx context.Context) *GetSipgateIoSettingsParams {
	var ()
	return &GetSipgateIoSettingsParams{

		Context: ctx,
	}
}

// NewGetSipgateIoSettingsParamsWithHTTPClient creates a new GetSipgateIoSettingsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetSipgateIoSettingsParamsWithHTTPClient(client *http.Client) *GetSipgateIoSettingsParams {
	var ()
	return &GetSipgateIoSettingsParams{
		HTTPClient: client,
	}
}

/*GetSipgateIoSettingsParams contains all the parameters to send to the API endpoint
for the get sipgate io settings operation typically these are written to a http.Request
*/
type GetSipgateIoSettingsParams struct {

	/*PhonelineID
	  The unique phone line identifier

	*/
	PhonelineID string
	/*UserID
	  The unique user identifier

	*/
	UserID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get sipgate io settings params
func (o *GetSipgateIoSettingsParams) WithTimeout(timeout time.Duration) *GetSipgateIoSettingsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get sipgate io settings params
func (o *GetSipgateIoSettingsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get sipgate io settings params
func (o *GetSipgateIoSettingsParams) WithContext(ctx context.Context) *GetSipgateIoSettingsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get sipgate io settings params
func (o *GetSipgateIoSettingsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get sipgate io settings params
func (o *GetSipgateIoSettingsParams) WithHTTPClient(client *http.Client) *GetSipgateIoSettingsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get sipgate io settings params
func (o *GetSipgateIoSettingsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPhonelineID adds the phonelineID to the get sipgate io settings params
func (o *GetSipgateIoSettingsParams) WithPhonelineID(phonelineID string) *GetSipgateIoSettingsParams {
	o.SetPhonelineID(phonelineID)
	return o
}

// SetPhonelineID adds the phonelineId to the get sipgate io settings params
func (o *GetSipgateIoSettingsParams) SetPhonelineID(phonelineID string) {
	o.PhonelineID = phonelineID
}

// WithUserID adds the userID to the get sipgate io settings params
func (o *GetSipgateIoSettingsParams) WithUserID(userID string) *GetSipgateIoSettingsParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the get sipgate io settings params
func (o *GetSipgateIoSettingsParams) SetUserID(userID string) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *GetSipgateIoSettingsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param phonelineId
	if err := r.SetPathParam("phonelineId", o.PhonelineID); err != nil {
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
