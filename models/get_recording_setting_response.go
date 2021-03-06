// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// GetRecordingSettingResponse get recording setting response
//
// swagger:model GetRecordingSettingResponse
type GetRecordingSettingResponse struct {

	// active
	Active *bool `json:"active,omitempty"`
}

// Validate validates this get recording setting response
func (m *GetRecordingSettingResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GetRecordingSettingResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetRecordingSettingResponse) UnmarshalBinary(b []byte) error {
	var res GetRecordingSettingResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
