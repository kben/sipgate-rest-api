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

// AddGreetingRequest add greeting request
//
// swagger:model AddGreetingRequest
type AddGreetingRequest struct {

	// base64 content
	// Max Length: 8000000
	// Min Length: 0
	Base64Content *string `json:"base64Content,omitempty"`

	// filename
	Filename string `json:"filename,omitempty"`
}

// Validate validates this add greeting request
func (m *AddGreetingRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBase64Content(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AddGreetingRequest) validateBase64Content(formats strfmt.Registry) error {

	if swag.IsZero(m.Base64Content) { // not required
		return nil
	}

	if err := validate.MinLength("base64Content", "body", string(*m.Base64Content), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("base64Content", "body", string(*m.Base64Content), 8000000); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AddGreetingRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AddGreetingRequest) UnmarshalBinary(b []byte) error {
	var res AddGreetingRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
