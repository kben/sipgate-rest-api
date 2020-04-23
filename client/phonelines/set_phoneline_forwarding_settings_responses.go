// Code generated by go-swagger; DO NOT EDIT.

package phonelines

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// SetPhonelineForwardingSettingsReader is a Reader for the SetPhonelineForwardingSettings structure.
type SetPhonelineForwardingSettingsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SetPhonelineForwardingSettingsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 400:
		result := NewSetPhonelineForwardingSettingsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewSetPhonelineForwardingSettingsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewSetPhonelineForwardingSettingsBadRequest creates a SetPhonelineForwardingSettingsBadRequest with default headers values
func NewSetPhonelineForwardingSettingsBadRequest() *SetPhonelineForwardingSettingsBadRequest {
	return &SetPhonelineForwardingSettingsBadRequest{}
}

/*SetPhonelineForwardingSettingsBadRequest handles this case with default header values.

Forwarding settings could not be changed due to bad request
*/
type SetPhonelineForwardingSettingsBadRequest struct {
}

func (o *SetPhonelineForwardingSettingsBadRequest) Error() string {
	return fmt.Sprintf("[PUT /{userId}/phonelines/{phonelineId}/forwardings][%d] setPhonelineForwardingSettingsBadRequest ", 400)
}

func (o *SetPhonelineForwardingSettingsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSetPhonelineForwardingSettingsForbidden creates a SetPhonelineForwardingSettingsForbidden with default headers values
func NewSetPhonelineForwardingSettingsForbidden() *SetPhonelineForwardingSettingsForbidden {
	return &SetPhonelineForwardingSettingsForbidden{}
}

/*SetPhonelineForwardingSettingsForbidden handles this case with default header values.

Feature not booked
*/
type SetPhonelineForwardingSettingsForbidden struct {
}

func (o *SetPhonelineForwardingSettingsForbidden) Error() string {
	return fmt.Sprintf("[PUT /{userId}/phonelines/{phonelineId}/forwardings][%d] setPhonelineForwardingSettingsForbidden ", 403)
}

func (o *SetPhonelineForwardingSettingsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
