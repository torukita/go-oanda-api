// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// CloseTradeOKBody close trade o k body
// swagger:model closeTradeOKBody
type CloseTradeOKBody struct {

	// The ID of the most recent Transaction created for the Account
	LastTransactionID string `json:"lastTransactionID,omitempty"`

	// order cancel transaction
	OrderCancelTransaction *OrderCancelTransaction `json:"orderCancelTransaction,omitempty"`

	// order create transaction
	OrderCreateTransaction *MarketOrderTransaction `json:"orderCreateTransaction,omitempty"`

	// order fill transaction
	OrderFillTransaction *OrderFillTransaction `json:"orderFillTransaction,omitempty"`

	// The IDs of all Transactions that were created while satisfying the request.
	RelatedTransactionIds []string `json:"relatedTransactionIDs"`
}

// Validate validates this close trade o k body
func (m *CloseTradeOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOrderCancelTransaction(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateOrderCreateTransaction(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateOrderFillTransaction(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateRelatedTransactionIds(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CloseTradeOKBody) validateOrderCancelTransaction(formats strfmt.Registry) error {

	if swag.IsZero(m.OrderCancelTransaction) { // not required
		return nil
	}

	if m.OrderCancelTransaction != nil {

		if err := m.OrderCancelTransaction.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("orderCancelTransaction")
			}
			return err
		}
	}

	return nil
}

func (m *CloseTradeOKBody) validateOrderCreateTransaction(formats strfmt.Registry) error {

	if swag.IsZero(m.OrderCreateTransaction) { // not required
		return nil
	}

	if m.OrderCreateTransaction != nil {

		if err := m.OrderCreateTransaction.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("orderCreateTransaction")
			}
			return err
		}
	}

	return nil
}

func (m *CloseTradeOKBody) validateOrderFillTransaction(formats strfmt.Registry) error {

	if swag.IsZero(m.OrderFillTransaction) { // not required
		return nil
	}

	if m.OrderFillTransaction != nil {

		if err := m.OrderFillTransaction.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("orderFillTransaction")
			}
			return err
		}
	}

	return nil
}

func (m *CloseTradeOKBody) validateRelatedTransactionIds(formats strfmt.Registry) error {

	if swag.IsZero(m.RelatedTransactionIds) { // not required
		return nil
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CloseTradeOKBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CloseTradeOKBody) UnmarshalBinary(b []byte) error {
	var res CloseTradeOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}