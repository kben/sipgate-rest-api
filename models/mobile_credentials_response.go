// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// MobileCredentialsResponse mobile credentials response
//
// swagger:model MobileCredentialsResponse
type MobileCredentialsResponse struct {

	// puk1
	Puk1 string `json:"puk1,omitempty"`

	// puk2
	Puk2 string `json:"puk2,omitempty"`

	// sim Id
	SimID string `json:"simId,omitempty"`
}

// Validate validates this mobile credentials response
func (m *MobileCredentialsResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *MobileCredentialsResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MobileCredentialsResponse) UnmarshalBinary(b []byte) error {
	var res MobileCredentialsResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
