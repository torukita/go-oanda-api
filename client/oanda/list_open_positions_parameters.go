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

// NewListOpenPositionsParams creates a new ListOpenPositionsParams object
// with the default values initialized.
func NewListOpenPositionsParams() *ListOpenPositionsParams {
	var ()
	return &ListOpenPositionsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewListOpenPositionsParamsWithTimeout creates a new ListOpenPositionsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListOpenPositionsParamsWithTimeout(timeout time.Duration) *ListOpenPositionsParams {
	var ()
	return &ListOpenPositionsParams{

		timeout: timeout,
	}
}

// NewListOpenPositionsParamsWithContext creates a new ListOpenPositionsParams object
// with the default values initialized, and the ability to set a context for a request
func NewListOpenPositionsParamsWithContext(ctx context.Context) *ListOpenPositionsParams {
	var ()
	return &ListOpenPositionsParams{

		Context: ctx,
	}
}

// NewListOpenPositionsParamsWithHTTPClient creates a new ListOpenPositionsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListOpenPositionsParamsWithHTTPClient(client *http.Client) *ListOpenPositionsParams {
	var ()
	return &ListOpenPositionsParams{
		HTTPClient: client,
	}
}

/*ListOpenPositionsParams contains all the parameters to send to the API endpoint
for the list open positions operation typically these are written to a http.Request
*/
type ListOpenPositionsParams struct {

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

// WithTimeout adds the timeout to the list open positions params
func (o *ListOpenPositionsParams) WithTimeout(timeout time.Duration) *ListOpenPositionsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list open positions params
func (o *ListOpenPositionsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list open positions params
func (o *ListOpenPositionsParams) WithContext(ctx context.Context) *ListOpenPositionsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list open positions params
func (o *ListOpenPositionsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list open positions params
func (o *ListOpenPositionsParams) WithHTTPClient(client *http.Client) *ListOpenPositionsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list open positions params
func (o *ListOpenPositionsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the list open positions params
func (o *ListOpenPositionsParams) WithAuthorization(authorization string) *ListOpenPositionsParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the list open positions params
func (o *ListOpenPositionsParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithAccountID adds the accountID to the list open positions params
func (o *ListOpenPositionsParams) WithAccountID(accountID string) *ListOpenPositionsParams {
	o.SetAccountID(accountID)
	return o
}

// SetAccountID adds the accountId to the list open positions params
func (o *ListOpenPositionsParams) SetAccountID(accountID string) {
	o.AccountID = accountID
}

// WriteToRequest writes these params to a swagger request
func (o *ListOpenPositionsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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