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
)

// NewDeleteNotificationParams creates a new DeleteNotificationParams object
// with the default values initialized.
func NewDeleteNotificationParams() *DeleteNotificationParams {
	var ()
	return &DeleteNotificationParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteNotificationParamsWithTimeout creates a new DeleteNotificationParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteNotificationParamsWithTimeout(timeout time.Duration) *DeleteNotificationParams {
	var ()
	return &DeleteNotificationParams{

		timeout: timeout,
	}
}

// NewDeleteNotificationParamsWithContext creates a new DeleteNotificationParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteNotificationParamsWithContext(ctx context.Context) *DeleteNotificationParams {
	var ()
	return &DeleteNotificationParams{

		Context: ctx,
	}
}

// NewDeleteNotificationParamsWithHTTPClient creates a new DeleteNotificationParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteNotificationParamsWithHTTPClient(client *http.Client) *DeleteNotificationParams {
	var ()
	return &DeleteNotificationParams{
		HTTPClient: client,
	}
}

/*DeleteNotificationParams contains all the parameters to send to the API endpoint
for the delete notification operation typically these are written to a http.Request
*/
type DeleteNotificationParams struct {

	/*NotificationID
	  The unique notification identifier

	*/
	NotificationID string
	/*UserID
	  The unique user identifier

	*/
	UserID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete notification params
func (o *DeleteNotificationParams) WithTimeout(timeout time.Duration) *DeleteNotificationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete notification params
func (o *DeleteNotificationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete notification params
func (o *DeleteNotificationParams) WithContext(ctx context.Context) *DeleteNotificationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete notification params
func (o *DeleteNotificationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete notification params
func (o *DeleteNotificationParams) WithHTTPClient(client *http.Client) *DeleteNotificationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete notification params
func (o *DeleteNotificationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNotificationID adds the notificationID to the delete notification params
func (o *DeleteNotificationParams) WithNotificationID(notificationID string) *DeleteNotificationParams {
	o.SetNotificationID(notificationID)
	return o
}

// SetNotificationID adds the notificationId to the delete notification params
func (o *DeleteNotificationParams) SetNotificationID(notificationID string) {
	o.NotificationID = notificationID
}

// WithUserID adds the userID to the delete notification params
func (o *DeleteNotificationParams) WithUserID(userID string) *DeleteNotificationParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the delete notification params
func (o *DeleteNotificationParams) SetUserID(userID string) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteNotificationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param notificationId
	if err := r.SetPathParam("notificationId", o.NotificationID); err != nil {
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
