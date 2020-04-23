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

// ChangeVoicemailRequest change voicemail request
//
// swagger:model ChangeVoicemailRequest
type ChangeVoicemailRequest struct {

	// active
	// Required: true
	Active bool `json:"active"`

	// timeout
	// Minimum: 0
	Timeout *int32 `json:"timeout,omitempty"`

	// transcription
	// Required: true
	Transcription bool `json:"transcription"`
}

// Validate validates this change voicemail request
func (m *ChangeVoicemailRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateActive(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTimeout(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTranscription(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ChangeVoicemailRequest) validateActive(formats strfmt.Registry) error {

	if err := validate.Required("active", "body", bool(m.Active)); err != nil {
		return err
	}

	return nil
}

func (m *ChangeVoicemailRequest) validateTimeout(formats strfmt.Registry) error {

	if swag.IsZero(m.Timeout) { // not required
		return nil
	}

	if err := validate.MinimumInt("timeout", "body", int64(*m.Timeout), 0, false); err != nil {
		return err
	}

	return nil
}

func (m *ChangeVoicemailRequest) validateTranscription(formats strfmt.Registry) error {

	if err := validate.Required("transcription", "body", bool(m.Transcription)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ChangeVoicemailRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ChangeVoicemailRequest) UnmarshalBinary(b []byte) error {
	var res ChangeVoicemailRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}