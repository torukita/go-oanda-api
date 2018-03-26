// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// SetTradeDependentOrdersBadRequestBody set trade dependent orders bad request body
// swagger:model setTradeDependentOrdersBadRequestBody
type SetTradeDependentOrdersBadRequestBody struct {

	// The code of the error that has occurred. This field may not be returned for some errors.
	ErrorCode string `json:"errorCode,omitempty"`

	// The human-readable description of the error that has occurred.
	ErrorMessage string `json:"errorMessage,omitempty"`

	// The ID of the most recent Transaction created for the Account.
	LastTransactionID string `json:"lastTransactionID,omitempty"`

	// The IDs of all Transactions that were created while satisfying the request.
	RelatedTransactionIds []string `json:"relatedTransactionIDs"`

	// stop loss order cancel reject transaction
	StopLossOrderCancelRejectTransaction *OrderCancelRejectTransaction `json:"stopLossOrderCancelRejectTransaction,omitempty"`

	// stop loss order reject transaction
	StopLossOrderRejectTransaction *StopLossOrderRejectTransaction `json:"stopLossOrderRejectTransaction,omitempty"`

	// take profit order cancel reject transaction
	TakeProfitOrderCancelRejectTransaction *OrderCancelRejectTransaction `json:"takeProfitOrderCancelRejectTransaction,omitempty"`

	// take profit order reject transaction
	TakeProfitOrderRejectTransaction *TakeProfitOrderRejectTransaction `json:"takeProfitOrderRejectTransaction,omitempty"`

	// trailing stop loss order cancel reject transaction
	TrailingStopLossOrderCancelRejectTransaction *OrderCancelRejectTransaction `json:"trailingStopLossOrderCancelRejectTransaction,omitempty"`

	// trailing stop loss order reject transaction
	TrailingStopLossOrderRejectTransaction *TrailingStopLossOrderRejectTransaction `json:"trailingStopLossOrderRejectTransaction,omitempty"`
}

// Validate validates this set trade dependent orders bad request body
func (m *SetTradeDependentOrdersBadRequestBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRelatedTransactionIds(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateStopLossOrderCancelRejectTransaction(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateStopLossOrderRejectTransaction(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateTakeProfitOrderCancelRejectTransaction(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateTakeProfitOrderRejectTransaction(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateTrailingStopLossOrderCancelRejectTransaction(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateTrailingStopLossOrderRejectTransaction(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SetTradeDependentOrdersBadRequestBody) validateRelatedTransactionIds(formats strfmt.Registry) error {

	if swag.IsZero(m.RelatedTransactionIds) { // not required
		return nil
	}

	return nil
}

func (m *SetTradeDependentOrdersBadRequestBody) validateStopLossOrderCancelRejectTransaction(formats strfmt.Registry) error {

	if swag.IsZero(m.StopLossOrderCancelRejectTransaction) { // not required
		return nil
	}

	if m.StopLossOrderCancelRejectTransaction != nil {

		if err := m.StopLossOrderCancelRejectTransaction.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("stopLossOrderCancelRejectTransaction")
			}
			return err
		}
	}

	return nil
}

func (m *SetTradeDependentOrdersBadRequestBody) validateStopLossOrderRejectTransaction(formats strfmt.Registry) error {

	if swag.IsZero(m.StopLossOrderRejectTransaction) { // not required
		return nil
	}

	if m.StopLossOrderRejectTransaction != nil {

		if err := m.StopLossOrderRejectTransaction.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("stopLossOrderRejectTransaction")
			}
			return err
		}
	}

	return nil
}

func (m *SetTradeDependentOrdersBadRequestBody) validateTakeProfitOrderCancelRejectTransaction(formats strfmt.Registry) error {

	if swag.IsZero(m.TakeProfitOrderCancelRejectTransaction) { // not required
		return nil
	}

	if m.TakeProfitOrderCancelRejectTransaction != nil {

		if err := m.TakeProfitOrderCancelRejectTransaction.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("takeProfitOrderCancelRejectTransaction")
			}
			return err
		}
	}

	return nil
}

func (m *SetTradeDependentOrdersBadRequestBody) validateTakeProfitOrderRejectTransaction(formats strfmt.Registry) error {

	if swag.IsZero(m.TakeProfitOrderRejectTransaction) { // not required
		return nil
	}

	if m.TakeProfitOrderRejectTransaction != nil {

		if err := m.TakeProfitOrderRejectTransaction.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("takeProfitOrderRejectTransaction")
			}
			return err
		}
	}

	return nil
}

func (m *SetTradeDependentOrdersBadRequestBody) validateTrailingStopLossOrderCancelRejectTransaction(formats strfmt.Registry) error {

	if swag.IsZero(m.TrailingStopLossOrderCancelRejectTransaction) { // not required
		return nil
	}

	if m.TrailingStopLossOrderCancelRejectTransaction != nil {

		if err := m.TrailingStopLossOrderCancelRejectTransaction.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("trailingStopLossOrderCancelRejectTransaction")
			}
			return err
		}
	}

	return nil
}

func (m *SetTradeDependentOrdersBadRequestBody) validateTrailingStopLossOrderRejectTransaction(formats strfmt.Registry) error {

	if swag.IsZero(m.TrailingStopLossOrderRejectTransaction) { // not required
		return nil
	}

	if m.TrailingStopLossOrderRejectTransaction != nil {

		if err := m.TrailingStopLossOrderRejectTransaction.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("trailingStopLossOrderRejectTransaction")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SetTradeDependentOrdersBadRequestBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SetTradeDependentOrdersBadRequestBody) UnmarshalBinary(b []byte) error {
	var res SetTradeDependentOrdersBadRequestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}