// Code generated by go-swagger; DO NOT EDIT.

package notifications

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

// NewAddFaxSmsNotificationParams creates a new AddFaxSmsNotificationParams object
// with the default values initialized.
func NewAddFaxSmsNotificationParams() *AddFaxSmsNotificationParams {
	var ()
	return &AddFaxSmsNotificationParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAddFaxSmsNotificationParamsWithTimeout creates a new AddFaxSmsNotificationParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAddFaxSmsNotificationParamsWithTimeout(timeout time.Duration) *AddFaxSmsNotificationParams {
	var ()
	return &AddFaxSmsNotificationParams{

		timeout: timeout,
	}
}

// NewAddFaxSmsNotificationParamsWithContext creates a new AddFaxSmsNotificationParams object
// with the default values initialized, and the ability to set a context for a request
func NewAddFaxSmsNotificationParamsWithContext(ctx context.Context) *AddFaxSmsNotificationParams {
	var ()
	return &AddFaxSmsNotificationParams{

		Context: ctx,
	}
}

// NewAddFaxSmsNotificationParamsWithHTTPClient creates a new AddFaxSmsNotificationParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAddFaxSmsNotificationParamsWithHTTPClient(client *http.Client) *AddFaxSmsNotificationParams {
	var ()
	return &AddFaxSmsNotificationParams{
		HTTPClient: client,
	}
}

/*AddFaxSmsNotificationParams contains all the parameters to send to the API endpoint
for the add fax sms notification operation typically these are written to a http.Request
*/
type AddFaxSmsNotificationParams struct {

	/*Body*/
	Body *models.AddFaxSmsNotificationRequest
	/*UserID
	  The unique user identifier

	*/
	UserID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the add fax sms notification params
func (o *AddFaxSmsNotificationParams) WithTimeout(timeout time.Duration) *AddFaxSmsNotificationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add fax sms notification params
func (o *AddFaxSmsNotificationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add fax sms notification params
func (o *AddFaxSmsNotificationParams) WithContext(ctx context.Context) *AddFaxSmsNotificationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add fax sms notification params
func (o *AddFaxSmsNotificationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add fax sms notification params
func (o *AddFaxSmsNotificationParams) WithHTTPClient(client *http.Client) *AddFaxSmsNotificationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add fax sms notification params
func (o *AddFaxSmsNotificationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the add fax sms notification params
func (o *AddFaxSmsNotificationParams) WithBody(body *models.AddFaxSmsNotificationRequest) *AddFaxSmsNotificationParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the add fax sms notification params
func (o *AddFaxSmsNotificationParams) SetBody(body *models.AddFaxSmsNotificationRequest) {
	o.Body = body
}

// WithUserID adds the userID to the add fax sms notification params
func (o *AddFaxSmsNotificationParams) WithUserID(userID string) *AddFaxSmsNotificationParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the add fax sms notification params
func (o *AddFaxSmsNotificationParams) SetUserID(userID string) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *AddFaxSmsNotificationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
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
