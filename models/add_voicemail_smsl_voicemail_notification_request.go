// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// AddVoicemailSmslVoicemailNotificationRequest add voicemail smsl voicemail notification request
//
// swagger:model AddVoicemailSmslVoicemailNotificationRequest
type AddVoicemailSmslVoicemailNotificationRequest struct {

	// number
	// Required: true
	Number *string `json:"number"`

	// voicemail Id
	// Required: true
	VoicemailID *string `json:"voicemailId"`
}

// Validate validates this add voicemail smsl voicemail notification request
func (m *AddVoicemailSmslVoicemailNotificationRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNumber(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVoicemailID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AddVoicemailSmslVoicemailNotificationRequest) validateNumber(formats strfmt.Registry) error {

	if err := validate.Required("number", "body", m.Number); err != nil {
		return err
	}

	return nil
}

func (m *AddVoicemailSmslVoicemailNotificationRequest) validateVoicemailID(formats strfmt.Registry) error {

	if err := validate.Required("voicemailId", "body", m.VoicemailID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AddVoicemailSmslVoicemailNotificationRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AddVoicemailSmslVoicemailNotificationRequest) UnmarshalBinary(b []byte) error {
	var res AddVoicemailSmslVoicemailNotificationRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
