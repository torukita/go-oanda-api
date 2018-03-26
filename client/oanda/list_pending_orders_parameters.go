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

	strfmt "github.com/go-openapi/strfmt"
)

// NewListPendingOrdersParams creates a new ListPendingOrdersParams object
// with the default values initialized.
func NewListPendingOrdersParams() *ListPendingOrdersParams {
	var ()
	return &ListPendingOrdersParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewListPendingOrdersParamsWithTimeout creates a new ListPendingOrdersParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListPendingOrdersParamsWithTimeout(timeout time.Duration) *ListPendingOrdersParams {
	var ()
	return &ListPendingOrdersParams{

		timeout: timeout,
	}
}

// NewListPendingOrdersParamsWithContext creates a new ListPendingOrdersParams object
// with the default values initialized, and the ability to set a context for a request
func NewListPendingOrdersParamsWithContext(ctx context.Context) *ListPendingOrdersParams {
	var ()
	return &ListPendingOrdersParams{

		Context: ctx,
	}
}

// NewListPendingOrdersParamsWithHTTPClient creates a new ListPendingOrdersParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListPendingOrdersParamsWithHTTPClient(client *http.Client) *ListPendingOrdersParams {
	var ()
	return &ListPendingOrdersParams{
		HTTPClient: client,
	}
}

/*ListPendingOrdersParams contains all the parameters to send to the API endpoint
for the list pending orders operation typically these are written to a http.Request
*/
type ListPendingOrdersParams struct {

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

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list pending orders params
func (o *ListPendingOrdersParams) WithTimeout(timeout time.Duration) *ListPendingOrdersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list pending orders params
func (o *ListPendingOrdersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list pending orders params
func (o *ListPendingOrdersParams) WithContext(ctx context.Context) *ListPendingOrdersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list pending orders params
func (o *ListPendingOrdersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list pending orders params
func (o *ListPendingOrdersParams) WithHTTPClient(client *http.Client) *ListPendingOrdersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list pending orders params
func (o *ListPendingOrdersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAcceptDatetimeFormat adds the acceptDatetimeFormat to the list pending orders params
func (o *ListPendingOrdersParams) WithAcceptDatetimeFormat(acceptDatetimeFormat *string) *ListPendingOrdersParams {
	o.SetAcceptDatetimeFormat(acceptDatetimeFormat)
	return o
}

// SetAcceptDatetimeFormat adds the acceptDatetimeFormat to the list pending orders params
func (o *ListPendingOrdersParams) SetAcceptDatetimeFormat(acceptDatetimeFormat *string) {
	o.AcceptDatetimeFormat = acceptDatetimeFormat
}

// WithAuthorization adds the authorization to the list pending orders params
func (o *ListPendingOrdersParams) WithAuthorization(authorization string) *ListPendingOrdersParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the list pending orders params
func (o *ListPendingOrdersParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithAccountID adds the accountID to the list pending orders params
func (o *ListPendingOrdersParams) WithAccountID(accountID string) *ListPendingOrdersParams {
	o.SetAccountID(accountID)
	return o
}

// SetAccountID adds the accountId to the list pending orders params
func (o *ListPendingOrdersParams) SetAccountID(accountID string) {
	o.AccountID = accountID
}

// WriteToRequest writes these params to a swagger request
func (o *ListPendingOrdersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
