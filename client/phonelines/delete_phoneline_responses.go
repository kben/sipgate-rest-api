// Code generated by go-swagger; DO NOT EDIT.

package phonelines

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeletePhonelineReader is a Reader for the DeletePhoneline structure.
type DeletePhonelineReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeletePhonelineReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 404:
		result := NewDeletePhonelineNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeletePhonelineNotFound creates a DeletePhonelineNotFound with default headers values
func NewDeletePhonelineNotFound() *DeletePhonelineNotFound {
	return &DeletePhonelineNotFound{}
}

/*DeletePhonelineNotFound handles this case with default header values.

Phoneline not found
*/
type DeletePhonelineNotFound struct {
}

func (o *DeletePhonelineNotFound) Error() string {
	return fmt.Sprintf("[DELETE /{userId}/phonelines/{phonelineId}][%d] deletePhonelineNotFound ", 404)
}

func (o *DeletePhonelineNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
