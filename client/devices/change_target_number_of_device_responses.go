// Code generated by go-swagger; DO NOT EDIT.

package devices

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// ChangeTargetNumberOfDeviceReader is a Reader for the ChangeTargetNumberOfDevice structure.
type ChangeTargetNumberOfDeviceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ChangeTargetNumberOfDeviceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	result := NewChangeTargetNumberOfDeviceDefault(response.Code())
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	if response.Code()/100 == 2 {
		return result, nil
	}
	return nil, result
}

// NewChangeTargetNumberOfDeviceDefault creates a ChangeTargetNumberOfDeviceDefault with default headers values
func NewChangeTargetNumberOfDeviceDefault(code int) *ChangeTargetNumberOfDeviceDefault {
	return &ChangeTargetNumberOfDeviceDefault{
		_statusCode: code,
	}
}

/*ChangeTargetNumberOfDeviceDefault handles this case with default header values.

successful operation
*/
type ChangeTargetNumberOfDeviceDefault struct {
	_statusCode int
}

// Code gets the status code for the change target number of device default response
func (o *ChangeTargetNumberOfDeviceDefault) Code() int {
	return o._statusCode
}

func (o *ChangeTargetNumberOfDeviceDefault) Error() string {
	return fmt.Sprintf("[PUT /devices/{deviceId}/external/targetnumber][%d] changeTargetNumberOfDevice default ", o._statusCode)
}

func (o *ChangeTargetNumberOfDeviceDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}