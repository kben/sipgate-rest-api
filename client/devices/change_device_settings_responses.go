// Code generated by go-swagger; DO NOT EDIT.

package devices

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// ChangeDeviceSettingsReader is a Reader for the ChangeDeviceSettings structure.
type ChangeDeviceSettingsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ChangeDeviceSettingsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	result := NewChangeDeviceSettingsDefault(response.Code())
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	if response.Code()/100 == 2 {
		return result, nil
	}
	return nil, result
}

// NewChangeDeviceSettingsDefault creates a ChangeDeviceSettingsDefault with default headers values
func NewChangeDeviceSettingsDefault(code int) *ChangeDeviceSettingsDefault {
	return &ChangeDeviceSettingsDefault{
		_statusCode: code,
	}
}

/*ChangeDeviceSettingsDefault handles this case with default header values.

successful operation
*/
type ChangeDeviceSettingsDefault struct {
	_statusCode int
}

// Code gets the status code for the change device settings default response
func (o *ChangeDeviceSettingsDefault) Code() int {
	return o._statusCode
}

func (o *ChangeDeviceSettingsDefault) Error() string {
	return fmt.Sprintf("[PUT /devices/{deviceId}][%d] changeDeviceSettings default ", o._statusCode)
}

func (o *ChangeDeviceSettingsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
