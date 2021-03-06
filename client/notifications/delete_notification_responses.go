// Code generated by go-swagger; DO NOT EDIT.

package notifications

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeleteNotificationReader is a Reader for the DeleteNotification structure.
type DeleteNotificationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteNotificationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	result := NewDeleteNotificationDefault(response.Code())
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	if response.Code()/100 == 2 {
		return result, nil
	}
	return nil, result
}

// NewDeleteNotificationDefault creates a DeleteNotificationDefault with default headers values
func NewDeleteNotificationDefault(code int) *DeleteNotificationDefault {
	return &DeleteNotificationDefault{
		_statusCode: code,
	}
}

/*DeleteNotificationDefault handles this case with default header values.

successful operation
*/
type DeleteNotificationDefault struct {
	_statusCode int
}

// Code gets the status code for the delete notification default response
func (o *DeleteNotificationDefault) Code() int {
	return o._statusCode
}

func (o *DeleteNotificationDefault) Error() string {
	return fmt.Sprintf("[DELETE /{userId}/notifications/{notificationId}][%d] deleteNotification default ", o._statusCode)
}

func (o *DeleteNotificationDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
