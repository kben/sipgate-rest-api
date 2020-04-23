// Code generated by go-swagger; DO NOT EDIT.

package groups

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

// NewGetGroupVoicemailParams creates a new GetGroupVoicemailParams object
// with the default values initialized.
func NewGetGroupVoicemailParams() *GetGroupVoicemailParams {
	var ()
	return &GetGroupVoicemailParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetGroupVoicemailParamsWithTimeout creates a new GetGroupVoicemailParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetGroupVoicemailParamsWithTimeout(timeout time.Duration) *GetGroupVoicemailParams {
	var ()
	return &GetGroupVoicemailParams{

		timeout: timeout,
	}
}

// NewGetGroupVoicemailParamsWithContext creates a new GetGroupVoicemailParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetGroupVoicemailParamsWithContext(ctx context.Context) *GetGroupVoicemailParams {
	var ()
	return &GetGroupVoicemailParams{

		Context: ctx,
	}
}

// NewGetGroupVoicemailParamsWithHTTPClient creates a new GetGroupVoicemailParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetGroupVoicemailParamsWithHTTPClient(client *http.Client) *GetGroupVoicemailParams {
	var ()
	return &GetGroupVoicemailParams{
		HTTPClient: client,
	}
}

/*GetGroupVoicemailParams contains all the parameters to send to the API endpoint
for the get group voicemail operation typically these are written to a http.Request
*/
type GetGroupVoicemailParams struct {

	/*GroupID
	  The unique group identifier

	*/
	GroupID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get group voicemail params
func (o *GetGroupVoicemailParams) WithTimeout(timeout time.Duration) *GetGroupVoicemailParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get group voicemail params
func (o *GetGroupVoicemailParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get group voicemail params
func (o *GetGroupVoicemailParams) WithContext(ctx context.Context) *GetGroupVoicemailParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get group voicemail params
func (o *GetGroupVoicemailParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get group voicemail params
func (o *GetGroupVoicemailParams) WithHTTPClient(client *http.Client) *GetGroupVoicemailParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get group voicemail params
func (o *GetGroupVoicemailParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithGroupID adds the groupID to the get group voicemail params
func (o *GetGroupVoicemailParams) WithGroupID(groupID string) *GetGroupVoicemailParams {
	o.SetGroupID(groupID)
	return o
}

// SetGroupID adds the groupId to the get group voicemail params
func (o *GetGroupVoicemailParams) SetGroupID(groupID string) {
	o.GroupID = groupID
}

// WriteToRequest writes these params to a swagger request
func (o *GetGroupVoicemailParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param groupId
	if err := r.SetPathParam("groupId", o.GroupID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}