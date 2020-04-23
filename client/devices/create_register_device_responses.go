// Code generated by go-swagger; DO NOT EDIT.

package devices

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/kben/sipgate-rest-api/models"
)

// CreateRegisterDeviceReader is a Reader for the CreateRegisterDevice structure.
type CreateRegisterDeviceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateRegisterDeviceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateRegisterDeviceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateRegisterDeviceOK creates a CreateRegisterDeviceOK with default headers values
func NewCreateRegisterDeviceOK() *CreateRegisterDeviceOK {
	return &CreateRegisterDeviceOK{}
}

/*CreateRegisterDeviceOK handles this case with default header values.

successful operation
*/
type CreateRegisterDeviceOK struct {
	Payload *models.RegisterDeviceResponse
}

func (o *CreateRegisterDeviceOK) Error() string {
	return fmt.Sprintf("[POST /{userId}/devices/register][%d] createRegisterDeviceOK  %+v", 200, o.Payload)
}

func (o *CreateRegisterDeviceOK) GetPayload() *models.RegisterDeviceResponse {
	return o.Payload
}

func (o *CreateRegisterDeviceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RegisterDeviceResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}