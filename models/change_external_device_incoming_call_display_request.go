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

// ChangeExternalDeviceIncomingCallDisplayRequest change external device incoming call display request
//
// swagger:model ChangeExternalDeviceIncomingCallDisplayRequest
type ChangeExternalDeviceIncomingCallDisplayRequest struct {

	// incoming call display
	// Required: true
	// Enum: [CALLED_NUMBER CALLER_NUMBER]
	IncomingCallDisplay *string `json:"incomingCallDisplay"`
}

// Validate validates this change external device incoming call display request
func (m *ChangeExternalDeviceIncomingCallDisplayRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateIncomingCallDisplay(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var changeExternalDeviceIncomingCallDisplayRequestTypeIncomingCallDisplayPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["CALLED_NUMBER","CALLER_NUMBER"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		changeExternalDeviceIncomingCallDisplayRequestTypeIncomingCallDisplayPropEnum = append(changeExternalDeviceIncomingCallDisplayRequestTypeIncomingCallDisplayPropEnum, v)
	}
}

const (

	// ChangeExternalDeviceIncomingCallDisplayRequestIncomingCallDisplayCALLEDNUMBER captures enum value "CALLED_NUMBER"
	ChangeExternalDeviceIncomingCallDisplayRequestIncomingCallDisplayCALLEDNUMBER string = "CALLED_NUMBER"

	// ChangeExternalDeviceIncomingCallDisplayRequestIncomingCallDisplayCALLERNUMBER captures enum value "CALLER_NUMBER"
	ChangeExternalDeviceIncomingCallDisplayRequestIncomingCallDisplayCALLERNUMBER string = "CALLER_NUMBER"
)

// prop value enum
func (m *ChangeExternalDeviceIncomingCallDisplayRequest) validateIncomingCallDisplayEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, changeExternalDeviceIncomingCallDisplayRequestTypeIncomingCallDisplayPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *ChangeExternalDeviceIncomingCallDisplayRequest) validateIncomingCallDisplay(formats strfmt.Registry) error {

	if err := validate.Required("incomingCallDisplay", "body", m.IncomingCallDisplay); err != nil {
		return err
	}

	// value enum
	if err := m.validateIncomingCallDisplayEnum("incomingCallDisplay", "body", *m.IncomingCallDisplay); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ChangeExternalDeviceIncomingCallDisplayRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ChangeExternalDeviceIncomingCallDisplayRequest) UnmarshalBinary(b []byte) error {
	var res ChangeExternalDeviceIncomingCallDisplayRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
