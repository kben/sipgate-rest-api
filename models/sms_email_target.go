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

// SmsEmailTarget sms email target
//
// swagger:model SmsEmailTarget
type SmsEmailTarget struct {

	// direction
	// Enum: [INCOMING OUTGOING]
	Direction string `json:"direction,omitempty"`

	// email
	Email string `json:"email,omitempty"`

	// id
	ID string `json:"id,omitempty"`
}

// Validate validates this sms email target
func (m *SmsEmailTarget) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDirection(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var smsEmailTargetTypeDirectionPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["INCOMING","OUTGOING"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		smsEmailTargetTypeDirectionPropEnum = append(smsEmailTargetTypeDirectionPropEnum, v)
	}
}

const (

	// SmsEmailTargetDirectionINCOMING captures enum value "INCOMING"
	SmsEmailTargetDirectionINCOMING string = "INCOMING"

	// SmsEmailTargetDirectionOUTGOING captures enum value "OUTGOING"
	SmsEmailTargetDirectionOUTGOING string = "OUTGOING"
)

// prop value enum
func (m *SmsEmailTarget) validateDirectionEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, smsEmailTargetTypeDirectionPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *SmsEmailTarget) validateDirection(formats strfmt.Registry) error {

	if swag.IsZero(m.Direction) { // not required
		return nil
	}

	// value enum
	if err := m.validateDirectionEnum("direction", "body", m.Direction); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SmsEmailTarget) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SmsEmailTarget) UnmarshalBinary(b []byte) error {
	var res SmsEmailTarget
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}