// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// EventResponse event response
//
// swagger:model EventResponse
type EventResponse struct {

	// event Id
	EventID string `json:"eventId,omitempty"`

	// event name
	EventName string `json:"eventName,omitempty"`

	// event type
	EventType string `json:"eventType,omitempty"`

	// payload
	Payload map[string]string `json:"payload,omitempty"`
}

// Validate validates this event response
func (m *EventResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *EventResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EventResponse) UnmarshalBinary(b []byte) error {
	var res EventResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
