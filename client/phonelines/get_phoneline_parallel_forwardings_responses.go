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

// GetPhonelineParallelForwardingsReader is a Reader for the GetPhonelineParallelForwardings structure.
type GetPhonelineParallelForwardingsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPhonelineParallelForwardingsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetPhonelineParallelForwardingsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGetPhonelineParallelForwardingsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetPhonelineParallelForwardingsOK creates a GetPhonelineParallelForwardingsOK with default headers values
func NewGetPhonelineParallelForwardingsOK() *GetPhonelineParallelForwardingsOK {
	return &GetPhonelineParallelForwardingsOK{}
}

/*GetPhonelineParallelForwardingsOK handles this case with default header values.

successful operation
*/
type GetPhonelineParallelForwardingsOK struct {
	Payload *models.ParallelForwardingsResponse
}

func (o *GetPhonelineParallelForwardingsOK) Error() string {
	return fmt.Sprintf("[GET /{userId}/phonelines/{phonelineId}/parallelforwardings][%d] getPhonelineParallelForwardingsOK  %+v", 200, o.Payload)
}

func (o *GetPhonelineParallelForwardingsOK) GetPayload() *models.ParallelForwardingsResponse {
	return o.Payload
}

func (o *GetPhonelineParallelForwardingsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ParallelForwardingsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPhonelineParallelForwardingsForbidden creates a GetPhonelineParallelForwardingsForbidden with default headers values
func NewGetPhonelineParallelForwardingsForbidden() *GetPhonelineParallelForwardingsForbidden {
	return &GetPhonelineParallelForwardingsForbidden{}
}

/*GetPhonelineParallelForwardingsForbidden handles this case with default header values.

Feature not booked
*/
type GetPhonelineParallelForwardingsForbidden struct {
}

func (o *GetPhonelineParallelForwardingsForbidden) Error() string {
	return fmt.Sprintf("[GET /{userId}/phonelines/{phonelineId}/parallelforwardings][%d] getPhonelineParallelForwardingsForbidden ", 403)
}

func (o *GetPhonelineParallelForwardingsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}