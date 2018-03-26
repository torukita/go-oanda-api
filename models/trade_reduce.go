// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// TradeReduce A TradeReduce object represents a Trade for an instrument that was reduced (either partially or fully) in an Account. It is found embedded in Transactions that affect the position of an instrument in the account, specifically the OrderFill Transaction.
// swagger:model TradeReduce
type TradeReduce struct {

	// The financing paid/collected when reducing the Trade
	Financing string `json:"financing,omitempty"`

	// The PL realized when reducing the Trade
	RealizedPL string `json:"realizedPL,omitempty"`

	// The ID of the Trade that was reduced or closed
	TradeID string `json:"tradeID,omitempty"`

	// The number of units that the Trade was reduced by
	Units string `json:"units,omitempty"`
}

// Validate validates this trade reduce
func (m *TradeReduce) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *TradeReduce) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TradeReduce) UnmarshalBinary(b []byte) error {
	var res TradeReduce
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}