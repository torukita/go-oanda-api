// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// GetTradeOKBody get trade o k body
// swagger:model getTradeOKBody
type GetTradeOKBody struct {

	// The ID of the most recent Transaction created for the Account
	LastTransactionID string `json:"lastTransactionID,omitempty"`

	// trade
	Trade *Trade `json:"trade,omitempty"`
}

// Validate validates this get trade o k body
func (m *GetTradeOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTrade(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetTradeOKBody) validateTrade(formats strfmt.Registry) error {

	if swag.IsZero(m.Trade) { // not required
		return nil
	}

	if m.Trade != nil {

		if err := m.Trade.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("trade")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GetTradeOKBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetTradeOKBody) UnmarshalBinary(b []byte) error {
	var res GetTradeOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
