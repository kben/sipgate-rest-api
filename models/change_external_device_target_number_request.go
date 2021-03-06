// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ChangeExternalDeviceTargetNumberRequest change external device target number request
//
// swagger:model ChangeExternalDeviceTargetNumberRequest
type ChangeExternalDeviceTargetNumberRequest struct {

	// number
	Number string `json:"number,omitempty"`
}

// Validate validates this change external device target number request
func (m *ChangeExternalDeviceTargetNumberRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ChangeExternalDeviceTargetNumberRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ChangeExternalDeviceTargetNumberRequest) UnmarshalBinary(b []byte) error {
	var res ChangeExternalDeviceTargetNumberRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
