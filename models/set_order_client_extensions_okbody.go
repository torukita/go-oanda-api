// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// SetOrderClientExtensionsOKBody set order client extensions o k body
// swagger:model setOrderClientExtensionsOKBody
type SetOrderClientExtensionsOKBody struct {

	// The ID of the most recent Transaction created for the Account
	LastTransactionID string `json:"lastTransactionID,omitempty"`

	// order client extensions modify transaction
	OrderClientExtensionsModifyTransaction *OrderClientExtensionsModifyTransaction `json:"orderClientExtensionsModifyTransaction,omitempty"`

	// The IDs of all Transactions that were created while satisfying the request.
	RelatedTransactionIds []string `json:"relatedTransactionIDs"`
}

// Validate validates this set order client extensions o k body
func (m *SetOrderClientExtensionsOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOrderClientExtensionsModifyTransaction(formats); err != nil {
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

func (m *SetOrderClientExtensionsOKBody) validateOrderClientExtensionsModifyTransaction(formats strfmt.Registry) error {

	if swag.IsZero(m.OrderClientExtensionsModifyTransaction) { // not required
		return nil
	}

	if m.OrderClientExtensionsModifyTransaction != nil {

		if err := m.OrderClientExtensionsModifyTransaction.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("orderClientExtensionsModifyTransaction")
			}
			return err
		}
	}

	return nil
}

func (m *SetOrderClientExtensionsOKBody) validateRelatedTransactionIds(formats strfmt.Registry) error {

	if swag.IsZero(m.RelatedTransactionIds) { // not required
		return nil
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SetOrderClientExtensionsOKBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SetOrderClientExtensionsOKBody) UnmarshalBinary(b []byte) error {
	var res SetOrderClientExtensionsOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
