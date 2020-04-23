// Code generated by go-swagger; DO NOT EDIT.

package phonelines

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeleteParallelForwardingSettingsReader is a Reader for the DeleteParallelForwardingSettings structure.
type DeleteParallelForwardingSettingsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteParallelForwardingSettingsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 404:
		result := NewDeleteParallelForwardingSettingsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteParallelForwardingSettingsNotFound creates a DeleteParallelForwardingSettingsNotFound with default headers values
func NewDeleteParallelForwardingSettingsNotFound() *DeleteParallelForwardingSettingsNotFound {
	return &DeleteParallelForwardingSettingsNotFound{}
}

/*DeleteParallelForwardingSettingsNotFound handles this case with default header values.

Parallel forwarding not found
*/
type DeleteParallelForwardingSettingsNotFound struct {
}

func (o *DeleteParallelForwardingSettingsNotFound) Error() string {
	return fmt.Sprintf("[DELETE /{userId}/phonelines/{phonelineId}/parallelforwardings/{parallelForwardingId}][%d] deleteParallelForwardingSettingsNotFound ", 404)
}

func (o *DeleteParallelForwardingSettingsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
