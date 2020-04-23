// Code generated by go-swagger; DO NOT EDIT.

package voicemails

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new voicemails API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for voicemails API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	GetVoicemail(params *GetVoicemailParams, authInfo runtime.ClientAuthInfoWriter) (*GetVoicemailOK, error)

	GetVoicemails(params *GetVoicemailsParams, authInfo runtime.ClientAuthInfoWriter) (*GetVoicemailsOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  GetVoicemail gets a single voicemail
*/
func (a *Client) GetVoicemail(params *GetVoicemailParams, authInfo runtime.ClientAuthInfoWriter) (*GetVoicemailOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetVoicemailParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getVoicemail",
		Method:             "GET",
		PathPattern:        "/voicemails/{voicemailId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetVoicemailReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetVoicemailOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getVoicemail: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetVoicemails lists all voicemails
*/
func (a *Client) GetVoicemails(params *GetVoicemailsParams, authInfo runtime.ClientAuthInfoWriter) (*GetVoicemailsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetVoicemailsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getVoicemails",
		Method:             "GET",
		PathPattern:        "/voicemails",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetVoicemailsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetVoicemailsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getVoicemails: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
