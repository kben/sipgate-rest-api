// Code generated by go-swagger; DO NOT EDIT.

package phonelines

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// RemoveDeviceFromPhonelineReader is a Reader for the RemoveDeviceFromPhoneline structure.
type RemoveDeviceFromPhonelineReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RemoveDeviceFromPhonelineReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	result := NewRemoveDeviceFromPhonelineDefault(response.Code())
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	if response.Code()/100 == 2 {
		return result, nil
	}
	return nil, result
}

// NewRemoveDeviceFromPhonelineDefault creates a RemoveDeviceFromPhonelineDefault with default headers values
func NewRemoveDeviceFromPhonelineDefault(code int) *RemoveDeviceFromPhonelineDefault {
	return &RemoveDeviceFromPhonelineDefault{
		_statusCode: code,
	}
}

/*RemoveDeviceFromPhonelineDefault handles this case with default header values.

successful operation
*/
type RemoveDeviceFromPhonelineDefault struct {
	_statusCode int
}

// Code gets the status code for the remove device from phoneline default response
func (o *RemoveDeviceFromPhonelineDefault) Code() int {
	return o._statusCode
}

func (o *RemoveDeviceFromPhonelineDefault) Error() string {
	return fmt.Sprintf("[DELETE /{userId}/phonelines/{phonelineId}/devices/{deviceId}][%d] removeDeviceFromPhoneline default ", o._statusCode)
}

func (o *RemoveDeviceFromPhonelineDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
