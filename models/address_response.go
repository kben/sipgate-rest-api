// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AddressResponse address response
//
// swagger:model AddressResponse
type AddressResponse struct {

	// country
	Country string `json:"country,omitempty"`

	// extended address
	ExtendedAddress string `json:"extendedAddress,omitempty"`

	// locality
	Locality string `json:"locality,omitempty"`

	// po box
	PoBox string `json:"poBox,omitempty"`

	// postal code
	PostalCode string `json:"postalCode,omitempty"`

	// region
	Region string `json:"region,omitempty"`

	// street address
	StreetAddress string `json:"streetAddress,omitempty"`
}

// Validate validates this address response
func (m *AddressResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AddressResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AddressResponse) UnmarshalBinary(b []byte) error {
	var res AddressResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
