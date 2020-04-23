// Code generated by go-swagger; DO NOT EDIT.

package addresses

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
	"github.com/go-openapi/swag"

	"github.com/kben/sipgate-rest-api/models"
)

// NewModifyAddressParams creates a new ModifyAddressParams object
// with the default values initialized.
func NewModifyAddressParams() *ModifyAddressParams {
	var ()
	return &ModifyAddressParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewModifyAddressParamsWithTimeout creates a new ModifyAddressParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewModifyAddressParamsWithTimeout(timeout time.Duration) *ModifyAddressParams {
	var ()
	return &ModifyAddressParams{

		timeout: timeout,
	}
}

// NewModifyAddressParamsWithContext creates a new ModifyAddressParams object
// with the default values initialized, and the ability to set a context for a request
func NewModifyAddressParamsWithContext(ctx context.Context) *ModifyAddressParams {
	var ()
	return &ModifyAddressParams{

		Context: ctx,
	}
}

// NewModifyAddressParamsWithHTTPClient creates a new ModifyAddressParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewModifyAddressParamsWithHTTPClient(client *http.Client) *ModifyAddressParams {
	var ()
	return &ModifyAddressParams{
		HTTPClient: client,
	}
}

/*ModifyAddressParams contains all the parameters to send to the API endpoint
for the modify address operation typically these are written to a http.Request
*/
type ModifyAddressParams struct {

	/*AddressID
	  The unique address identifier

	*/
	AddressID int32
	/*Body*/
	Body *models.ModifyAddressRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the modify address params
func (o *ModifyAddressParams) WithTimeout(timeout time.Duration) *ModifyAddressParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the modify address params
func (o *ModifyAddressParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the modify address params
func (o *ModifyAddressParams) WithContext(ctx context.Context) *ModifyAddressParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the modify address params
func (o *ModifyAddressParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the modify address params
func (o *ModifyAddressParams) WithHTTPClient(client *http.Client) *ModifyAddressParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the modify address params
func (o *ModifyAddressParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAddressID adds the addressID to the modify address params
func (o *ModifyAddressParams) WithAddressID(addressID int32) *ModifyAddressParams {
	o.SetAddressID(addressID)
	return o
}

// SetAddressID adds the addressId to the modify address params
func (o *ModifyAddressParams) SetAddressID(addressID int32) {
	o.AddressID = addressID
}

// WithBody adds the body to the modify address params
func (o *ModifyAddressParams) WithBody(body *models.ModifyAddressRequest) *ModifyAddressParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the modify address params
func (o *ModifyAddressParams) SetBody(body *models.ModifyAddressRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *ModifyAddressParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param addressId
	if err := r.SetPathParam("addressId", swag.FormatInt32(o.AddressID)); err != nil {
		return err
	}

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