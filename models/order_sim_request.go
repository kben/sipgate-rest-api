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

// OrderSimRequest order sim request
//
// swagger:model OrderSimRequest
type OrderSimRequest struct {

	// address Id
	// Required: true
	AddressID *string `json:"addressId"`
}

// Validate validates this order sim request
func (m *OrderSimRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAddressID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OrderSimRequest) validateAddressID(formats strfmt.Registry) error {

	if err := validate.Required("addressId", "body", m.AddressID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *OrderSimRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OrderSimRequest) UnmarshalBinary(b []byte) error {
	var res OrderSimRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
