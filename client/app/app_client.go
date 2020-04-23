// Code generated by go-swagger; DO NOT EDIT.

package app

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new app API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for app API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	DeleteEvent(params *DeleteEventParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteEventOK, error)

	GetEvents(params *GetEventsParams, authInfo runtime.ClientAuthInfoWriter) (*GetEventsOK, error)

	GetLanguage(params *GetLanguageParams) error

	GetLinks(params *GetLinksParams) (*GetLinksOK, error)

	GetProperties(params *GetPropertiesParams, authInfo runtime.ClientAuthInfoWriter) (*GetPropertiesOK, error)

	GetTac(params *GetTacParams, authInfo runtime.ClientAuthInfoWriter) (*GetTacOK, error)

	GetUsers(params *GetUsersParams, authInfo runtime.ClientAuthInfoWriter) (*GetUsersOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  DeleteEvent deletes a specific event
*/
func (a *Client) DeleteEvent(params *DeleteEventParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteEventOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteEventParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteEvent",
		Method:             "DELETE",
		PathPattern:        "/app/events/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteEventReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteEventOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteEvent: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetEvents lists all active events for this account
*/
func (a *Client) GetEvents(params *GetEventsParams, authInfo runtime.ClientAuthInfoWriter) (*GetEventsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetEventsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getEvents",
		Method:             "GET",
		PathPattern:        "/app/events",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetEventsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetEventsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getEvents: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetLanguage lists sipgate translations for a specific language
*/
func (a *Client) GetLanguage(params *GetLanguageParams) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetLanguageParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getLanguage",
		Method:             "GET",
		PathPattern:        "/translations/{language}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetLanguageReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return err
	}
	return nil
}

/*
  GetLinks lists all external u r ls of app sipgate com
*/
func (a *Client) GetLinks(params *GetLinksParams) (*GetLinksOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetLinksParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getLinks",
		Method:             "GET",
		PathPattern:        "/app/links",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetLinksReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetLinksOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getLinks: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetProperties lists all properties of app sipgate com
*/
func (a *Client) GetProperties(params *GetPropertiesParams, authInfo runtime.ClientAuthInfoWriter) (*GetPropertiesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPropertiesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getProperties",
		Method:             "GET",
		PathPattern:        "/app/properties",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetPropertiesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetPropertiesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getProperties: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetTac checks if terms and conditions have been accepted
*/
func (a *Client) GetTac(params *GetTacParams, authInfo runtime.ClientAuthInfoWriter) (*GetTacOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetTacParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getTac",
		Method:             "GET",
		PathPattern:        "/app/tacs",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetTacReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetTacOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getTac: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetUsers fetches all users with numbers addresses and call restrictions
*/
func (a *Client) GetUsers(params *GetUsersParams, authInfo runtime.ClientAuthInfoWriter) (*GetUsersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetUsersParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getUsers",
		Method:             "GET",
		PathPattern:        "/app/users",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetUsersReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetUsersOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getUsers: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
