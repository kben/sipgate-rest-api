// Code generated by go-swagger; DO NOT EDIT.

package contacts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// ImportFromCSVReader is a Reader for the ImportFromCSV structure.
type ImportFromCSVReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ImportFromCSVReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	result := NewImportFromCSVDefault(response.Code())
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	if response.Code()/100 == 2 {
		return result, nil
	}
	return nil, result
}

// NewImportFromCSVDefault creates a ImportFromCSVDefault with default headers values
func NewImportFromCSVDefault(code int) *ImportFromCSVDefault {
	return &ImportFromCSVDefault{
		_statusCode: code,
	}
}

/*ImportFromCSVDefault handles this case with default header values.

successful operation
*/
type ImportFromCSVDefault struct {
	_statusCode int
}

// Code gets the status code for the import from c s v default response
func (o *ImportFromCSVDefault) Code() int {
	return o._statusCode
}

func (o *ImportFromCSVDefault) Error() string {
	return fmt.Sprintf("[POST /contacts/import/csv][%d] importFromCSV default ", o._statusCode)
}

func (o *ImportFromCSVDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
