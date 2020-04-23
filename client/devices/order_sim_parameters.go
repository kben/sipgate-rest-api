// Code generated by go-swagger; DO NOT EDIT.

package devices

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

// NewOrderSimParams creates a new OrderSimParams object
// with the default values initialized.
func NewOrderSimParams() *OrderSimParams {
	var ()
	return &OrderSimParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewOrderSimParamsWithTimeout creates a new OrderSimParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewOrderSimParamsWithTimeout(timeout time.Duration) *OrderSimParams {
	var ()
	return &OrderSimParams{

		timeout: timeout,
	}
}

// NewOrderSimParamsWithContext creates a new OrderSimParams object
// with the default values initialized, and the ability to set a context for a request
func NewOrderSimParamsWithContext(ctx context.Context) *OrderSimParams {
	var ()
	return &OrderSimParams{

		Context: ctx,
	}
}

// NewOrderSimParamsWithHTTPClient creates a new OrderSimParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewOrderSimParamsWithHTTPClient(client *http.Client) *OrderSimParams {
	var ()
	return &OrderSimParams{
		HTTPClient: client,
	}
}

/*OrderSimParams contains all the parameters to send to the API endpoint
for the order sim operation typically these are written to a http.Request
*/
type OrderSimParams struct {

	/*Body*/
	Body *models.OrderSimRequest
	/*DeviceID
	  The unique mobile device identifier

	*/
	DeviceID string
	/*UserID
	  The unique user identifier

	*/
	UserID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the order sim params
func (o *OrderSimParams) WithTimeout(timeout time.Duration) *OrderSimParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the order sim params
func (o *OrderSimParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the order sim params
func (o *OrderSimParams) WithContext(ctx context.Context) *OrderSimParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the order sim params
func (o *OrderSimParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the order sim params
func (o *OrderSimParams) WithHTTPClient(client *http.Client) *OrderSimParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the order sim params
func (o *OrderSimParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the order sim params
func (o *OrderSimParams) WithBody(body *models.OrderSimRequest) *OrderSimParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the order sim params
func (o *OrderSimParams) SetBody(body *models.OrderSimRequest) {
	o.Body = body
}

// WithDeviceID adds the deviceID to the order sim params
func (o *OrderSimParams) WithDeviceID(deviceID string) *OrderSimParams {
	o.SetDeviceID(deviceID)
	return o
}

// SetDeviceID adds the deviceId to the order sim params
func (o *OrderSimParams) SetDeviceID(deviceID string) {
	o.DeviceID = deviceID
}

// WithUserID adds the userID to the order sim params
func (o *OrderSimParams) WithUserID(userID string) *OrderSimParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the order sim params
func (o *OrderSimParams) SetUserID(userID string) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *OrderSimParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param deviceId
	if err := r.SetPathParam("deviceId", o.DeviceID); err != nil {
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