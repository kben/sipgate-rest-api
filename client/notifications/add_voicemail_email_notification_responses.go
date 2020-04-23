// Code generated by go-swagger; DO NOT EDIT.

package notifications

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// AddVoicemailEmailNotificationReader is a Reader for the AddVoicemailEmailNotification structure.
type AddVoicemailEmailNotificationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddVoicemailEmailNotificationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 403:
		result := NewAddVoicemailEmailNotificationForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewAddVoicemailEmailNotificationForbidden creates a AddVoicemailEmailNotificationForbidden with default headers values
func NewAddVoicemailEmailNotificationForbidden() *AddVoicemailEmailNotificationForbidden {
	return &AddVoicemailEmailNotificationForbidden{}
}

/*AddVoicemailEmailNotificationForbidden handles this case with default header values.

Voicemail feature not booked
*/
type AddVoicemailEmailNotificationForbidden struct {
}

func (o *AddVoicemailEmailNotificationForbidden) Error() string {
	return fmt.Sprintf("[POST /{userId}/notifications/voicemail/email][%d] addVoicemailEmailNotificationForbidden ", 403)
}

func (o *AddVoicemailEmailNotificationForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
