// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// StreamTransactionsOKBody The response body for the Transaction Stream uses chunked transfer encoding.  Each chunk contains Transaction and/or TransactionHeartbeat objects encoded as JSON.  Each JSON object is serialized into a single line of text, and multiple objects found in the same chunk are separated by newlines.
// TransactionHeartbeats are sent every 5 seconds.
// swagger:model streamTransactionsOKBody
type StreamTransactionsOKBody struct {

	// heartbeat
	Heartbeat *TransactionHeartbeat `json:"heartbeat,omitempty"`

	// transaction
	Transaction *Transaction `json:"transaction,omitempty"`
}

// Validate validates this stream transactions o k body
func (m *StreamTransactionsOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHeartbeat(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateTransaction(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *StreamTransactionsOKBody) validateHeartbeat(formats strfmt.Registry) error {

	if swag.IsZero(m.Heartbeat) { // not required
		return nil
	}

	if m.Heartbeat != nil {

		if err := m.Heartbeat.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("heartbeat")
			}
			return err
		}
	}

	return nil
}

func (m *StreamTransactionsOKBody) validateTransaction(formats strfmt.Registry) error {

	if swag.IsZero(m.Transaction) { // not required
		return nil
	}

	if m.Transaction != nil {

		if err := m.Transaction.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("transaction")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *StreamTransactionsOKBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StreamTransactionsOKBody) UnmarshalBinary(b []byte) error {
	var res StreamTransactionsOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}