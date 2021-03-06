// Code generated by go-swagger; DO NOT EDIT.

package addresses

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new addresses API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for addresses API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	GetAddress(params *GetAddressParams, authInfo runtime.ClientAuthInfoWriter) (*GetAddressOK, error)

	GetAddressNumbers(params *GetAddressNumbersParams, authInfo runtime.ClientAuthInfoWriter) (*GetAddressNumbersOK, error)

	GetAddresses(params *GetAddressesParams, authInfo runtime.ClientAuthInfoWriter) (*GetAddressesOK, error)

	GetItemCount(params *GetItemCountParams, authInfo runtime.ClientAuthInfoWriter) (*GetItemCountOK, error)

	ModifyAddress(params *ModifyAddressParams, authInfo runtime.ClientAuthInfoWriter) error

	SetTransport(transport runtime.ClientTransport)
}

/*
  GetAddress gets a specific address
*/
func (a *Client) GetAddress(params *GetAddressParams, authInfo runtime.ClientAuthInfoWriter) (*GetAddressOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAddressParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getAddress",
		Method:             "GET",
		PathPattern:        "/addresses/{addressId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetAddressReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetAddressOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getAddress: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetAddressNumbers lists all phone numbers of a specific address
*/
func (a *Client) GetAddressNumbers(params *GetAddressNumbersParams, authInfo runtime.ClientAuthInfoWriter) (*GetAddressNumbersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAddressNumbersParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getAddressNumbers",
		Method:             "GET",
		PathPattern:        "/addresses/{addressId}/numbers",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetAddressNumbersReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetAddressNumbersOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getAddressNumbers: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetAddresses lists all addresses
*/
func (a *Client) GetAddresses(params *GetAddressesParams, authInfo runtime.ClientAuthInfoWriter) (*GetAddressesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAddressesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getAddresses",
		Method:             "GET",
		PathPattern:        "/addresses",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetAddressesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetAddressesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getAddresses: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetItemCount gets amount of addresses
*/
func (a *Client) GetItemCount(params *GetItemCountParams, authInfo runtime.ClientAuthInfoWriter) (*GetItemCountOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetItemCountParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getItemCount",
		Method:             "GET",
		PathPattern:        "/addresses/itemCount",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetItemCountReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetItemCountOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getItemCount: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ModifyAddress updates an existing address

  <strong>Warning:</strong> Depending on your country of residence, a change of your address may lead to the deactivation of associated telephone numbers.
*/
func (a *Client) ModifyAddress(params *ModifyAddressParams, authInfo runtime.ClientAuthInfoWriter) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewModifyAddressParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "modifyAddress",
		Method:             "PUT",
		PathPattern:        "/addresses/{addressId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ModifyAddressReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return err
	}
	return nil
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
