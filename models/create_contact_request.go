// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// CreateContactRequest create contact request
//
// swagger:model CreateContactRequest
type CreateContactRequest struct {

	// addresses
	Addresses []*AddressRequest `json:"addresses"`

	// emails
	Emails []*EmailRequest `json:"emails"`

	// family
	Family string `json:"family,omitempty"`

	// given
	Given string `json:"given,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// numbers
	Numbers []*NumberRequest `json:"numbers"`

	// organization
	Organization [][]string `json:"organization"`

	// picture
	Picture string `json:"picture,omitempty"`

	// scope
	// Required: true
	// Enum: [PRIVATE SHARED]
	Scope *string `json:"scope"`
}

// Validate validates this create contact request
func (m *CreateContactRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAddresses(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEmails(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNumbers(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateScope(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateContactRequest) validateAddresses(formats strfmt.Registry) error {

	if swag.IsZero(m.Addresses) { // not required
		return nil
	}

	for i := 0; i < len(m.Addresses); i++ {
		if swag.IsZero(m.Addresses[i]) { // not required
			continue
		}

		if m.Addresses[i] != nil {
			if err := m.Addresses[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("addresses" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *CreateContactRequest) validateEmails(formats strfmt.Registry) error {

	if swag.IsZero(m.Emails) { // not required
		return nil
	}

	for i := 0; i < len(m.Emails); i++ {
		if swag.IsZero(m.Emails[i]) { // not required
			continue
		}

		if m.Emails[i] != nil {
			if err := m.Emails[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("emails" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *CreateContactRequest) validateNumbers(formats strfmt.Registry) error {

	if swag.IsZero(m.Numbers) { // not required
		return nil
	}

	for i := 0; i < len(m.Numbers); i++ {
		if swag.IsZero(m.Numbers[i]) { // not required
			continue
		}

		if m.Numbers[i] != nil {
			if err := m.Numbers[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("numbers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

var createContactRequestTypeScopePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["PRIVATE","SHARED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		createContactRequestTypeScopePropEnum = append(createContactRequestTypeScopePropEnum, v)
	}
}

const (

	// CreateContactRequestScopePRIVATE captures enum value "PRIVATE"
	CreateContactRequestScopePRIVATE string = "PRIVATE"

	// CreateContactRequestScopeSHARED captures enum value "SHARED"
	CreateContactRequestScopeSHARED string = "SHARED"
)

// prop value enum
func (m *CreateContactRequest) validateScopeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, createContactRequestTypeScopePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *CreateContactRequest) validateScope(formats strfmt.Registry) error {

	if err := validate.Required("scope", "body", m.Scope); err != nil {
		return err
	}

	// value enum
	if err := m.validateScopeEnum("scope", "body", *m.Scope); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CreateContactRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateContactRequest) UnmarshalBinary(b []byte) error {
	var res CreateContactRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
