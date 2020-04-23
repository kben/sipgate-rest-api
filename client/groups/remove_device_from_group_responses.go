// Code generated by go-swagger; DO NOT EDIT.

package groups

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// RemoveDeviceFromGroupReader is a Reader for the RemoveDeviceFromGroup structure.
type RemoveDeviceFromGroupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RemoveDeviceFromGroupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	result := NewRemoveDeviceFromGroupDefault(response.Code())
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	if response.Code()/100 == 2 {
		return result, nil
	}
	return nil, result
}

// NewRemoveDeviceFromGroupDefault creates a RemoveDeviceFromGroupDefault with default headers values
func NewRemoveDeviceFromGroupDefault(code int) *RemoveDeviceFromGroupDefault {
	return &RemoveDeviceFromGroupDefault{
		_statusCode: code,
	}
}

/*RemoveDeviceFromGroupDefault handles this case with default header values.

successful operation
*/
type RemoveDeviceFromGroupDefault struct {
	_statusCode int
}

// Code gets the status code for the remove device from group default response
func (o *RemoveDeviceFromGroupDefault) Code() int {
	return o._statusCode
}

func (o *RemoveDeviceFromGroupDefault) Error() string {
	return fmt.Sprintf("[DELETE /groups/{groupId}/devices/{deviceId}][%d] removeDeviceFromGroup default ", o._statusCode)
}

func (o *RemoveDeviceFromGroupDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
