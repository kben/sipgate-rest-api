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

// GetPhonelineForwardingsReader is a Reader for the GetPhonelineForwardings structure.
type GetPhonelineForwardingsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPhonelineForwardingsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetPhonelineForwardingsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGetPhonelineForwardingsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetPhonelineForwardingsOK creates a GetPhonelineForwardingsOK with default headers values
func NewGetPhonelineForwardingsOK() *GetPhonelineForwardingsOK {
	return &GetPhonelineForwardingsOK{}
}

/*GetPhonelineForwardingsOK handles this case with default header values.

successful operation
*/
type GetPhonelineForwardingsOK struct {
	Payload *models.ForwardingsResponse
}

func (o *GetPhonelineForwardingsOK) Error() string {
	return fmt.Sprintf("[GET /{userId}/phonelines/{phonelineId}/forwardings][%d] getPhonelineForwardingsOK  %+v", 200, o.Payload)
}

func (o *GetPhonelineForwardingsOK) GetPayload() *models.ForwardingsResponse {
	return o.Payload
}

func (o *GetPhonelineForwardingsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ForwardingsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPhonelineForwardingsForbidden creates a GetPhonelineForwardingsForbidden with default headers values
func NewGetPhonelineForwardingsForbidden() *GetPhonelineForwardingsForbidden {
	return &GetPhonelineForwardingsForbidden{}
}

/*GetPhonelineForwardingsForbidden handles this case with default header values.

Feature not booked
*/
type GetPhonelineForwardingsForbidden struct {
}

func (o *GetPhonelineForwardingsForbidden) Error() string {
	return fmt.Sprintf("[GET /{userId}/phonelines/{phonelineId}/forwardings][%d] getPhonelineForwardingsForbidden ", 403)
}

func (o *GetPhonelineForwardingsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
