// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// ClosePositionNotFoundBody close position not found body
// swagger:model closePositionNotFoundBody
type ClosePositionNotFoundBody struct {

	// The code of the error that has occurred. This field may not be returned for some errors.
	ErrorCode string `json:"errorCode,omitempty"`

	// The human-readable description of the error that has occurred.
	ErrorMessage string `json:"errorMessage,omitempty"`

	// The ID of the most recent Transaction created for the Account. Only present if the Account exists.
	LastTransactionID string `json:"lastTransactionID,omitempty"`

	// long order reject transaction
	LongOrderRejectTransaction *MarketOrderRejectTransaction `json:"longOrderRejectTransaction,omitempty"`

	// The IDs of all Transactions that were created while satisfying the request. Only present if the Account exists.
	RelatedTransactionIds []string `json:"relatedTransactionIDs"`

	// short order reject transaction
	ShortOrderRejectTransaction *MarketOrderRejectTransaction `json:"shortOrderRejectTransaction,omitempty"`
}

// Validate validates this close position not found body
func (m *ClosePositionNotFoundBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLongOrderRejectTransaction(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateRelatedTransactionIds(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateShortOrderRejectTransaction(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClosePositionNotFoundBody) validateLongOrderRejectTransaction(formats strfmt.Registry) error {

	if swag.IsZero(m.LongOrderRejectTransaction) { // not required
		return nil
	}

	if m.LongOrderRejectTransaction != nil {

		if err := m.LongOrderRejectTransaction.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("longOrderRejectTransaction")
			}
			return err
		}
	}

	return nil
}

func (m *ClosePositionNotFoundBody) validateRelatedTransactionIds(formats strfmt.Registry) error {

	if swag.IsZero(m.RelatedTransactionIds) { // not required
		return nil
	}

	return nil
}

func (m *ClosePositionNotFoundBody) validateShortOrderRejectTransaction(formats strfmt.Registry) error {

	if swag.IsZero(m.ShortOrderRejectTransaction) { // not required
		return nil
	}

	if m.ShortOrderRejectTransaction != nil {

		if err := m.ShortOrderRejectTransaction.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("shortOrderRejectTransaction")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ClosePositionNotFoundBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClosePositionNotFoundBody) UnmarshalBinary(b []byte) error {
	var res ClosePositionNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}