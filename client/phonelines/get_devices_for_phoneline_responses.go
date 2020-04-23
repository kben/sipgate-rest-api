// Code generated by go-swagger; DO NOT EDIT.

package phonelines

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/kben/sipgate-rest-api/models"
)

// GetDevicesForPhonelineReader is a Reader for the GetDevicesForPhoneline structure.
type GetDevicesForPhonelineReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDevicesForPhonelineReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDevicesForPhonelineOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetDevicesForPhonelineOK creates a GetDevicesForPhonelineOK with default headers values
func NewGetDevicesForPhonelineOK() *GetDevicesForPhonelineOK {
	return &GetDevicesForPhonelineOK{}
}

/*GetDevicesForPhonelineOK handles this case with default header values.

successful operation
*/
type GetDevicesForPhonelineOK struct {
	Payload *models.DevicesResponse
}

func (o *GetDevicesForPhonelineOK) Error() string {
	return fmt.Sprintf("[GET /{userId}/phonelines/{phonelineId}/devices][%d] getDevicesForPhonelineOK  %+v", 200, o.Payload)
}

func (o *GetDevicesForPhonelineOK) GetPayload() *models.DevicesResponse {
	return o.Payload
}

func (o *GetDevicesForPhonelineOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DevicesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}