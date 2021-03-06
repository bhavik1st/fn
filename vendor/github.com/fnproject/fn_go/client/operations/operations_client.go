// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new operations API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for operations API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
DeleteAppsAppCallsCallLog deletes call log entry

Delete call log entry
*/
func (a *Client) DeleteAppsAppCallsCallLog(params *DeleteAppsAppCallsCallLogParams) (*DeleteAppsAppCallsCallLogAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteAppsAppCallsCallLogParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteAppsAppCallsCallLog",
		Method:             "DELETE",
		PathPattern:        "/apps/{app}/calls/{call}/log",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteAppsAppCallsCallLogReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteAppsAppCallsCallLogAccepted), nil

}

/*
GetAppsAppCallsCallLog gets call logs

Get call logs
*/
func (a *Client) GetAppsAppCallsCallLog(params *GetAppsAppCallsCallLogParams) (*GetAppsAppCallsCallLogOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAppsAppCallsCallLogParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetAppsAppCallsCallLog",
		Method:             "GET",
		PathPattern:        "/apps/{app}/calls/{call}/log",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetAppsAppCallsCallLogReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetAppsAppCallsCallLogOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
