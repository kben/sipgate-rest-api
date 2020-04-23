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
)

// NewGetTariffAnnouncementParams creates a new GetTariffAnnouncementParams object
// with the default values initialized.
func NewGetTariffAnnouncementParams() *GetTariffAnnouncementParams {
	var ()
	return &GetTariffAnnouncementParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetTariffAnnouncementParamsWithTimeout creates a new GetTariffAnnouncementParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetTariffAnnouncementParamsWithTimeout(timeout time.Duration) *GetTariffAnnouncementParams {
	var ()
	return &GetTariffAnnouncementParams{

		timeout: timeout,
	}
}

// NewGetTariffAnnouncementParamsWithContext creates a new GetTariffAnnouncementParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetTariffAnnouncementParamsWithContext(ctx context.Context) *GetTariffAnnouncementParams {
	var ()
	return &GetTariffAnnouncementParams{

		Context: ctx,
	}
}

// NewGetTariffAnnouncementParamsWithHTTPClient creates a new GetTariffAnnouncementParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetTariffAnnouncementParamsWithHTTPClient(client *http.Client) *GetTariffAnnouncementParams {
	var ()
	return &GetTariffAnnouncementParams{
		HTTPClient: client,
	}
}

/*GetTariffAnnouncementParams contains all the parameters to send to the API endpoint
for the get tariff announcement operation typically these are written to a http.Request
*/
type GetTariffAnnouncementParams struct {

	/*DeviceID
	  The unique device identifier

	*/
	DeviceID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get tariff announcement params
func (o *GetTariffAnnouncementParams) WithTimeout(timeout time.Duration) *GetTariffAnnouncementParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get tariff announcement params
func (o *GetTariffAnnouncementParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get tariff announcement params
func (o *GetTariffAnnouncementParams) WithContext(ctx context.Context) *GetTariffAnnouncementParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get tariff announcement params
func (o *GetTariffAnnouncementParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get tariff announcement params
func (o *GetTariffAnnouncementParams) WithHTTPClient(client *http.Client) *GetTariffAnnouncementParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get tariff announcement params
func (o *GetTariffAnnouncementParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDeviceID adds the deviceID to the get tariff announcement params
func (o *GetTariffAnnouncementParams) WithDeviceID(deviceID string) *GetTariffAnnouncementParams {
	o.SetDeviceID(deviceID)
	return o
}

// SetDeviceID adds the deviceId to the get tariff announcement params
func (o *GetTariffAnnouncementParams) SetDeviceID(deviceID string) {
	o.DeviceID = deviceID
}

// WriteToRequest writes these params to a swagger request
func (o *GetTariffAnnouncementParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param deviceId
	if err := r.SetPathParam("deviceId", o.DeviceID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
