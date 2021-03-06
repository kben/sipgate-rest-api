// Code generated by go-swagger; DO NOT EDIT.

package authorization

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/kben/sipgate-rest-api/models"
)

// GetClientReader is a Reader for the GetClient structure.
type GetClientReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetClientReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetClientOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetClientOK creates a GetClientOK with default headers values
func NewGetClientOK() *GetClientOK {
	return &GetClientOK{}
}

/*GetClientOK handles this case with default header values.

Success
*/
type GetClientOK struct {
	Payload *models.OAuthClient
}

func (o *GetClientOK) Error() string {
	return fmt.Sprintf("[GET /authorization/oauth2/clients/{clientId}][%d] getClientOK  %+v", 200, o.Payload)
}

func (o *GetClientOK) GetPayload() *models.OAuthClient {
	return o.Payload
}

func (o *GetClientOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OAuthClient)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
