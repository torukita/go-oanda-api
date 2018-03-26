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

// TakeProfitOrderRequest A TakeProfitOrderRequest specifies the parameters that may be set when creating a Take Profit Order.
// swagger:model TakeProfitOrderRequest
type TakeProfitOrderRequest struct {

	// client extensions
	ClientExtensions *ClientExtensions `json:"clientExtensions,omitempty"`

	// The client ID of the Trade to be closed when the price threshold is breached.
	ClientTradeID string `json:"clientTradeID,omitempty"`

	// The date/time when the TakeProfit Order will be cancelled if its timeInForce is "GTD".
	GtdTime string `json:"gtdTime,omitempty"`

	// The price threshold specified for the TakeProfit Order. The associated Trade will be closed by a market price that is equal to or better than this threshold.
	Price string `json:"price,omitempty"`

	// The time-in-force requested for the TakeProfit Order. Restricted to "GTC", "GFD" and "GTD" for TakeProfit Orders.
	TimeInForce string `json:"timeInForce,omitempty"`

	// The ID of the Trade to close when the price threshold is breached.
	TradeID string `json:"tradeID,omitempty"`

	// Specification of what component of a price should be used for comparison when determining if the Order should be filled.
	TriggerCondition string `json:"triggerCondition,omitempty"`

	// The type of the Order to Create. Must be set to "TAKE_PROFIT" when creating a Take Profit Order.
	Type string `json:"type,omitempty"`
}

// Validate validates this take profit order request
func (m *TakeProfitOrderRequest) Validate(formats strfmt.Registry) error {
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

func (m *TakeProfitOrderRequest) validateClientExtensions(formats strfmt.Registry) error {

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

var takeProfitOrderRequestTypeTimeInForcePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["GTC","GTD","GFD","FOK","IOC"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		takeProfitOrderRequestTypeTimeInForcePropEnum = append(takeProfitOrderRequestTypeTimeInForcePropEnum, v)
	}
}

const (
	// TakeProfitOrderRequestTimeInForceGTC captures enum value "GTC"
	TakeProfitOrderRequestTimeInForceGTC string = "GTC"
	// TakeProfitOrderRequestTimeInForceGTD captures enum value "GTD"
	TakeProfitOrderRequestTimeInForceGTD string = "GTD"
	// TakeProfitOrderRequestTimeInForceGFD captures enum value "GFD"
	TakeProfitOrderRequestTimeInForceGFD string = "GFD"
	// TakeProfitOrderRequestTimeInForceFOK captures enum value "FOK"
	TakeProfitOrderRequestTimeInForceFOK string = "FOK"
	// TakeProfitOrderRequestTimeInForceIOC captures enum value "IOC"
	TakeProfitOrderRequestTimeInForceIOC string = "IOC"
)

// prop value enum
func (m *TakeProfitOrderRequest) validateTimeInForceEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, takeProfitOrderRequestTypeTimeInForcePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *TakeProfitOrderRequest) validateTimeInForce(formats strfmt.Registry) error {

	if swag.IsZero(m.TimeInForce) { // not required
		return nil
	}

	// value enum
	if err := m.validateTimeInForceEnum("timeInForce", "body", m.TimeInForce); err != nil {
		return err
	}

	return nil
}

var takeProfitOrderRequestTypeTriggerConditionPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["DEFAULT","INVERSE","BID","ASK","MID"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		takeProfitOrderRequestTypeTriggerConditionPropEnum = append(takeProfitOrderRequestTypeTriggerConditionPropEnum, v)
	}
}

const (
	// TakeProfitOrderRequestTriggerConditionDEFAULT captures enum value "DEFAULT"
	TakeProfitOrderRequestTriggerConditionDEFAULT string = "DEFAULT"
	// TakeProfitOrderRequestTriggerConditionINVERSE captures enum value "INVERSE"
	TakeProfitOrderRequestTriggerConditionINVERSE string = "INVERSE"
	// TakeProfitOrderRequestTriggerConditionBID captures enum value "BID"
	TakeProfitOrderRequestTriggerConditionBID string = "BID"
	// TakeProfitOrderRequestTriggerConditionASK captures enum value "ASK"
	TakeProfitOrderRequestTriggerConditionASK string = "ASK"
	// TakeProfitOrderRequestTriggerConditionMID captures enum value "MID"
	TakeProfitOrderRequestTriggerConditionMID string = "MID"
)

// prop value enum
func (m *TakeProfitOrderRequest) validateTriggerConditionEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, takeProfitOrderRequestTypeTriggerConditionPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *TakeProfitOrderRequest) validateTriggerCondition(formats strfmt.Registry) error {

	if swag.IsZero(m.TriggerCondition) { // not required
		return nil
	}

	// value enum
	if err := m.validateTriggerConditionEnum("triggerCondition", "body", m.TriggerCondition); err != nil {
		return err
	}

	return nil
}

var takeProfitOrderRequestTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["MARKET","LIMIT","STOP","MARKET_IF_TOUCHED","TAKE_PROFIT","STOP_LOSS","TRAILING_STOP_LOSS"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		takeProfitOrderRequestTypeTypePropEnum = append(takeProfitOrderRequestTypeTypePropEnum, v)
	}
}

const (
	// TakeProfitOrderRequestTypeMARKET captures enum value "MARKET"
	TakeProfitOrderRequestTypeMARKET string = "MARKET"
	// TakeProfitOrderRequestTypeLIMIT captures enum value "LIMIT"
	TakeProfitOrderRequestTypeLIMIT string = "LIMIT"
	// TakeProfitOrderRequestTypeSTOP captures enum value "STOP"
	TakeProfitOrderRequestTypeSTOP string = "STOP"
	// TakeProfitOrderRequestTypeMARKETIFTOUCHED captures enum value "MARKET_IF_TOUCHED"
	TakeProfitOrderRequestTypeMARKETIFTOUCHED string = "MARKET_IF_TOUCHED"
	// TakeProfitOrderRequestTypeTAKEPROFIT captures enum value "TAKE_PROFIT"
	TakeProfitOrderRequestTypeTAKEPROFIT string = "TAKE_PROFIT"
	// TakeProfitOrderRequestTypeSTOPLOSS captures enum value "STOP_LOSS"
	TakeProfitOrderRequestTypeSTOPLOSS string = "STOP_LOSS"
	// TakeProfitOrderRequestTypeTRAILINGSTOPLOSS captures enum value "TRAILING_STOP_LOSS"
	TakeProfitOrderRequestTypeTRAILINGSTOPLOSS string = "TRAILING_STOP_LOSS"
)

// prop value enum
func (m *TakeProfitOrderRequest) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, takeProfitOrderRequestTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *TakeProfitOrderRequest) validateType(formats strfmt.Registry) error {

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
func (m *TakeProfitOrderRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TakeProfitOrderRequest) UnmarshalBinary(b []byte) error {
	var res TakeProfitOrderRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
