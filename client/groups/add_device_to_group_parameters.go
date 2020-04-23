// Code generated by go-swagger; DO NOT EDIT.

package groups

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

// NewAddDeviceToGroupParams creates a new AddDeviceToGroupParams object
// with the default values initialized.
func NewAddDeviceToGroupParams() *AddDeviceToGroupParams {
	var ()
	return &AddDeviceToGroupParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAddDeviceToGroupParamsWithTimeout creates a new AddDeviceToGroupParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAddDeviceToGroupParamsWithTimeout(timeout time.Duration) *AddDeviceToGroupParams {
	var ()
	return &AddDeviceToGroupParams{

		timeout: timeout,
	}
}

// NewAddDeviceToGroupParamsWithContext creates a new AddDeviceToGroupParams object
// with the default values initialized, and the ability to set a context for a request
func NewAddDeviceToGroupParamsWithContext(ctx context.Context) *AddDeviceToGroupParams {
	var ()
	return &AddDeviceToGroupParams{

		Context: ctx,
	}
}

// NewAddDeviceToGroupParamsWithHTTPClient creates a new AddDeviceToGroupParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAddDeviceToGroupParamsWithHTTPClient(client *http.Client) *AddDeviceToGroupParams {
	var ()
	return &AddDeviceToGroupParams{
		HTTPClient: client,
	}
}

/*AddDeviceToGroupParams contains all the parameters to send to the API endpoint
for the add device to group operation typically these are written to a http.Request
*/
type AddDeviceToGroupParams struct {

	/*Body*/
	Body *models.RouteDeviceRequest
	/*GroupID
	  The unique group identifier

	*/
	GroupID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the add device to group params
func (o *AddDeviceToGroupParams) WithTimeout(timeout time.Duration) *AddDeviceToGroupParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add device to group params
func (o *AddDeviceToGroupParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add device to group params
func (o *AddDeviceToGroupParams) WithContext(ctx context.Context) *AddDeviceToGroupParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add device to group params
func (o *AddDeviceToGroupParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add device to group params
func (o *AddDeviceToGroupParams) WithHTTPClient(client *http.Client) *AddDeviceToGroupParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add device to group params
func (o *AddDeviceToGroupParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the add device to group params
func (o *AddDeviceToGroupParams) WithBody(body *models.RouteDeviceRequest) *AddDeviceToGroupParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the add device to group params
func (o *AddDeviceToGroupParams) SetBody(body *models.RouteDeviceRequest) {
	o.Body = body
}

// WithGroupID adds the groupID to the add device to group params
func (o *AddDeviceToGroupParams) WithGroupID(groupID string) *AddDeviceToGroupParams {
	o.SetGroupID(groupID)
	return o
}

// SetGroupID adds the groupId to the add device to group params
func (o *AddDeviceToGroupParams) SetGroupID(groupID string) {
	o.GroupID = groupID
}

// WriteToRequest writes these params to a swagger request
func (o *AddDeviceToGroupParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param groupId
	if err := r.SetPathParam("groupId", o.GroupID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
