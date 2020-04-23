// Code generated by go-swagger; DO NOT EDIT.

package authorization

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new authorization API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for authorization API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	ChangePassword(params *ChangePasswordParams) error

	CreateClient(params *CreateClientParams, authInfo runtime.ClientAuthInfoWriter) (*CreateClientOK, error)

	DeleteClient(params *DeleteClientParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteClientAccepted, error)

	GetClient(params *GetClientParams, authInfo runtime.ClientAuthInfoWriter) (*GetClientOK, error)

	GetClients(params *GetClientsParams, authInfo runtime.ClientAuthInfoWriter) (*GetClientsOK, error)

	GetGdprUrls(params *GetGdprUrlsParams) (*GetGdprUrlsOK, error)

	RequestPasswordReset(params *RequestPasswordResetParams) error

	UpdateClient(params *UpdateClientParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateClientOK, error)

	Userinfo(params *UserinfoParams) (*UserinfoOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  ChangePassword changes the password
*/
func (a *Client) ChangePassword(params *ChangePasswordParams) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewChangePasswordParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "changePassword",
		Method:             "POST",
		PathPattern:        "/passwordreset/{nonce}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ChangePasswordReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return err
	}
	return nil
}

/*
  CreateClient creates a new o auth 2 0 client
*/
func (a *Client) CreateClient(params *CreateClientParams, authInfo runtime.ClientAuthInfoWriter) (*CreateClientOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateClientParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createClient",
		Method:             "POST",
		PathPattern:        "/authorization/oauth2/clients",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateClientReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreateClientOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for createClient: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DeleteClient deletes an existing o auth 2 0 client
*/
func (a *Client) DeleteClient(params *DeleteClientParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteClientAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteClientParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteClient",
		Method:             "DELETE",
		PathPattern:        "/authorization/oauth2/clients/{clientId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteClientReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteClientAccepted)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteClient: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetClient gets a specific o auth 2 0 client
*/
func (a *Client) GetClient(params *GetClientParams, authInfo runtime.ClientAuthInfoWriter) (*GetClientOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetClientParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getClient",
		Method:             "GET",
		PathPattern:        "/authorization/oauth2/clients/{clientId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetClientReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetClientOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getClient: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetClients lists all o auth 2 0 clients
*/
func (a *Client) GetClients(params *GetClientsParams, authInfo runtime.ClientAuthInfoWriter) (*GetClientsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetClientsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getClients",
		Method:             "GET",
		PathPattern:        "/authorization/oauth2/clients",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetClientsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetClientsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getClients: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetGdprUrls gets g d p r urls for specific o auth 2 0 client
*/
func (a *Client) GetGdprUrls(params *GetGdprUrlsParams) (*GetGdprUrlsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetGdprUrlsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getGdprUrls",
		Method:             "GET",
		PathPattern:        "/authorization/oauth2/clients/{clientId}/gdpr",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetGdprUrlsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetGdprUrlsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getGdprUrls: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  RequestPasswordReset requests a password change
*/
func (a *Client) RequestPasswordReset(params *RequestPasswordResetParams) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRequestPasswordResetParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "requestPasswordReset",
		Method:             "POST",
		PathPattern:        "/passwordreset",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &RequestPasswordResetReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return err
	}
	return nil
}

/*
  UpdateClient updates an existing o auth 2 0 client
*/
func (a *Client) UpdateClient(params *UpdateClientParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateClientOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateClientParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateClient",
		Method:             "PUT",
		PathPattern:        "/authorization/oauth2/clients/{clientId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateClientReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpdateClientOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateClient: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  Userinfo gets information about the logged in user
*/
func (a *Client) Userinfo(params *UserinfoParams) (*UserinfoOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUserinfoParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "userinfo",
		Method:             "GET",
		PathPattern:        "/authorization/userinfo",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UserinfoReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UserinfoOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for userinfo: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
