// Code generated by go-swagger; DO NOT EDIT.

package numbers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/kben/sipgate-rest-api/models"
)

// GetUserNumbersReader is a Reader for the GetUserNumbers structure.
type GetUserNumbersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetUserNumbersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetUserNumbersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetUserNumbersOK creates a GetUserNumbersOK with default headers values
func NewGetUserNumbersOK() *GetUserNumbersOK {
	return &GetUserNumbersOK{}
}

/*GetUserNumbersOK handles this case with default header values.

successful operation
*/
type GetUserNumbersOK struct {
	Payload *models.NumbersResponse
}

func (o *GetUserNumbersOK) Error() string {
	return fmt.Sprintf("[GET /{userId}/numbers][%d] getUserNumbersOK  %+v", 200, o.Payload)
}

func (o *GetUserNumbersOK) GetPayload() *models.NumbersResponse {
	return o.Payload
}

func (o *GetUserNumbersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NumbersResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
