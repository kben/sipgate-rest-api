// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CredentialsResponse credentials response
//
// swagger:model CredentialsResponse
type CredentialsResponse struct {

	// outbound proxy
	OutboundProxy string `json:"outboundProxy,omitempty"`

	// password
	Password string `json:"password,omitempty"`

	// sip server
	SipServer string `json:"sipServer,omitempty"`

	// sip server websocket Url
	SipServerWebsocketURL string `json:"sipServerWebsocketUrl,omitempty"`

	// username
	Username string `json:"username,omitempty"`
}

// Validate validates this credentials response
func (m *CredentialsResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CredentialsResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CredentialsResponse) UnmarshalBinary(b []byte) error {
	var res CredentialsResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}