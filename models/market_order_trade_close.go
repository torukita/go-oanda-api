// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// MarketOrderTradeClose A MarketOrderTradeClose specifies the extensions to a Market Order that has been created specifically to close a Trade.
// swagger:model MarketOrderTradeClose
type MarketOrderTradeClose struct {

	// The client ID of the Trade requested to be closed
	ClientTradeID string `json:"clientTradeID,omitempty"`

	// The ID of the Trade requested to be closed
	TradeID string `json:"tradeID,omitempty"`

	// Indication of how much of the Trade to close. Either "ALL", or a DecimalNumber reflection a partial close of the Trade.
	Units string `json:"units,omitempty"`
}

// Validate validates this market order trade close
func (m *MarketOrderTradeClose) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *MarketOrderTradeClose) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MarketOrderTradeClose) UnmarshalBinary(b []byte) error {
	var res MarketOrderTradeClose
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}