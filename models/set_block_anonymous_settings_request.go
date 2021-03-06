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

// SetBlockAnonymousSettingsRequest set block anonymous settings request
//
// swagger:model SetBlockAnonymousSettingsRequest
type SetBlockAnonymousSettingsRequest struct {

	// enabled
	Enabled *bool `json:"enabled,omitempty"`

	// target
	// Enum: [REJECT VOICEMAIL]
	Target string `json:"target,omitempty"`
}

// Validate validates this set block anonymous settings request
func (m *SetBlockAnonymousSettingsRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTarget(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var setBlockAnonymousSettingsRequestTypeTargetPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["REJECT","VOICEMAIL"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		setBlockAnonymousSettingsRequestTypeTargetPropEnum = append(setBlockAnonymousSettingsRequestTypeTargetPropEnum, v)
	}
}

const (

	// SetBlockAnonymousSettingsRequestTargetREJECT captures enum value "REJECT"
	SetBlockAnonymousSettingsRequestTargetREJECT string = "REJECT"

	// SetBlockAnonymousSettingsRequestTargetVOICEMAIL captures enum value "VOICEMAIL"
	SetBlockAnonymousSettingsRequestTargetVOICEMAIL string = "VOICEMAIL"
)

// prop value enum
func (m *SetBlockAnonymousSettingsRequest) validateTargetEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, setBlockAnonymousSettingsRequestTypeTargetPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *SetBlockAnonymousSettingsRequest) validateTarget(formats strfmt.Registry) error {

	if swag.IsZero(m.Target) { // not required
		return nil
	}

	// value enum
	if err := m.validateTargetEnum("target", "body", m.Target); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SetBlockAnonymousSettingsRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SetBlockAnonymousSettingsRequest) UnmarshalBinary(b []byte) error {
	var res SetBlockAnonymousSettingsRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
