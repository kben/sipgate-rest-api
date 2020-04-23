// Code generated by go-swagger; DO NOT EDIT.

package phonelines

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

// NewAddGreetingParams creates a new AddGreetingParams object
// with the default values initialized.
func NewAddGreetingParams() *AddGreetingParams {
	var ()
	return &AddGreetingParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAddGreetingParamsWithTimeout creates a new AddGreetingParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAddGreetingParamsWithTimeout(timeout time.Duration) *AddGreetingParams {
	var ()
	return &AddGreetingParams{

		timeout: timeout,
	}
}

// NewAddGreetingParamsWithContext creates a new AddGreetingParams object
// with the default values initialized, and the ability to set a context for a request
func NewAddGreetingParamsWithContext(ctx context.Context) *AddGreetingParams {
	var ()
	return &AddGreetingParams{

		Context: ctx,
	}
}

// NewAddGreetingParamsWithHTTPClient creates a new AddGreetingParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAddGreetingParamsWithHTTPClient(client *http.Client) *AddGreetingParams {
	var ()
	return &AddGreetingParams{
		HTTPClient: client,
	}
}

/*AddGreetingParams contains all the parameters to send to the API endpoint
for the add greeting operation typically these are written to a http.Request
*/
type AddGreetingParams struct {

	/*Body*/
	Body *models.AddGreetingRequest
	/*PhonelineID
	  The unique phone line identifier

	*/
	PhonelineID string
	/*UserID
	  The unique user identifier

	*/
	UserID string
	/*VoicemailID
	  The unique voicemail identifier

	*/
	VoicemailID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the add greeting params
func (o *AddGreetingParams) WithTimeout(timeout time.Duration) *AddGreetingParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add greeting params
func (o *AddGreetingParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add greeting params
func (o *AddGreetingParams) WithContext(ctx context.Context) *AddGreetingParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add greeting params
func (o *AddGreetingParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add greeting params
func (o *AddGreetingParams) WithHTTPClient(client *http.Client) *AddGreetingParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add greeting params
func (o *AddGreetingParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the add greeting params
func (o *AddGreetingParams) WithBody(body *models.AddGreetingRequest) *AddGreetingParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the add greeting params
func (o *AddGreetingParams) SetBody(body *models.AddGreetingRequest) {
	o.Body = body
}

// WithPhonelineID adds the phonelineID to the add greeting params
func (o *AddGreetingParams) WithPhonelineID(phonelineID string) *AddGreetingParams {
	o.SetPhonelineID(phonelineID)
	return o
}

// SetPhonelineID adds the phonelineId to the add greeting params
func (o *AddGreetingParams) SetPhonelineID(phonelineID string) {
	o.PhonelineID = phonelineID
}

// WithUserID adds the userID to the add greeting params
func (o *AddGreetingParams) WithUserID(userID string) *AddGreetingParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the add greeting params
func (o *AddGreetingParams) SetUserID(userID string) {
	o.UserID = userID
}

// WithVoicemailID adds the voicemailID to the add greeting params
func (o *AddGreetingParams) WithVoicemailID(voicemailID string) *AddGreetingParams {
	o.SetVoicemailID(voicemailID)
	return o
}

// SetVoicemailID adds the voicemailId to the add greeting params
func (o *AddGreetingParams) SetVoicemailID(voicemailID string) {
	o.VoicemailID = voicemailID
}

// WriteToRequest writes these params to a swagger request
func (o *AddGreetingParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param phonelineId
	if err := r.SetPathParam("phonelineId", o.PhonelineID); err != nil {
		return err
	}

	// path param userId
	if err := r.SetPathParam("userId", o.UserID); err != nil {
		return err
	}

	// path param voicemailId
	if err := r.SetPathParam("voicemailId", o.VoicemailID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
