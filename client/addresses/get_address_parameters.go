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
)

// NewGetAddressParams creates a new GetAddressParams object
// with the default values initialized.
func NewGetAddressParams() *GetAddressParams {
	var ()
	return &GetAddressParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetAddressParamsWithTimeout creates a new GetAddressParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetAddressParamsWithTimeout(timeout time.Duration) *GetAddressParams {
	var ()
	return &GetAddressParams{

		timeout: timeout,
	}
}

// NewGetAddressParamsWithContext creates a new GetAddressParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetAddressParamsWithContext(ctx context.Context) *GetAddressParams {
	var ()
	return &GetAddressParams{

		Context: ctx,
	}
}

// NewGetAddressParamsWithHTTPClient creates a new GetAddressParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetAddressParamsWithHTTPClient(client *http.Client) *GetAddressParams {
	var ()
	return &GetAddressParams{
		HTTPClient: client,
	}
}

/*GetAddressParams contains all the parameters to send to the API endpoint
for the get address operation typically these are written to a http.Request
*/
type GetAddressParams struct {

	/*AddressID
	  The unique address identifier

	*/
	AddressID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get address params
func (o *GetAddressParams) WithTimeout(timeout time.Duration) *GetAddressParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get address params
func (o *GetAddressParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get address params
func (o *GetAddressParams) WithContext(ctx context.Context) *GetAddressParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get address params
func (o *GetAddressParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get address params
func (o *GetAddressParams) WithHTTPClient(client *http.Client) *GetAddressParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get address params
func (o *GetAddressParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAddressID adds the addressID to the get address params
func (o *GetAddressParams) WithAddressID(addressID int32) *GetAddressParams {
	o.SetAddressID(addressID)
	return o
}

// SetAddressID adds the addressId to the get address params
func (o *GetAddressParams) SetAddressID(addressID int32) {
	o.AddressID = addressID
}

// WriteToRequest writes these params to a swagger request
func (o *GetAddressParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param addressId
	if err := r.SetPathParam("addressId", swag.FormatInt32(o.AddressID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
