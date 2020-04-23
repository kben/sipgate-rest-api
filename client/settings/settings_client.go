// Code generated by go-swagger; DO NOT EDIT.

package settings

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new settings API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for settings API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	GetSipgateIoUrls(params *GetSipgateIoUrlsParams, authInfo runtime.ClientAuthInfoWriter) (*GetSipgateIoUrlsOK, error)

	SetSipgateIoUrls(params *SetSipgateIoUrlsParams, authInfo runtime.ClientAuthInfoWriter) error

	SetTransport(transport runtime.ClientTransport)
}

/*
  GetSipgateIoUrls lists sipgate io global settings
*/
func (a *Client) GetSipgateIoUrls(params *GetSipgateIoUrlsParams, authInfo runtime.ClientAuthInfoWriter) (*GetSipgateIoUrlsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSipgateIoUrlsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getSipgateIoUrls",
		Method:             "GET",
		PathPattern:        "/settings/sipgateio",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetSipgateIoUrlsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetSipgateIoUrlsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getSipgateIoUrls: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  SetSipgateIoUrls updates sipgate io global settings
*/
func (a *Client) SetSipgateIoUrls(params *SetSipgateIoUrlsParams, authInfo runtime.ClientAuthInfoWriter) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSetSipgateIoUrlsParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "setSipgateIoUrls",
		Method:             "PUT",
		PathPattern:        "/settings/sipgateio",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &SetSipgateIoUrlsReader{formats: a.formats},
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
