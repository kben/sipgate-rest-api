// Code generated by go-swagger; DO NOT EDIT.

package history

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewGetHistoryCSVParams creates a new GetHistoryCSVParams object
// with the default values initialized.
func NewGetHistoryCSVParams() *GetHistoryCSVParams {
	var (
		archivedDefault = bool(false)
		limitDefault    = int32(10)
		offsetDefault   = int32(0)
	)
	return &GetHistoryCSVParams{
		Archived: &archivedDefault,
		Limit:    &limitDefault,
		Offset:   &offsetDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetHistoryCSVParamsWithTimeout creates a new GetHistoryCSVParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetHistoryCSVParamsWithTimeout(timeout time.Duration) *GetHistoryCSVParams {
	var (
		archivedDefault = bool(false)
		limitDefault    = int32(10)
		offsetDefault   = int32(0)
	)
	return &GetHistoryCSVParams{
		Archived: &archivedDefault,
		Limit:    &limitDefault,
		Offset:   &offsetDefault,

		timeout: timeout,
	}
}

// NewGetHistoryCSVParamsWithContext creates a new GetHistoryCSVParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetHistoryCSVParamsWithContext(ctx context.Context) *GetHistoryCSVParams {
	var (
		archivedDefault = bool(false)
		limitDefault    = int32(10)
		offsetDefault   = int32(0)
	)
	return &GetHistoryCSVParams{
		Archived: &archivedDefault,
		Limit:    &limitDefault,
		Offset:   &offsetDefault,

		Context: ctx,
	}
}

// NewGetHistoryCSVParamsWithHTTPClient creates a new GetHistoryCSVParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetHistoryCSVParamsWithHTTPClient(client *http.Client) *GetHistoryCSVParams {
	var (
		archivedDefault = bool(false)
		limitDefault    = int32(10)
		offsetDefault   = int32(0)
	)
	return &GetHistoryCSVParams{
		Archived:   &archivedDefault,
		Limit:      &limitDefault,
		Offset:     &offsetDefault,
		HTTPClient: client,
	}
}

/*GetHistoryCSVParams contains all the parameters to send to the API endpoint
for the get history c s v operation typically these are written to a http.Request
*/
type GetHistoryCSVParams struct {

	/*Archived
	  Only show archived events

	*/
	Archived *bool
	/*ConnectionIds
	  List of extensions

	*/
	ConnectionIds []string
	/*Directions
	  Filter for incoming/outgoing/missed

	*/
	Directions []string
	/*From
	  Filter 'from' date time in ISO8601 format

	*/
	From *string
	/*Limit
	  Limit result rows

	*/
	Limit *int32
	/*Offset
	  The offset used for pagination

	*/
	Offset *int32
	/*Starred
	  Filter for starred/unstarred

	*/
	Starred []string
	/*To
	  Filter 'to' date time in ISO8601 format

	*/
	To *string
	/*Types
	  Filter by type

	*/
	Types []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get history c s v params
func (o *GetHistoryCSVParams) WithTimeout(timeout time.Duration) *GetHistoryCSVParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get history c s v params
func (o *GetHistoryCSVParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get history c s v params
func (o *GetHistoryCSVParams) WithContext(ctx context.Context) *GetHistoryCSVParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get history c s v params
func (o *GetHistoryCSVParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get history c s v params
func (o *GetHistoryCSVParams) WithHTTPClient(client *http.Client) *GetHistoryCSVParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get history c s v params
func (o *GetHistoryCSVParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithArchived adds the archived to the get history c s v params
func (o *GetHistoryCSVParams) WithArchived(archived *bool) *GetHistoryCSVParams {
	o.SetArchived(archived)
	return o
}

// SetArchived adds the archived to the get history c s v params
func (o *GetHistoryCSVParams) SetArchived(archived *bool) {
	o.Archived = archived
}

// WithConnectionIds adds the connectionIds to the get history c s v params
func (o *GetHistoryCSVParams) WithConnectionIds(connectionIds []string) *GetHistoryCSVParams {
	o.SetConnectionIds(connectionIds)
	return o
}

// SetConnectionIds adds the connectionIds to the get history c s v params
func (o *GetHistoryCSVParams) SetConnectionIds(connectionIds []string) {
	o.ConnectionIds = connectionIds
}

// WithDirections adds the directions to the get history c s v params
func (o *GetHistoryCSVParams) WithDirections(directions []string) *GetHistoryCSVParams {
	o.SetDirections(directions)
	return o
}

// SetDirections adds the directions to the get history c s v params
func (o *GetHistoryCSVParams) SetDirections(directions []string) {
	o.Directions = directions
}

// WithFrom adds the from to the get history c s v params
func (o *GetHistoryCSVParams) WithFrom(from *string) *GetHistoryCSVParams {
	o.SetFrom(from)
	return o
}

// SetFrom adds the from to the get history c s v params
func (o *GetHistoryCSVParams) SetFrom(from *string) {
	o.From = from
}

// WithLimit adds the limit to the get history c s v params
func (o *GetHistoryCSVParams) WithLimit(limit *int32) *GetHistoryCSVParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the get history c s v params
func (o *GetHistoryCSVParams) SetLimit(limit *int32) {
	o.Limit = limit
}

// WithOffset adds the offset to the get history c s v params
func (o *GetHistoryCSVParams) WithOffset(offset *int32) *GetHistoryCSVParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the get history c s v params
func (o *GetHistoryCSVParams) SetOffset(offset *int32) {
	o.Offset = offset
}

// WithStarred adds the starred to the get history c s v params
func (o *GetHistoryCSVParams) WithStarred(starred []string) *GetHistoryCSVParams {
	o.SetStarred(starred)
	return o
}

// SetStarred adds the starred to the get history c s v params
func (o *GetHistoryCSVParams) SetStarred(starred []string) {
	o.Starred = starred
}

// WithTo adds the to to the get history c s v params
func (o *GetHistoryCSVParams) WithTo(to *string) *GetHistoryCSVParams {
	o.SetTo(to)
	return o
}

// SetTo adds the to to the get history c s v params
func (o *GetHistoryCSVParams) SetTo(to *string) {
	o.To = to
}

// WithTypes adds the types to the get history c s v params
func (o *GetHistoryCSVParams) WithTypes(types []string) *GetHistoryCSVParams {
	o.SetTypes(types)
	return o
}

// SetTypes adds the types to the get history c s v params
func (o *GetHistoryCSVParams) SetTypes(types []string) {
	o.Types = types
}

// WriteToRequest writes these params to a swagger request
func (o *GetHistoryCSVParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Archived != nil {

		// query param archived
		var qrArchived bool
		if o.Archived != nil {
			qrArchived = *o.Archived
		}
		qArchived := swag.FormatBool(qrArchived)
		if qArchived != "" {
			if err := r.SetQueryParam("archived", qArchived); err != nil {
				return err
			}
		}

	}

	valuesConnectionIds := o.ConnectionIds

	joinedConnectionIds := swag.JoinByFormat(valuesConnectionIds, "multi")
	// query array param connectionIds
	if err := r.SetQueryParam("connectionIds", joinedConnectionIds...); err != nil {
		return err
	}

	valuesDirections := o.Directions

	joinedDirections := swag.JoinByFormat(valuesDirections, "multi")
	// query array param directions
	if err := r.SetQueryParam("directions", joinedDirections...); err != nil {
		return err
	}

	if o.From != nil {

		// query param from
		var qrFrom string
		if o.From != nil {
			qrFrom = *o.From
		}
		qFrom := qrFrom
		if qFrom != "" {
			if err := r.SetQueryParam("from", qFrom); err != nil {
				return err
			}
		}

	}

	if o.Limit != nil {

		// query param limit
		var qrLimit int32
		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt32(qrLimit)
		if qLimit != "" {
			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}

	}

	if o.Offset != nil {

		// query param offset
		var qrOffset int32
		if o.Offset != nil {
			qrOffset = *o.Offset
		}
		qOffset := swag.FormatInt32(qrOffset)
		if qOffset != "" {
			if err := r.SetQueryParam("offset", qOffset); err != nil {
				return err
			}
		}

	}

	valuesStarred := o.Starred

	joinedStarred := swag.JoinByFormat(valuesStarred, "multi")
	// query array param starred
	if err := r.SetQueryParam("starred", joinedStarred...); err != nil {
		return err
	}

	if o.To != nil {

		// query param to
		var qrTo string
		if o.To != nil {
			qrTo = *o.To
		}
		qTo := qrTo
		if qTo != "" {
			if err := r.SetQueryParam("to", qTo); err != nil {
				return err
			}
		}

	}

	valuesTypes := o.Types

	joinedTypes := swag.JoinByFormat(valuesTypes, "multi")
	// query array param types
	if err := r.SetQueryParam("types", joinedTypes...); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
