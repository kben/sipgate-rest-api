// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ErrorSource A more detailed info where the error occured
//
// swagger:model ErrorSource
type ErrorSource struct {

	// A pointer to indicate where the error occured
	Pointer string `json:"pointer,omitempty"`
}

// Validate validates this error source
func (m *ErrorSource) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ErrorSource) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ErrorSource) UnmarshalBinary(b []byte) error {
	var res ErrorSource
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}