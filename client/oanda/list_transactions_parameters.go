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

// NewListTransactionsParams creates a new ListTransactionsParams object
// with the default values initialized.
func NewListTransactionsParams() *ListTransactionsParams {
	var ()
	return &ListTransactionsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewListTransactionsParamsWithTimeout creates a new ListTransactionsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListTransactionsParamsWithTimeout(timeout time.Duration) *ListTransactionsParams {
	var ()
	return &ListTransactionsParams{

		timeout: timeout,
	}
}

// NewListTransactionsParamsWithContext creates a new ListTransactionsParams object
// with the default values initialized, and the ability to set a context for a request
func NewListTransactionsParamsWithContext(ctx context.Context) *ListTransactionsParams {
	var ()
	return &ListTransactionsParams{

		Context: ctx,
	}
}

// NewListTransactionsParamsWithHTTPClient creates a new ListTransactionsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListTransactionsParamsWithHTTPClient(client *http.Client) *ListTransactionsParams {
	var ()
	return &ListTransactionsParams{
		HTTPClient: client,
	}
}

/*ListTransactionsParams contains all the parameters to send to the API endpoint
for the list transactions operation typically these are written to a http.Request
*/
type ListTransactionsParams struct {

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
	/*From
	  The starting time (inclusive) of the time range for the Transactions being queried.

	*/
	From *string
	/*PageSize
	  The number of Transactions to include in each page of the results.

	*/
	PageSize *int64
	/*To
	  The ending time (inclusive) of the time range for the Transactions being queried.

	*/
	To *string
	/*Type
	  A filter for restricting the types of Transactions to retreive.

	*/
	Type []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list transactions params
func (o *ListTransactionsParams) WithTimeout(timeout time.Duration) *ListTransactionsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list transactions params
func (o *ListTransactionsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list transactions params
func (o *ListTransactionsParams) WithContext(ctx context.Context) *ListTransactionsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list transactions params
func (o *ListTransactionsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list transactions params
func (o *ListTransactionsParams) WithHTTPClient(client *http.Client) *ListTransactionsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list transactions params
func (o *ListTransactionsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAcceptDatetimeFormat adds the acceptDatetimeFormat to the list transactions params
func (o *ListTransactionsParams) WithAcceptDatetimeFormat(acceptDatetimeFormat *string) *ListTransactionsParams {
	o.SetAcceptDatetimeFormat(acceptDatetimeFormat)
	return o
}

// SetAcceptDatetimeFormat adds the acceptDatetimeFormat to the list transactions params
func (o *ListTransactionsParams) SetAcceptDatetimeFormat(acceptDatetimeFormat *string) {
	o.AcceptDatetimeFormat = acceptDatetimeFormat
}

// WithAuthorization adds the authorization to the list transactions params
func (o *ListTransactionsParams) WithAuthorization(authorization string) *ListTransactionsParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the list transactions params
func (o *ListTransactionsParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithAccountID adds the accountID to the list transactions params
func (o *ListTransactionsParams) WithAccountID(accountID string) *ListTransactionsParams {
	o.SetAccountID(accountID)
	return o
}

// SetAccountID adds the accountId to the list transactions params
func (o *ListTransactionsParams) SetAccountID(accountID string) {
	o.AccountID = accountID
}

// WithFrom adds the from to the list transactions params
func (o *ListTransactionsParams) WithFrom(from *string) *ListTransactionsParams {
	o.SetFrom(from)
	return o
}

// SetFrom adds the from to the list transactions params
func (o *ListTransactionsParams) SetFrom(from *string) {
	o.From = from
}

// WithPageSize adds the pageSize to the list transactions params
func (o *ListTransactionsParams) WithPageSize(pageSize *int64) *ListTransactionsParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the list transactions params
func (o *ListTransactionsParams) SetPageSize(pageSize *int64) {
	o.PageSize = pageSize
}

// WithTo adds the to to the list transactions params
func (o *ListTransactionsParams) WithTo(to *string) *ListTransactionsParams {
	o.SetTo(to)
	return o
}

// SetTo adds the to to the list transactions params
func (o *ListTransactionsParams) SetTo(to *string) {
	o.To = to
}

// WithType adds the typeVar to the list transactions params
func (o *ListTransactionsParams) WithType(typeVar []string) *ListTransactionsParams {
	o.SetType(typeVar)
	return o
}

// SetType adds the type to the list transactions params
func (o *ListTransactionsParams) SetType(typeVar []string) {
	o.Type = typeVar
}

// WriteToRequest writes these params to a swagger request
func (o *ListTransactionsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.PageSize != nil {

		// query param pageSize
		var qrPageSize int64
		if o.PageSize != nil {
			qrPageSize = *o.PageSize
		}
		qPageSize := swag.FormatInt64(qrPageSize)
		if qPageSize != "" {
			if err := r.SetQueryParam("pageSize", qPageSize); err != nil {
				return err
			}
		}

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

	valuesType := o.Type

	joinedType := swag.JoinByFormat(valuesType, "csv")
	// query array param type
	if err := r.SetQueryParam("type", joinedType...); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
