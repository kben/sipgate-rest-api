// Code generated by go-swagger; DO NOT EDIT.

package faxlines

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// ChangeFaxlineCallerIDSettingsReader is a Reader for the ChangeFaxlineCallerIDSettings structure.
type ChangeFaxlineCallerIDSettingsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ChangeFaxlineCallerIDSettingsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 403:
		result := NewChangeFaxlineCallerIDSettingsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewChangeFaxlineCallerIDSettingsForbidden creates a ChangeFaxlineCallerIDSettingsForbidden with default headers values
func NewChangeFaxlineCallerIDSettingsForbidden() *ChangeFaxlineCallerIDSettingsForbidden {
	return &ChangeFaxlineCallerIDSettingsForbidden{}
}

/*ChangeFaxlineCallerIDSettingsForbidden handles this case with default header values.

Feature not booked
*/
type ChangeFaxlineCallerIDSettingsForbidden struct {
}

func (o *ChangeFaxlineCallerIDSettingsForbidden) Error() string {
	return fmt.Sprintf("[PUT /{userId}/faxlines/{faxlineId}/callerid][%d] changeFaxlineCallerIdSettingsForbidden ", 403)
}

func (o *ChangeFaxlineCallerIDSettingsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}