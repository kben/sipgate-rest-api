// Code generated by go-swagger; DO NOT EDIT.

package callrestrictions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new callrestrictions API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for callrestrictions API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	GetCallRestrictions(params *GetCallRestrictionsParams, authInfo runtime.ClientAuthInfoWriter) (*GetCallRestrictionsOK, error)

	SetCallRestriction(params *SetCallRestrictionParams, authInfo runtime.ClientAuthInfoWriter) (*SetCallRestrictionOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  GetCallRestrictions gets call restrictions data
*/
func (a *Client) GetCallRestrictions(params *GetCallRestrictionsParams, authInfo runtime.ClientAuthInfoWriter) (*GetCallRestrictionsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCallRestrictionsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getCallRestrictions",
		Method:             "GET",
		PathPattern:        "/callrestrictions",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetCallRestrictionsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetCallRestrictionsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getCallRestrictions: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  SetCallRestriction enables disable restriction
*/
func (a *Client) SetCallRestriction(params *SetCallRestrictionParams, authInfo runtime.ClientAuthInfoWriter) (*SetCallRestrictionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSetCallRestrictionParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "setCallRestriction",
		Method:             "POST",
		PathPattern:        "/{userId}/callrestrictions/{restriction}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &SetCallRestrictionReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*SetCallRestrictionOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for setCallRestriction: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
