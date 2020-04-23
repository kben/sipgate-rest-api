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

// NewImportFromCSVParams creates a new ImportFromCSVParams object
// with the default values initialized.
func NewImportFromCSVParams() *ImportFromCSVParams {
	var ()
	return &ImportFromCSVParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewImportFromCSVParamsWithTimeout creates a new ImportFromCSVParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewImportFromCSVParamsWithTimeout(timeout time.Duration) *ImportFromCSVParams {
	var ()
	return &ImportFromCSVParams{

		timeout: timeout,
	}
}

// NewImportFromCSVParamsWithContext creates a new ImportFromCSVParams object
// with the default values initialized, and the ability to set a context for a request
func NewImportFromCSVParamsWithContext(ctx context.Context) *ImportFromCSVParams {
	var ()
	return &ImportFromCSVParams{

		Context: ctx,
	}
}

// NewImportFromCSVParamsWithHTTPClient creates a new ImportFromCSVParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewImportFromCSVParamsWithHTTPClient(client *http.Client) *ImportFromCSVParams {
	var ()
	return &ImportFromCSVParams{
		HTTPClient: client,
	}
}

/*ImportFromCSVParams contains all the parameters to send to the API endpoint
for the import from c s v operation typically these are written to a http.Request
*/
type ImportFromCSVParams struct {

	/*Body*/
	Body *models.ImportCSVRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the import from c s v params
func (o *ImportFromCSVParams) WithTimeout(timeout time.Duration) *ImportFromCSVParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the import from c s v params
func (o *ImportFromCSVParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the import from c s v params
func (o *ImportFromCSVParams) WithContext(ctx context.Context) *ImportFromCSVParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the import from c s v params
func (o *ImportFromCSVParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the import from c s v params
func (o *ImportFromCSVParams) WithHTTPClient(client *http.Client) *ImportFromCSVParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the import from c s v params
func (o *ImportFromCSVParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the import from c s v params
func (o *ImportFromCSVParams) WithBody(body *models.ImportCSVRequest) *ImportFromCSVParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the import from c s v params
func (o *ImportFromCSVParams) SetBody(body *models.ImportCSVRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *ImportFromCSVParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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