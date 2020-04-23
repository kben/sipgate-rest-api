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

// NewChangeVoicemailSettingsParams creates a new ChangeVoicemailSettingsParams object
// with the default values initialized.
func NewChangeVoicemailSettingsParams() *ChangeVoicemailSettingsParams {
	var ()
	return &ChangeVoicemailSettingsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewChangeVoicemailSettingsParamsWithTimeout creates a new ChangeVoicemailSettingsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewChangeVoicemailSettingsParamsWithTimeout(timeout time.Duration) *ChangeVoicemailSettingsParams {
	var ()
	return &ChangeVoicemailSettingsParams{

		timeout: timeout,
	}
}

// NewChangeVoicemailSettingsParamsWithContext creates a new ChangeVoicemailSettingsParams object
// with the default values initialized, and the ability to set a context for a request
func NewChangeVoicemailSettingsParamsWithContext(ctx context.Context) *ChangeVoicemailSettingsParams {
	var ()
	return &ChangeVoicemailSettingsParams{

		Context: ctx,
	}
}

// NewChangeVoicemailSettingsParamsWithHTTPClient creates a new ChangeVoicemailSettingsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewChangeVoicemailSettingsParamsWithHTTPClient(client *http.Client) *ChangeVoicemailSettingsParams {
	var ()
	return &ChangeVoicemailSettingsParams{
		HTTPClient: client,
	}
}

/*ChangeVoicemailSettingsParams contains all the parameters to send to the API endpoint
for the change voicemail settings operation typically these are written to a http.Request
*/
type ChangeVoicemailSettingsParams struct {

	/*Body*/
	Body *models.ChangeVoicemailRequest
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

// WithTimeout adds the timeout to the change voicemail settings params
func (o *ChangeVoicemailSettingsParams) WithTimeout(timeout time.Duration) *ChangeVoicemailSettingsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the change voicemail settings params
func (o *ChangeVoicemailSettingsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the change voicemail settings params
func (o *ChangeVoicemailSettingsParams) WithContext(ctx context.Context) *ChangeVoicemailSettingsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the change voicemail settings params
func (o *ChangeVoicemailSettingsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the change voicemail settings params
func (o *ChangeVoicemailSettingsParams) WithHTTPClient(client *http.Client) *ChangeVoicemailSettingsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the change voicemail settings params
func (o *ChangeVoicemailSettingsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the change voicemail settings params
func (o *ChangeVoicemailSettingsParams) WithBody(body *models.ChangeVoicemailRequest) *ChangeVoicemailSettingsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the change voicemail settings params
func (o *ChangeVoicemailSettingsParams) SetBody(body *models.ChangeVoicemailRequest) {
	o.Body = body
}

// WithPhonelineID adds the phonelineID to the change voicemail settings params
func (o *ChangeVoicemailSettingsParams) WithPhonelineID(phonelineID string) *ChangeVoicemailSettingsParams {
	o.SetPhonelineID(phonelineID)
	return o
}

// SetPhonelineID adds the phonelineId to the change voicemail settings params
func (o *ChangeVoicemailSettingsParams) SetPhonelineID(phonelineID string) {
	o.PhonelineID = phonelineID
}

// WithUserID adds the userID to the change voicemail settings params
func (o *ChangeVoicemailSettingsParams) WithUserID(userID string) *ChangeVoicemailSettingsParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the change voicemail settings params
func (o *ChangeVoicemailSettingsParams) SetUserID(userID string) {
	o.UserID = userID
}

// WithVoicemailID adds the voicemailID to the change voicemail settings params
func (o *ChangeVoicemailSettingsParams) WithVoicemailID(voicemailID string) *ChangeVoicemailSettingsParams {
	o.SetVoicemailID(voicemailID)
	return o
}

// SetVoicemailID adds the voicemailId to the change voicemail settings params
func (o *ChangeVoicemailSettingsParams) SetVoicemailID(voicemailID string) {
	o.VoicemailID = voicemailID
}

// WriteToRequest writes these params to a swagger request
func (o *ChangeVoicemailSettingsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
