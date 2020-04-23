// Code generated by go-swagger; DO NOT EDIT.

package numbers

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

// NewChangeQuickDialNumberParams creates a new ChangeQuickDialNumberParams object
// with the default values initialized.
func NewChangeQuickDialNumberParams() *ChangeQuickDialNumberParams {
	var ()
	return &ChangeQuickDialNumberParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewChangeQuickDialNumberParamsWithTimeout creates a new ChangeQuickDialNumberParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewChangeQuickDialNumberParamsWithTimeout(timeout time.Duration) *ChangeQuickDialNumberParams {
	var ()
	return &ChangeQuickDialNumberParams{

		timeout: timeout,
	}
}

// NewChangeQuickDialNumberParamsWithContext creates a new ChangeQuickDialNumberParams object
// with the default values initialized, and the ability to set a context for a request
func NewChangeQuickDialNumberParamsWithContext(ctx context.Context) *ChangeQuickDialNumberParams {
	var ()
	return &ChangeQuickDialNumberParams{

		Context: ctx,
	}
}

// NewChangeQuickDialNumberParamsWithHTTPClient creates a new ChangeQuickDialNumberParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewChangeQuickDialNumberParamsWithHTTPClient(client *http.Client) *ChangeQuickDialNumberParams {
	var ()
	return &ChangeQuickDialNumberParams{
		HTTPClient: client,
	}
}

/*ChangeQuickDialNumberParams contains all the parameters to send to the API endpoint
for the change quick dial number operation typically these are written to a http.Request
*/
type ChangeQuickDialNumberParams struct {

	/*Body*/
	Body *models.ChangeDirectDialRequest
	/*QuickdialID*/
	QuickdialID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the change quick dial number params
func (o *ChangeQuickDialNumberParams) WithTimeout(timeout time.Duration) *ChangeQuickDialNumberParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the change quick dial number params
func (o *ChangeQuickDialNumberParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the change quick dial number params
func (o *ChangeQuickDialNumberParams) WithContext(ctx context.Context) *ChangeQuickDialNumberParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the change quick dial number params
func (o *ChangeQuickDialNumberParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the change quick dial number params
func (o *ChangeQuickDialNumberParams) WithHTTPClient(client *http.Client) *ChangeQuickDialNumberParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the change quick dial number params
func (o *ChangeQuickDialNumberParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the change quick dial number params
func (o *ChangeQuickDialNumberParams) WithBody(body *models.ChangeDirectDialRequest) *ChangeQuickDialNumberParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the change quick dial number params
func (o *ChangeQuickDialNumberParams) SetBody(body *models.ChangeDirectDialRequest) {
	o.Body = body
}

// WithQuickdialID adds the quickdialID to the change quick dial number params
func (o *ChangeQuickDialNumberParams) WithQuickdialID(quickdialID string) *ChangeQuickDialNumberParams {
	o.SetQuickdialID(quickdialID)
	return o
}

// SetQuickdialID adds the quickdialId to the change quick dial number params
func (o *ChangeQuickDialNumberParams) SetQuickdialID(quickdialID string) {
	o.QuickdialID = quickdialID
}

// WriteToRequest writes these params to a swagger request
func (o *ChangeQuickDialNumberParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param quickdialId
	if err := r.SetPathParam("quickdialId", o.QuickdialID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}