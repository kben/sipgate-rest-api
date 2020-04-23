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

// NotificationsResponse notifications response
//
// swagger:model NotificationsResponse
type NotificationsResponse struct {

	// call
	Call []*CallNotificationResponse `json:"call"`

	// fax
	Fax []*FaxNotificationResponse `json:"fax"`

	// sms
	Sms []*SmsNotificationResponse `json:"sms"`

	// voicemail
	Voicemail []*VoicemailNotificationResponse `json:"voicemail"`
}

// Validate validates this notifications response
func (m *NotificationsResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCall(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFax(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSms(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVoicemail(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NotificationsResponse) validateCall(formats strfmt.Registry) error {

	if swag.IsZero(m.Call) { // not required
		return nil
	}

	for i := 0; i < len(m.Call); i++ {
		if swag.IsZero(m.Call[i]) { // not required
			continue
		}

		if m.Call[i] != nil {
			if err := m.Call[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("call" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *NotificationsResponse) validateFax(formats strfmt.Registry) error {

	if swag.IsZero(m.Fax) { // not required
		return nil
	}

	for i := 0; i < len(m.Fax); i++ {
		if swag.IsZero(m.Fax[i]) { // not required
			continue
		}

		if m.Fax[i] != nil {
			if err := m.Fax[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("fax" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *NotificationsResponse) validateSms(formats strfmt.Registry) error {

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

func (m *NotificationsResponse) validateVoicemail(formats strfmt.Registry) error {

	if swag.IsZero(m.Voicemail) { // not required
		return nil
	}

	for i := 0; i < len(m.Voicemail); i++ {
		if swag.IsZero(m.Voicemail[i]) { // not required
			continue
		}

		if m.Voicemail[i] != nil {
			if err := m.Voicemail[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("voicemail" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *NotificationsResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NotificationsResponse) UnmarshalBinary(b []byte) error {
	var res NotificationsResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}