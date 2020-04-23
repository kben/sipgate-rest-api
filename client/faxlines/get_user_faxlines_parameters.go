// Code generated by go-swagger; DO NOT EDIT.

package faxlines

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

// NewGetUserFaxlinesParams creates a new GetUserFaxlinesParams object
// with the default values initialized.
func NewGetUserFaxlinesParams() *GetUserFaxlinesParams {
	var ()
	return &GetUserFaxlinesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetUserFaxlinesParamsWithTimeout creates a new GetUserFaxlinesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetUserFaxlinesParamsWithTimeout(timeout time.Duration) *GetUserFaxlinesParams {
	var ()
	return &GetUserFaxlinesParams{

		timeout: timeout,
	}
}

// NewGetUserFaxlinesParamsWithContext creates a new GetUserFaxlinesParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetUserFaxlinesParamsWithContext(ctx context.Context) *GetUserFaxlinesParams {
	var ()
	return &GetUserFaxlinesParams{

		Context: ctx,
	}
}

// NewGetUserFaxlinesParamsWithHTTPClient creates a new GetUserFaxlinesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetUserFaxlinesParamsWithHTTPClient(client *http.Client) *GetUserFaxlinesParams {
	var ()
	return &GetUserFaxlinesParams{
		HTTPClient: client,
	}
}

/*GetUserFaxlinesParams contains all the parameters to send to the API endpoint
for the get user faxlines operation typically these are written to a http.Request
*/
type GetUserFaxlinesParams struct {

	/*UserID
	  The unique user identifier

	*/
	UserID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get user faxlines params
func (o *GetUserFaxlinesParams) WithTimeout(timeout time.Duration) *GetUserFaxlinesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get user faxlines params
func (o *GetUserFaxlinesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get user faxlines params
func (o *GetUserFaxlinesParams) WithContext(ctx context.Context) *GetUserFaxlinesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get user faxlines params
func (o *GetUserFaxlinesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get user faxlines params
func (o *GetUserFaxlinesParams) WithHTTPClient(client *http.Client) *GetUserFaxlinesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get user faxlines params
func (o *GetUserFaxlinesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUserID adds the userID to the get user faxlines params
func (o *GetUserFaxlinesParams) WithUserID(userID string) *GetUserFaxlinesParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the get user faxlines params
func (o *GetUserFaxlinesParams) SetUserID(userID string) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *GetUserFaxlinesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param userId
	if err := r.SetPathParam("userId", o.UserID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
