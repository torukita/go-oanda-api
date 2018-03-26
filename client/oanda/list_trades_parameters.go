// Code generated by go-swagger; DO NOT EDIT.

package oanda

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewListTradesParams creates a new ListTradesParams object
// with the default values initialized.
func NewListTradesParams() *ListTradesParams {
	var ()
	return &ListTradesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewListTradesParamsWithTimeout creates a new ListTradesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListTradesParamsWithTimeout(timeout time.Duration) *ListTradesParams {
	var ()
	return &ListTradesParams{

		timeout: timeout,
	}
}

// NewListTradesParamsWithContext creates a new ListTradesParams object
// with the default values initialized, and the ability to set a context for a request
func NewListTradesParamsWithContext(ctx context.Context) *ListTradesParams {
	var ()
	return &ListTradesParams{

		Context: ctx,
	}
}

// NewListTradesParamsWithHTTPClient creates a new ListTradesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListTradesParamsWithHTTPClient(client *http.Client) *ListTradesParams {
	var ()
	return &ListTradesParams{
		HTTPClient: client,
	}
}

/*ListTradesParams contains all the parameters to send to the API endpoint
for the list trades operation typically these are written to a http.Request
*/
type ListTradesParams struct {

	/*AcceptDatetimeFormat
	  Format of DateTime fields in the request and response.

	*/
	AcceptDatetimeFormat *string
	/*Authorization
	  The authorization bearer token previously obtained by the client

	*/
	Authorization string
	/*AccountID
	  Account Identifier

	*/
	AccountID string
	/*BeforeID
	  The maximum Trade ID to return. If not provided the most recent Trades in the Account are returned.

	*/
	BeforeID *string
	/*Count
	  The maximum number of Trades to return.

	*/
	Count *int64
	/*Ids
	  List of Trade IDs to retrieve.

	*/
	Ids []string
	/*Instrument
	  The instrument to filter the requested Trades by.

	*/
	Instrument *string
	/*State
	  The state to filter the requested Trades by.

	*/
	State *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list trades params
func (o *ListTradesParams) WithTimeout(timeout time.Duration) *ListTradesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list trades params
func (o *ListTradesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list trades params
func (o *ListTradesParams) WithContext(ctx context.Context) *ListTradesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list trades params
func (o *ListTradesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list trades params
func (o *ListTradesParams) WithHTTPClient(client *http.Client) *ListTradesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list trades params
func (o *ListTradesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAcceptDatetimeFormat adds the acceptDatetimeFormat to the list trades params
func (o *ListTradesParams) WithAcceptDatetimeFormat(acceptDatetimeFormat *string) *ListTradesParams {
	o.SetAcceptDatetimeFormat(acceptDatetimeFormat)
	return o
}

// SetAcceptDatetimeFormat adds the acceptDatetimeFormat to the list trades params
func (o *ListTradesParams) SetAcceptDatetimeFormat(acceptDatetimeFormat *string) {
	o.AcceptDatetimeFormat = acceptDatetimeFormat
}

// WithAuthorization adds the authorization to the list trades params
func (o *ListTradesParams) WithAuthorization(authorization string) *ListTradesParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the list trades params
func (o *ListTradesParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithAccountID adds the accountID to the list trades params
func (o *ListTradesParams) WithAccountID(accountID string) *ListTradesParams {
	o.SetAccountID(accountID)
	return o
}

// SetAccountID adds the accountId to the list trades params
func (o *ListTradesParams) SetAccountID(accountID string) {
	o.AccountID = accountID
}

// WithBeforeID adds the beforeID to the list trades params
func (o *ListTradesParams) WithBeforeID(beforeID *string) *ListTradesParams {
	o.SetBeforeID(beforeID)
	return o
}

// SetBeforeID adds the beforeId to the list trades params
func (o *ListTradesParams) SetBeforeID(beforeID *string) {
	o.BeforeID = beforeID
}

// WithCount adds the count to the list trades params
func (o *ListTradesParams) WithCount(count *int64) *ListTradesParams {
	o.SetCount(count)
	return o
}

// SetCount adds the count to the list trades params
func (o *ListTradesParams) SetCount(count *int64) {
	o.Count = count
}

// WithIds adds the ids to the list trades params
func (o *ListTradesParams) WithIds(ids []string) *ListTradesParams {
	o.SetIds(ids)
	return o
}

// SetIds adds the ids to the list trades params
func (o *ListTradesParams) SetIds(ids []string) {
	o.Ids = ids
}

// WithInstrument adds the instrument to the list trades params
func (o *ListTradesParams) WithInstrument(instrument *string) *ListTradesParams {
	o.SetInstrument(instrument)
	return o
}

// SetInstrument adds the instrument to the list trades params
func (o *ListTradesParams) SetInstrument(instrument *string) {
	o.Instrument = instrument
}

// WithState adds the state to the list trades params
func (o *ListTradesParams) WithState(state *string) *ListTradesParams {
	o.SetState(state)
	return o
}

// SetState adds the state to the list trades params
func (o *ListTradesParams) SetState(state *string) {
	o.State = state
}

// WriteToRequest writes these params to a swagger request
func (o *ListTradesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.AcceptDatetimeFormat != nil {

		// header param Accept-Datetime-Format
		if err := r.SetHeaderParam("Accept-Datetime-Format", *o.AcceptDatetimeFormat); err != nil {
			return err
		}

	}

	// header param Authorization
	if err := r.SetHeaderParam("Authorization", o.Authorization); err != nil {
		return err
	}

	// path param accountID
	if err := r.SetPathParam("accountID", o.AccountID); err != nil {
		return err
	}

	if o.BeforeID != nil {

		// query param beforeID
		var qrBeforeID string
		if o.BeforeID != nil {
			qrBeforeID = *o.BeforeID
		}
		qBeforeID := qrBeforeID
		if qBeforeID != "" {
			if err := r.SetQueryParam("beforeID", qBeforeID); err != nil {
				return err
			}
		}

	}

	if o.Count != nil {

		// query param count
		var qrCount int64
		if o.Count != nil {
			qrCount = *o.Count
		}
		qCount := swag.FormatInt64(qrCount)
		if qCount != "" {
			if err := r.SetQueryParam("count", qCount); err != nil {
				return err
			}
		}

	}

	valuesIds := o.Ids

	joinedIds := swag.JoinByFormat(valuesIds, "csv")
	// query array param ids
	if err := r.SetQueryParam("ids", joinedIds...); err != nil {
		return err
	}

	if o.Instrument != nil {

		// query param instrument
		var qrInstrument string
		if o.Instrument != nil {
			qrInstrument = *o.Instrument
		}
		qInstrument := qrInstrument
		if qInstrument != "" {
			if err := r.SetQueryParam("instrument", qInstrument); err != nil {
				return err
			}
		}

	}

	if o.State != nil {

		// query param state
		var qrState string
		if o.State != nil {
			qrState = *o.State
		}
		qState := qrState
		if qState != "" {
			if err := r.SetQueryParam("state", qState); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}