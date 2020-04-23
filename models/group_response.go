// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// GroupResponse group response
//
// swagger:model GroupResponse
type GroupResponse struct {

	// alias
	Alias string `json:"alias,omitempty"`

	// department Id
	DepartmentID int64 `json:"departmentId,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// voicemails
	Voicemails []*VoicemailSettingsResponse `json:"voicemails"`
}

// Validate validates this group response
func (m *GroupResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateVoicemails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GroupResponse) validateVoicemails(formats strfmt.Registry) error {

	if swag.IsZero(m.Voicemails) { // not required
		return nil
	}

	for i := 0; i < len(m.Voicemails); i++ {
		if swag.IsZero(m.Voicemails[i]) { // not required
			continue
		}

		if m.Voicemails[i] != nil {
			if err := m.Voicemails[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("voicemails" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *GroupResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GroupResponse) UnmarshalBinary(b []byte) error {
	var res GroupResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
