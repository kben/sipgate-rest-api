// Code generated by go-swagger; DO NOT EDIT.

package addresses

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// ModifyAddressReader is a Reader for the ModifyAddress structure.
type ModifyAddressReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ModifyAddressReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 400:
		result := NewModifyAddressBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewModifyAddressBadRequest creates a ModifyAddressBadRequest with default headers values
func NewModifyAddressBadRequest() *ModifyAddressBadRequest {
	return &ModifyAddressBadRequest{}
}

/*ModifyAddressBadRequest handles this case with default header values.

Bad request
*/
type ModifyAddressBadRequest struct {
}

func (o *ModifyAddressBadRequest) Error() string {
	return fmt.Sprintf("[PUT /addresses/{addressId}][%d] modifyAddressBadRequest ", 400)
}

func (o *ModifyAddressBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
