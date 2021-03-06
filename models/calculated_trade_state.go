// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// CalculatedTradeState The dynamic (calculated) state of an open Trade
// swagger:model CalculatedTradeState
type CalculatedTradeState struct {

	// The Trade's ID.
	ID string `json:"id,omitempty"`

	// The Trade's unrealized profit/loss.
	UnrealizedPL string `json:"unrealizedPL,omitempty"`
}

// Validate validates this calculated trade state
func (m *CalculatedTradeState) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *CalculatedTradeState) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CalculatedTradeState) UnmarshalBinary(b []byte) error {
	var res CalculatedTradeState
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
