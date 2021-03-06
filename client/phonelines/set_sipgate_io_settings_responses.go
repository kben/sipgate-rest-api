// Code generated by go-swagger; DO NOT EDIT.

package phonelines

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// SetSipgateIoSettingsReader is a Reader for the SetSipgateIoSettings structure.
type SetSipgateIoSettingsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SetSipgateIoSettingsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 403:
		result := NewSetSipgateIoSettingsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewSetSipgateIoSettingsForbidden creates a SetSipgateIoSettingsForbidden with default headers values
func NewSetSipgateIoSettingsForbidden() *SetSipgateIoSettingsForbidden {
	return &SetSipgateIoSettingsForbidden{}
}

/*SetSipgateIoSettingsForbidden handles this case with default header values.

sipgate.io feature not booked
*/
type SetSipgateIoSettingsForbidden struct {
}

func (o *SetSipgateIoSettingsForbidden) Error() string {
	return fmt.Sprintf("[PUT /{userId}/phonelines/{phonelineId}/sipgateio][%d] setSipgateIoSettingsForbidden ", 403)
}

func (o *SetSipgateIoSettingsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
