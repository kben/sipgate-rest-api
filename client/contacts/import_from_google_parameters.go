// Code generated by go-swagger; DO NOT EDIT.

package contacts

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

// NewImportFromGoogleParams creates a new ImportFromGoogleParams object
// with the default values initialized.
func NewImportFromGoogleParams() *ImportFromGoogleParams {
	var ()
	return &ImportFromGoogleParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewImportFromGoogleParamsWithTimeout creates a new ImportFromGoogleParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewImportFromGoogleParamsWithTimeout(timeout time.Duration) *ImportFromGoogleParams {
	var ()
	return &ImportFromGoogleParams{

		timeout: timeout,
	}
}

// NewImportFromGoogleParamsWithContext creates a new ImportFromGoogleParams object
// with the default values initialized, and the ability to set a context for a request
func NewImportFromGoogleParamsWithContext(ctx context.Context) *ImportFromGoogleParams {
	var ()
	return &ImportFromGoogleParams{

		Context: ctx,
	}
}

// NewImportFromGoogleParamsWithHTTPClient creates a new ImportFromGoogleParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewImportFromGoogleParamsWithHTTPClient(client *http.Client) *ImportFromGoogleParams {
	var ()
	return &ImportFromGoogleParams{
		HTTPClient: client,
	}
}

/*ImportFromGoogleParams contains all the parameters to send to the API endpoint
for the import from google operation typically these are written to a http.Request
*/
type ImportFromGoogleParams struct {

	/*Body*/
	Body *models.ImportGoogleRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the import from google params
func (o *ImportFromGoogleParams) WithTimeout(timeout time.Duration) *ImportFromGoogleParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the import from google params
func (o *ImportFromGoogleParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the import from google params
func (o *ImportFromGoogleParams) WithContext(ctx context.Context) *ImportFromGoogleParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the import from google params
func (o *ImportFromGoogleParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the import from google params
func (o *ImportFromGoogleParams) WithHTTPClient(client *http.Client) *ImportFromGoogleParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the import from google params
func (o *ImportFromGoogleParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the import from google params
func (o *ImportFromGoogleParams) WithBody(body *models.ImportGoogleRequest) *ImportFromGoogleParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the import from google params
func (o *ImportFromGoogleParams) SetBody(body *models.ImportGoogleRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *ImportFromGoogleParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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