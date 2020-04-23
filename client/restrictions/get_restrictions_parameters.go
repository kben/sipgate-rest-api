// Code generated by go-swagger; DO NOT EDIT.

package restrictions

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

// NewGetRestrictionsParams creates a new GetRestrictionsParams object
// with the default values initialized.
func NewGetRestrictionsParams() *GetRestrictionsParams {
	var ()
	return &GetRestrictionsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetRestrictionsParamsWithTimeout creates a new GetRestrictionsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetRestrictionsParamsWithTimeout(timeout time.Duration) *GetRestrictionsParams {
	var ()
	return &GetRestrictionsParams{

		timeout: timeout,
	}
}

// NewGetRestrictionsParamsWithContext creates a new GetRestrictionsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetRestrictionsParamsWithContext(ctx context.Context) *GetRestrictionsParams {
	var ()
	return &GetRestrictionsParams{

		Context: ctx,
	}
}

// NewGetRestrictionsParamsWithHTTPClient creates a new GetRestrictionsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetRestrictionsParamsWithHTTPClient(client *http.Client) *GetRestrictionsParams {
	var ()
	return &GetRestrictionsParams{
		HTTPClient: client,
	}
}

/*GetRestrictionsParams contains all the parameters to send to the API endpoint
for the get restrictions operation typically these are written to a http.Request
*/
type GetRestrictionsParams struct {

	/*Restriction
	  The requested restrictions

	*/
	Restriction []string
	/*UserID
	  The unique user identifier

	*/
	UserID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get restrictions params
func (o *GetRestrictionsParams) WithTimeout(timeout time.Duration) *GetRestrictionsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get restrictions params
func (o *GetRestrictionsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get restrictions params
func (o *GetRestrictionsParams) WithContext(ctx context.Context) *GetRestrictionsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get restrictions params
func (o *GetRestrictionsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get restrictions params
func (o *GetRestrictionsParams) WithHTTPClient(client *http.Client) *GetRestrictionsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get restrictions params
func (o *GetRestrictionsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRestriction adds the restriction to the get restrictions params
func (o *GetRestrictionsParams) WithRestriction(restriction []string) *GetRestrictionsParams {
	o.SetRestriction(restriction)
	return o
}

// SetRestriction adds the restriction to the get restrictions params
func (o *GetRestrictionsParams) SetRestriction(restriction []string) {
	o.Restriction = restriction
}

// WithUserID adds the userID to the get restrictions params
func (o *GetRestrictionsParams) WithUserID(userID string) *GetRestrictionsParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the get restrictions params
func (o *GetRestrictionsParams) SetUserID(userID string) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *GetRestrictionsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	valuesRestriction := o.Restriction

	joinedRestriction := swag.JoinByFormat(valuesRestriction, "multi")
	// query array param restriction
	if err := r.SetQueryParam("restriction", joinedRestriction...); err != nil {
		return err
	}

	// query param userId
	qrUserID := o.UserID
	qUserID := qrUserID
	if qUserID != "" {
		if err := r.SetQueryParam("userId", qUserID); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
