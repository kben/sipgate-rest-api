// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// SetUserRoleRequest set user role request
//
// swagger:model SetUserRoleRequest
type SetUserRoleRequest struct {

	// admin
	Admin *bool `json:"admin,omitempty"`
}

// Validate validates this set user role request
func (m *SetUserRoleRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SetUserRoleRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SetUserRoleRequest) UnmarshalBinary(b []byte) error {
	var res SetUserRoleRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}