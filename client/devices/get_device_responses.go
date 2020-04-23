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

// GetDeviceReader is a Reader for the GetDevice structure.
type GetDeviceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDeviceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDeviceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetDeviceOK creates a GetDeviceOK with default headers values
func NewGetDeviceOK() *GetDeviceOK {
	return &GetDeviceOK{}
}

/*GetDeviceOK handles this case with default header values.

successful operation
*/
type GetDeviceOK struct {
	Payload *models.DeviceResponse
}

func (o *GetDeviceOK) Error() string {
	return fmt.Sprintf("[GET /devices/{deviceId}][%d] getDeviceOK  %+v", 200, o.Payload)
}

func (o *GetDeviceOK) GetPayload() *models.DeviceResponse {
	return o.Payload
}

func (o *GetDeviceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DeviceResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
