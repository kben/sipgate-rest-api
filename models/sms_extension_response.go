// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// SmsExtensionResponse sms extension response
//
// swagger:model SmsExtensionResponse
type SmsExtensionResponse struct {

	// alias
	Alias string `json:"alias,omitempty"`

	// caller Id
	CallerID string `json:"callerId,omitempty"`

	// id
	ID string `json:"id,omitempty"`
}

// Validate validates this sms extension response
func (m *SmsExtensionResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SmsExtensionResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SmsExtensionResponse) UnmarshalBinary(b []byte) error {
	var res SmsExtensionResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
