// Code generated by go-swagger; DO NOT EDIT.

package app

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/kben/sipgate-rest-api/models"
)

// GetLinksReader is a Reader for the GetLinks structure.
type GetLinksReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLinksReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLinksOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetLinksOK creates a GetLinksOK with default headers values
func NewGetLinksOK() *GetLinksOK {
	return &GetLinksOK{}
}

/*GetLinksOK handles this case with default header values.

successful operation
*/
type GetLinksOK struct {
	Payload *models.LinksResponse
}

func (o *GetLinksOK) Error() string {
	return fmt.Sprintf("[GET /app/links][%d] getLinksOK  %+v", 200, o.Payload)
}

func (o *GetLinksOK) GetPayload() *models.LinksResponse {
	return o.Payload
}

func (o *GetLinksOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LinksResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}