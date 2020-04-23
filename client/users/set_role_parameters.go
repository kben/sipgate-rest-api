// Code generated by go-swagger; DO NOT EDIT.

package users

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

// NewSetRoleParams creates a new SetRoleParams object
// with the default values initialized.
func NewSetRoleParams() *SetRoleParams {
	var ()
	return &SetRoleParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSetRoleParamsWithTimeout creates a new SetRoleParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSetRoleParamsWithTimeout(timeout time.Duration) *SetRoleParams {
	var ()
	return &SetRoleParams{

		timeout: timeout,
	}
}

// NewSetRoleParamsWithContext creates a new SetRoleParams object
// with the default values initialized, and the ability to set a context for a request
func NewSetRoleParamsWithContext(ctx context.Context) *SetRoleParams {
	var ()
	return &SetRoleParams{

		Context: ctx,
	}
}

// NewSetRoleParamsWithHTTPClient creates a new SetRoleParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSetRoleParamsWithHTTPClient(client *http.Client) *SetRoleParams {
	var ()
	return &SetRoleParams{
		HTTPClient: client,
	}
}

/*SetRoleParams contains all the parameters to send to the API endpoint
for the set role operation typically these are written to a http.Request
*/
type SetRoleParams struct {

	/*Body*/
	Body *models.SetUserRoleRequest
	/*UserID
	  The unique user identifier

	*/
	UserID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the set role params
func (o *SetRoleParams) WithTimeout(timeout time.Duration) *SetRoleParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the set role params
func (o *SetRoleParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the set role params
func (o *SetRoleParams) WithContext(ctx context.Context) *SetRoleParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the set role params
func (o *SetRoleParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the set role params
func (o *SetRoleParams) WithHTTPClient(client *http.Client) *SetRoleParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the set role params
func (o *SetRoleParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the set role params
func (o *SetRoleParams) WithBody(body *models.SetUserRoleRequest) *SetRoleParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the set role params
func (o *SetRoleParams) SetBody(body *models.SetUserRoleRequest) {
	o.Body = body
}

// WithUserID adds the userID to the set role params
func (o *SetRoleParams) WithUserID(userID string) *SetRoleParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the set role params
func (o *SetRoleParams) SetUserID(userID string) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *SetRoleParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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