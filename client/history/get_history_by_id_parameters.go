// Code generated by go-swagger; DO NOT EDIT.

package history

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

// NewGetHistoryByIDParams creates a new GetHistoryByIDParams object
// with the default values initialized.
func NewGetHistoryByIDParams() *GetHistoryByIDParams {
	var ()
	return &GetHistoryByIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetHistoryByIDParamsWithTimeout creates a new GetHistoryByIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetHistoryByIDParamsWithTimeout(timeout time.Duration) *GetHistoryByIDParams {
	var ()
	return &GetHistoryByIDParams{

		timeout: timeout,
	}
}

// NewGetHistoryByIDParamsWithContext creates a new GetHistoryByIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetHistoryByIDParamsWithContext(ctx context.Context) *GetHistoryByIDParams {
	var ()
	return &GetHistoryByIDParams{

		Context: ctx,
	}
}

// NewGetHistoryByIDParamsWithHTTPClient creates a new GetHistoryByIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetHistoryByIDParamsWithHTTPClient(client *http.Client) *GetHistoryByIDParams {
	var ()
	return &GetHistoryByIDParams{
		HTTPClient: client,
	}
}

/*GetHistoryByIDParams contains all the parameters to send to the API endpoint
for the get history by Id operation typically these are written to a http.Request
*/
type GetHistoryByIDParams struct {

	/*EntryID
	  The unique call, fax, sms or voicemail identifier

	*/
	EntryID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get history by Id params
func (o *GetHistoryByIDParams) WithTimeout(timeout time.Duration) *GetHistoryByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get history by Id params
func (o *GetHistoryByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get history by Id params
func (o *GetHistoryByIDParams) WithContext(ctx context.Context) *GetHistoryByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get history by Id params
func (o *GetHistoryByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get history by Id params
func (o *GetHistoryByIDParams) WithHTTPClient(client *http.Client) *GetHistoryByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get history by Id params
func (o *GetHistoryByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEntryID adds the entryID to the get history by Id params
func (o *GetHistoryByIDParams) WithEntryID(entryID string) *GetHistoryByIDParams {
	o.SetEntryID(entryID)
	return o
}

// SetEntryID adds the entryId to the get history by Id params
func (o *GetHistoryByIDParams) SetEntryID(entryID string) {
	o.EntryID = entryID
}

// WriteToRequest writes these params to a swagger request
func (o *GetHistoryByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param entryId
	if err := r.SetPathParam("entryId", o.EntryID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
