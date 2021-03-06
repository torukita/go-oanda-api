// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// TradeState The current state of the Trade.
// swagger:model TradeState
type TradeState string

const (
	// TradeStateOPEN captures enum value "OPEN"
//	TradeStateOPEN TradeState = "OPEN"
	// TradeStateCLOSED captures enum value "CLOSED"
//	TradeStateCLOSED TradeState = "CLOSED"
	// TradeStateCLOSEWHENTRADEABLE captures enum value "CLOSE_WHEN_TRADEABLE"
//	TradeStateCLOSEWHENTRADEABLE TradeState = "CLOSE_WHEN_TRADEABLE"
)

// for schema
var tradeStateEnum []interface{}

func init() {
	var res []TradeState
	if err := json.Unmarshal([]byte(`["OPEN","CLOSED","CLOSE_WHEN_TRADEABLE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		tradeStateEnum = append(tradeStateEnum, v)
	}
}

func (m TradeState) validateTradeStateEnum(path, location string, value TradeState) error {
	if err := validate.Enum(path, location, value, tradeStateEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this trade state
func (m TradeState) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateTradeStateEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
