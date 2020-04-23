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

// GetLocalprefixReader is a Reader for the GetLocalprefix structure.
type GetLocalprefixReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLocalprefixReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLocalprefixOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetLocalprefixOK creates a GetLocalprefixOK with default headers values
func NewGetLocalprefixOK() *GetLocalprefixOK {
	return &GetLocalprefixOK{}
}

/*GetLocalprefixOK handles this case with default header values.

successful operation
*/
type GetLocalprefixOK struct {
	Payload *models.LocalprefixResponse
}

func (o *GetLocalprefixOK) Error() string {
	return fmt.Sprintf("[GET /devices/{deviceId}/localprefix][%d] getLocalprefixOK  %+v", 200, o.Payload)
}

func (o *GetLocalprefixOK) GetPayload() *models.LocalprefixResponse {
	return o.Payload
}

func (o *GetLocalprefixOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LocalprefixResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}