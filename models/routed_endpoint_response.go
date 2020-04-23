// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// RoutedEndpointResponse routed endpoint response
//
// swagger:model RoutedEndpointResponse
type RoutedEndpointResponse struct {

	// endpoint
	Endpoint *EndpointResponse `json:"endpoint,omitempty"`

	// type
	// Enum: [FORWARDED ROUTED PICKUP]
	Type string `json:"type,omitempty"`
}

// Validate validates this routed endpoint response
func (m *RoutedEndpointResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEndpoint(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RoutedEndpointResponse) validateEndpoint(formats strfmt.Registry) error {

	if swag.IsZero(m.Endpoint) { // not required
		return nil
	}

	if m.Endpoint != nil {
		if err := m.Endpoint.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("endpoint")
			}
			return err
		}
	}

	return nil
}

var routedEndpointResponseTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["FORWARDED","ROUTED","PICKUP"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		routedEndpointResponseTypeTypePropEnum = append(routedEndpointResponseTypeTypePropEnum, v)
	}
}

const (

	// RoutedEndpointResponseTypeFORWARDED captures enum value "FORWARDED"
	RoutedEndpointResponseTypeFORWARDED string = "FORWARDED"

	// RoutedEndpointResponseTypeROUTED captures enum value "ROUTED"
	RoutedEndpointResponseTypeROUTED string = "ROUTED"

	// RoutedEndpointResponseTypePICKUP captures enum value "PICKUP"
	RoutedEndpointResponseTypePICKUP string = "PICKUP"
)

// prop value enum
func (m *RoutedEndpointResponse) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, routedEndpointResponseTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *RoutedEndpointResponse) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RoutedEndpointResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RoutedEndpointResponse) UnmarshalBinary(b []byte) error {
	var res RoutedEndpointResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}