// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CreateExternalDeviceRequest create external device request
//
// swagger:model CreateExternalDeviceRequest
type CreateExternalDeviceRequest struct {

	// alias
	Alias string `json:"alias,omitempty"`

	// number
	Number string `json:"number,omitempty"`
}

// Validate validates this create external device request
func (m *CreateExternalDeviceRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CreateExternalDeviceRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateExternalDeviceRequest) UnmarshalBinary(b []byte) error {
	var res CreateExternalDeviceRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}