// Code generated by go-swagger; DO NOT EDIT.

package history

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// UpdateHistoryEntriesReader is a Reader for the UpdateHistoryEntries structure.
type UpdateHistoryEntriesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateHistoryEntriesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	result := NewUpdateHistoryEntriesDefault(response.Code())
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	if response.Code()/100 == 2 {
		return result, nil
	}
	return nil, result
}

// NewUpdateHistoryEntriesDefault creates a UpdateHistoryEntriesDefault with default headers values
func NewUpdateHistoryEntriesDefault(code int) *UpdateHistoryEntriesDefault {
	return &UpdateHistoryEntriesDefault{
		_statusCode: code,
	}
}

/*UpdateHistoryEntriesDefault handles this case with default header values.

successful operation
*/
type UpdateHistoryEntriesDefault struct {
	_statusCode int
}

// Code gets the status code for the update history entries default response
func (o *UpdateHistoryEntriesDefault) Code() int {
	return o._statusCode
}

func (o *UpdateHistoryEntriesDefault) Error() string {
	return fmt.Sprintf("[PUT /history][%d] updateHistoryEntries default ", o._statusCode)
}

func (o *UpdateHistoryEntriesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
