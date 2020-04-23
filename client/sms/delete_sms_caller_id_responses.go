// Code generated by go-swagger; DO NOT EDIT.

package sms

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeleteSmsCallerIDReader is a Reader for the DeleteSmsCallerID structure.
type DeleteSmsCallerIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteSmsCallerIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	result := NewDeleteSmsCallerIDDefault(response.Code())
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	if response.Code()/100 == 2 {
		return result, nil
	}
	return nil, result
}

// NewDeleteSmsCallerIDDefault creates a DeleteSmsCallerIDDefault with default headers values
func NewDeleteSmsCallerIDDefault(code int) *DeleteSmsCallerIDDefault {
	return &DeleteSmsCallerIDDefault{
		_statusCode: code,
	}
}

/*DeleteSmsCallerIDDefault handles this case with default header values.

successful operation
*/
type DeleteSmsCallerIDDefault struct {
	_statusCode int
}

// Code gets the status code for the delete sms caller Id default response
func (o *DeleteSmsCallerIDDefault) Code() int {
	return o._statusCode
}

func (o *DeleteSmsCallerIDDefault) Error() string {
	return fmt.Sprintf("[DELETE /{userId}/sms/{smsId}/callerids/{callerId}][%d] deleteSmsCallerId default ", o._statusCode)
}

func (o *DeleteSmsCallerIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
