// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// VoicemailNotificationResponse voicemail notification response
//
// swagger:model VoicemailNotificationResponse
type VoicemailNotificationResponse struct {

	// emails
	Emails []*EmailTarget `json:"emails"`

	// endpoint alias
	EndpointAlias string `json:"endpointAlias,omitempty"`

	// endpoint Id
	EndpointID string `json:"endpointId,omitempty"`

	// sms
	Sms []*SmsTarget `json:"sms"`
}

// Validate validates this voicemail notification response
func (m *VoicemailNotificationResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEmails(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSms(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VoicemailNotificationResponse) validateEmails(formats strfmt.Registry) error {

	if swag.IsZero(m.Emails) { // not required
		return nil
	}

	for i := 0; i < len(m.Emails); i++ {
		if swag.IsZero(m.Emails[i]) { // not required
			continue
		}

		if m.Emails[i] != nil {
			if err := m.Emails[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("emails" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *VoicemailNotificationResponse) validateSms(formats strfmt.Registry) error {

	if swag.IsZero(m.Sms) { // not required
		return nil
	}

	for i := 0; i < len(m.Sms); i++ {
		if swag.IsZero(m.Sms[i]) { // not required
			continue
		}

		if m.Sms[i] != nil {
			if err := m.Sms[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("sms" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *VoicemailNotificationResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VoicemailNotificationResponse) UnmarshalBinary(b []byte) error {
	var res VoicemailNotificationResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
