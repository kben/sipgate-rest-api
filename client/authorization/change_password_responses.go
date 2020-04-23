// Code generated by go-swagger; DO NOT EDIT.

package authorization

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// ChangePasswordReader is a Reader for the ChangePassword structure.
type ChangePasswordReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ChangePasswordReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	result := NewChangePasswordDefault(response.Code())
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	if response.Code()/100 == 2 {
		return result, nil
	}
	return nil, result
}

// NewChangePasswordDefault creates a ChangePasswordDefault with default headers values
func NewChangePasswordDefault(code int) *ChangePasswordDefault {
	return &ChangePasswordDefault{
		_statusCode: code,
	}
}

/*ChangePasswordDefault handles this case with default header values.

successful operation
*/
type ChangePasswordDefault struct {
	_statusCode int
}

// Code gets the status code for the change password default response
func (o *ChangePasswordDefault) Code() int {
	return o._statusCode
}

func (o *ChangePasswordDefault) Error() string {
	return fmt.Sprintf("[POST /passwordreset/{nonce}][%d] changePassword default ", o._statusCode)
}

func (o *ChangePasswordDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
