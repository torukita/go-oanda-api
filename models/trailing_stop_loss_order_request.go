// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// TrailingStopLossOrderRequest A TrailingStopLossOrderRequest specifies the parameters that may be set when creating a Trailing Stop Loss Order.
// swagger:model TrailingStopLossOrderRequest
type TrailingStopLossOrderRequest struct {

	// client extensions
	ClientExtensions *ClientExtensions `json:"clientExtensions,omitempty"`

	// The client ID of the Trade to be closed when the price threshold is breached.
	ClientTradeID string `json:"clientTradeID,omitempty"`

	// The price distance specified for the TrailingStopLoss Order.
	Distance string `json:"distance,omitempty"`

	// The date/time when the StopLoss Order will be cancelled if its timeInForce is "GTD".
	GtdTime string `json:"gtdTime,omitempty"`

	// The time-in-force requested for the TrailingStopLoss Order. Restricted to "GTC", "GFD" and "GTD" for TrailingStopLoss Orders.
	TimeInForce string `json:"timeInForce,omitempty"`

	// The ID of the Trade to close when the price threshold is breached.
	TradeID string `json:"tradeID,omitempty"`

	// Specification of what component of a price should be used for comparison when determining if the Order should be filled.
	TriggerCondition string `json:"triggerCondition,omitempty"`

	// The type of the Order to Create. Must be set to "TRAILING_STOP_LOSS" when creating a Trailng Stop Loss Order.
	Type string `json:"type,omitempty"`
}

// Validate validates this trailing stop loss order request
func (m *TrailingStopLossOrderRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateClientExtensions(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateTimeInForce(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateTriggerCondition(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TrailingStopLossOrderRequest) validateClientExtensions(formats strfmt.Registry) error {

	if swag.IsZero(m.ClientExtensions) { // not required
		return nil
	}

	if m.ClientExtensions != nil {

		if err := m.ClientExtensions.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("clientExtensions")
			}
			return err
		}
	}

	return nil
}

var trailingStopLossOrderRequestTypeTimeInForcePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["GTC","GTD","GFD","FOK","IOC"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		trailingStopLossOrderRequestTypeTimeInForcePropEnum = append(trailingStopLossOrderRequestTypeTimeInForcePropEnum, v)
	}
}

const (
	// TrailingStopLossOrderRequestTimeInForceGTC captures enum value "GTC"
	TrailingStopLossOrderRequestTimeInForceGTC string = "GTC"
	// TrailingStopLossOrderRequestTimeInForceGTD captures enum value "GTD"
	TrailingStopLossOrderRequestTimeInForceGTD string = "GTD"
	// TrailingStopLossOrderRequestTimeInForceGFD captures enum value "GFD"
	TrailingStopLossOrderRequestTimeInForceGFD string = "GFD"
	// TrailingStopLossOrderRequestTimeInForceFOK captures enum value "FOK"
	TrailingStopLossOrderRequestTimeInForceFOK string = "FOK"
	// TrailingStopLossOrderRequestTimeInForceIOC captures enum value "IOC"
	TrailingStopLossOrderRequestTimeInForceIOC string = "IOC"
)

// prop value enum
func (m *TrailingStopLossOrderRequest) validateTimeInForceEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, trailingStopLossOrderRequestTypeTimeInForcePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *TrailingStopLossOrderRequest) validateTimeInForce(formats strfmt.Registry) error {

	if swag.IsZero(m.TimeInForce) { // not required
		return nil
	}

	// value enum
	if err := m.validateTimeInForceEnum("timeInForce", "body", m.TimeInForce); err != nil {
		return err
	}

	return nil
}

var trailingStopLossOrderRequestTypeTriggerConditionPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["DEFAULT","INVERSE","BID","ASK","MID"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		trailingStopLossOrderRequestTypeTriggerConditionPropEnum = append(trailingStopLossOrderRequestTypeTriggerConditionPropEnum, v)
	}
}

const (
	// TrailingStopLossOrderRequestTriggerConditionDEFAULT captures enum value "DEFAULT"
	TrailingStopLossOrderRequestTriggerConditionDEFAULT string = "DEFAULT"
	// TrailingStopLossOrderRequestTriggerConditionINVERSE captures enum value "INVERSE"
	TrailingStopLossOrderRequestTriggerConditionINVERSE string = "INVERSE"
	// TrailingStopLossOrderRequestTriggerConditionBID captures enum value "BID"
	TrailingStopLossOrderRequestTriggerConditionBID string = "BID"
	// TrailingStopLossOrderRequestTriggerConditionASK captures enum value "ASK"
	TrailingStopLossOrderRequestTriggerConditionASK string = "ASK"
	// TrailingStopLossOrderRequestTriggerConditionMID captures enum value "MID"
	TrailingStopLossOrderRequestTriggerConditionMID string = "MID"
)

// prop value enum
func (m *TrailingStopLossOrderRequest) validateTriggerConditionEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, trailingStopLossOrderRequestTypeTriggerConditionPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *TrailingStopLossOrderRequest) validateTriggerCondition(formats strfmt.Registry) error {

	if swag.IsZero(m.TriggerCondition) { // not required
		return nil
	}

	// value enum
	if err := m.validateTriggerConditionEnum("triggerCondition", "body", m.TriggerCondition); err != nil {
		return err
	}

	return nil
}

var trailingStopLossOrderRequestTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["MARKET","LIMIT","STOP","MARKET_IF_TOUCHED","TAKE_PROFIT","STOP_LOSS","TRAILING_STOP_LOSS"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		trailingStopLossOrderRequestTypeTypePropEnum = append(trailingStopLossOrderRequestTypeTypePropEnum, v)
	}
}

const (
	// TrailingStopLossOrderRequestTypeMARKET captures enum value "MARKET"
	TrailingStopLossOrderRequestTypeMARKET string = "MARKET"
	// TrailingStopLossOrderRequestTypeLIMIT captures enum value "LIMIT"
	TrailingStopLossOrderRequestTypeLIMIT string = "LIMIT"
	// TrailingStopLossOrderRequestTypeSTOP captures enum value "STOP"
	TrailingStopLossOrderRequestTypeSTOP string = "STOP"
	// TrailingStopLossOrderRequestTypeMARKETIFTOUCHED captures enum value "MARKET_IF_TOUCHED"
	TrailingStopLossOrderRequestTypeMARKETIFTOUCHED string = "MARKET_IF_TOUCHED"
	// TrailingStopLossOrderRequestTypeTAKEPROFIT captures enum value "TAKE_PROFIT"
	TrailingStopLossOrderRequestTypeTAKEPROFIT string = "TAKE_PROFIT"
	// TrailingStopLossOrderRequestTypeSTOPLOSS captures enum value "STOP_LOSS"
	TrailingStopLossOrderRequestTypeSTOPLOSS string = "STOP_LOSS"
	// TrailingStopLossOrderRequestTypeTRAILINGSTOPLOSS captures enum value "TRAILING_STOP_LOSS"
	TrailingStopLossOrderRequestTypeTRAILINGSTOPLOSS string = "TRAILING_STOP_LOSS"
)

// prop value enum
func (m *TrailingStopLossOrderRequest) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, trailingStopLossOrderRequestTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *TrailingStopLossOrderRequest) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TrailingStopLossOrderRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TrailingStopLossOrderRequest) UnmarshalBinary(b []byte) error {
	var res TrailingStopLossOrderRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
