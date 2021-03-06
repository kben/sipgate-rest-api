// Code generated by go-swagger; DO NOT EDIT.

package faxlines

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/kben/sipgate-rest-api/models"
)

// GetFaxlineCallerIDReader is a Reader for the GetFaxlineCallerID structure.
type GetFaxlineCallerIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetFaxlineCallerIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetFaxlineCallerIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGetFaxlineCallerIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetFaxlineCallerIDOK creates a GetFaxlineCallerIDOK with default headers values
func NewGetFaxlineCallerIDOK() *GetFaxlineCallerIDOK {
	return &GetFaxlineCallerIDOK{}
}

/*GetFaxlineCallerIDOK handles this case with default header values.

successful operation
*/
type GetFaxlineCallerIDOK struct {
	Payload *models.FaxlineCallerIDResponse
}

func (o *GetFaxlineCallerIDOK) Error() string {
	return fmt.Sprintf("[GET /{userId}/faxlines/{faxlineId}/callerid][%d] getFaxlineCallerIdOK  %+v", 200, o.Payload)
}

func (o *GetFaxlineCallerIDOK) GetPayload() *models.FaxlineCallerIDResponse {
	return o.Payload
}

func (o *GetFaxlineCallerIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FaxlineCallerIDResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFaxlineCallerIDForbidden creates a GetFaxlineCallerIDForbidden with default headers values
func NewGetFaxlineCallerIDForbidden() *GetFaxlineCallerIDForbidden {
	return &GetFaxlineCallerIDForbidden{}
}

/*GetFaxlineCallerIDForbidden handles this case with default header values.

Feature not booked
*/
type GetFaxlineCallerIDForbidden struct {
}

func (o *GetFaxlineCallerIDForbidden) Error() string {
	return fmt.Sprintf("[GET /{userId}/faxlines/{faxlineId}/callerid][%d] getFaxlineCallerIdForbidden ", 403)
}

func (o *GetFaxlineCallerIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
