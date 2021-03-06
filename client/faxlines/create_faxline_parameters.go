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

// NewCreateFaxlineParams creates a new CreateFaxlineParams object
// with the default values initialized.
func NewCreateFaxlineParams() *CreateFaxlineParams {
	var ()
	return &CreateFaxlineParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateFaxlineParamsWithTimeout creates a new CreateFaxlineParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateFaxlineParamsWithTimeout(timeout time.Duration) *CreateFaxlineParams {
	var ()
	return &CreateFaxlineParams{

		timeout: timeout,
	}
}

// NewCreateFaxlineParamsWithContext creates a new CreateFaxlineParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateFaxlineParamsWithContext(ctx context.Context) *CreateFaxlineParams {
	var ()
	return &CreateFaxlineParams{

		Context: ctx,
	}
}

// NewCreateFaxlineParamsWithHTTPClient creates a new CreateFaxlineParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateFaxlineParamsWithHTTPClient(client *http.Client) *CreateFaxlineParams {
	var ()
	return &CreateFaxlineParams{
		HTTPClient: client,
	}
}

/*CreateFaxlineParams contains all the parameters to send to the API endpoint
for the create faxline operation typically these are written to a http.Request
*/
type CreateFaxlineParams struct {

	/*UserID
	  The unique user identifier

	*/
	UserID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create faxline params
func (o *CreateFaxlineParams) WithTimeout(timeout time.Duration) *CreateFaxlineParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create faxline params
func (o *CreateFaxlineParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create faxline params
func (o *CreateFaxlineParams) WithContext(ctx context.Context) *CreateFaxlineParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create faxline params
func (o *CreateFaxlineParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create faxline params
func (o *CreateFaxlineParams) WithHTTPClient(client *http.Client) *CreateFaxlineParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create faxline params
func (o *CreateFaxlineParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUserID adds the userID to the create faxline params
func (o *CreateFaxlineParams) WithUserID(userID string) *CreateFaxlineParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the create faxline params
func (o *CreateFaxlineParams) SetUserID(userID string) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *CreateFaxlineParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
