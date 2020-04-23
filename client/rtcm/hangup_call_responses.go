// Code generated by go-swagger; DO NOT EDIT.

package rtcm

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// HangupCallReader is a Reader for the HangupCall structure.
type HangupCallReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *HangupCallReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewHangupCallNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewHangupCallNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewHangupCallNoContent creates a HangupCallNoContent with default headers values
func NewHangupCallNoContent() *HangupCallNoContent {
	return &HangupCallNoContent{}
}

/*HangupCallNoContent handles this case with default header values.

Success
*/
type HangupCallNoContent struct {
}

func (o *HangupCallNoContent) Error() string {
	return fmt.Sprintf("[DELETE /calls/{callId}][%d] hangupCallNoContent ", 204)
}

func (o *HangupCallNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewHangupCallNotFound creates a HangupCallNotFound with default headers values
func NewHangupCallNotFound() *HangupCallNotFound {
	return &HangupCallNotFound{}
}

/*HangupCallNotFound handles this case with default header values.

Call not found
*/
type HangupCallNotFound struct {
}

func (o *HangupCallNotFound) Error() string {
	return fmt.Sprintf("[DELETE /calls/{callId}][%d] hangupCallNotFound ", 404)
}

func (o *HangupCallNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
