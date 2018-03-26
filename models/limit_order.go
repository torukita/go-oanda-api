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

// LimitOrder A LimitOrder is an order that is created with a price threshold, and will only be filled by a price that is equal to or better than the threshold.
// swagger:model LimitOrder
type LimitOrder struct {

	// Date/time when the Order was cancelled (only provided when the state of the Order is CANCELLED)
	CancelledTime string `json:"cancelledTime,omitempty"`

	// ID of the Transaction that cancelled the Order (only provided when the Order's state is CANCELLED)
	CancellingTransactionID string `json:"cancellingTransactionID,omitempty"`

	// client extensions
	ClientExtensions *ClientExtensions `json:"clientExtensions,omitempty"`

	// The time when the Order was created.
	CreateTime string `json:"createTime,omitempty"`

	// Date/time when the Order was filled (only provided when the Order's state is FILLED)
	FilledTime string `json:"filledTime,omitempty"`

	// ID of the Transaction that filled this Order (only provided when the Order's state is FILLED)
	FillingTransactionID string `json:"fillingTransactionID,omitempty"`

	// The date/time when the Limit Order will be cancelled if its timeInForce is "GTD".
	GtdTime string `json:"gtdTime,omitempty"`

	// The Order's identifier, unique within the Order's Account.
	ID string `json:"id,omitempty"`

	// The Limit Order's Instrument.
	Instrument string `json:"instrument,omitempty"`

	// Specification of how Positions in the Account are modified when the Order is filled.
	PositionFill string `json:"positionFill,omitempty"`

	// The price threshold specified for the Limit Order. The Limit Order will only be filled by a market price that is equal to or better than this price.
	Price string `json:"price,omitempty"`

	// The ID of the Order that replaced this Order (only provided if this Order was cancelled as part of a cancel/replace).
	ReplacedByOrderID string `json:"replacedByOrderID,omitempty"`

	// The ID of the Order that was replaced by this Order (only provided if this Order was created as part of a cancel/replace).
	ReplacesOrderID string `json:"replacesOrderID,omitempty"`

	// The current state of the Order.
	State string `json:"state,omitempty"`

	// stop loss on fill
	StopLossOnFill *StopLossDetails `json:"stopLossOnFill,omitempty"`

	// take profit on fill
	TakeProfitOnFill *TakeProfitDetails `json:"takeProfitOnFill,omitempty"`

	// The time-in-force requested for the Limit Order.
	TimeInForce string `json:"timeInForce,omitempty"`

	// trade client extensions
	TradeClientExtensions *ClientExtensions `json:"tradeClientExtensions,omitempty"`

	// Trade IDs of Trades closed when the Order was filled (only provided when the Order's state is FILLED and one or more Trades were closed as a result of the fill)
	TradeClosedIds []string `json:"tradeClosedIDs"`

	// Trade ID of Trade opened when the Order was filled (only provided when the Order's state is FILLED and a Trade was opened as a result of the fill)
	TradeOpenedID string `json:"tradeOpenedID,omitempty"`

	// Trade ID of Trade reduced when the Order was filled (only provided when the Order's state is FILLED and a Trade was reduced as a result of the fill)
	TradeReducedID string `json:"tradeReducedID,omitempty"`

	// trailing stop loss on fill
	TrailingStopLossOnFill *TrailingStopLossDetails `json:"trailingStopLossOnFill,omitempty"`

	// Specification of what component of a price should be used for comparison when determining if the Order should be filled.
	TriggerCondition string `json:"triggerCondition,omitempty"`

	// The type of the Order. Always set to "LIMIT" for Limit Orders.
	Type string `json:"type,omitempty"`

	// The quantity requested to be filled by the Limit Order. A posititive number of units results in a long Order, and a negative number of units results in a short Order.
	Units string `json:"units,omitempty"`
}

// Validate validates this limit order
func (m *LimitOrder) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateClientExtensions(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validatePositionFill(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateState(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateStopLossOnFill(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateTakeProfitOnFill(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateTimeInForce(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateTradeClientExtensions(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateTradeClosedIds(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateTrailingStopLossOnFill(formats); err != nil {
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

func (m *LimitOrder) validateClientExtensions(formats strfmt.Registry) error {

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

var limitOrderTypePositionFillPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["OPEN_ONLY","REDUCE_FIRST","REDUCE_ONLY","DEFAULT"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		limitOrderTypePositionFillPropEnum = append(limitOrderTypePositionFillPropEnum, v)
	}
}

const (
	// LimitOrderPositionFillOPENONLY captures enum value "OPEN_ONLY"
	LimitOrderPositionFillOPENONLY string = "OPEN_ONLY"
	// LimitOrderPositionFillREDUCEFIRST captures enum value "REDUCE_FIRST"
	LimitOrderPositionFillREDUCEFIRST string = "REDUCE_FIRST"
	// LimitOrderPositionFillREDUCEONLY captures enum value "REDUCE_ONLY"
	LimitOrderPositionFillREDUCEONLY string = "REDUCE_ONLY"
	// LimitOrderPositionFillDEFAULT captures enum value "DEFAULT"
	LimitOrderPositionFillDEFAULT string = "DEFAULT"
)

// prop value enum
func (m *LimitOrder) validatePositionFillEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, limitOrderTypePositionFillPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *LimitOrder) validatePositionFill(formats strfmt.Registry) error {

	if swag.IsZero(m.PositionFill) { // not required
		return nil
	}

	// value enum
	if err := m.validatePositionFillEnum("positionFill", "body", m.PositionFill); err != nil {
		return err
	}

	return nil
}

var limitOrderTypeStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["PENDING","FILLED","TRIGGERED","CANCELLED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		limitOrderTypeStatePropEnum = append(limitOrderTypeStatePropEnum, v)
	}
}

const (
	// LimitOrderStatePENDING captures enum value "PENDING"
	LimitOrderStatePENDING string = "PENDING"
	// LimitOrderStateFILLED captures enum value "FILLED"
	LimitOrderStateFILLED string = "FILLED"
	// LimitOrderStateTRIGGERED captures enum value "TRIGGERED"
	LimitOrderStateTRIGGERED string = "TRIGGERED"
	// LimitOrderStateCANCELLED captures enum value "CANCELLED"
	LimitOrderStateCANCELLED string = "CANCELLED"
)

// prop value enum
func (m *LimitOrder) validateStateEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, limitOrderTypeStatePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *LimitOrder) validateState(formats strfmt.Registry) error {

	if swag.IsZero(m.State) { // not required
		return nil
	}

	// value enum
	if err := m.validateStateEnum("state", "body", m.State); err != nil {
		return err
	}

	return nil
}

func (m *LimitOrder) validateStopLossOnFill(formats strfmt.Registry) error {

	if swag.IsZero(m.StopLossOnFill) { // not required
		return nil
	}

	if m.StopLossOnFill != nil {

		if err := m.StopLossOnFill.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("stopLossOnFill")
			}
			return err
		}
	}

	return nil
}

func (m *LimitOrder) validateTakeProfitOnFill(formats strfmt.Registry) error {

	if swag.IsZero(m.TakeProfitOnFill) { // not required
		return nil
	}

	if m.TakeProfitOnFill != nil {

		if err := m.TakeProfitOnFill.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("takeProfitOnFill")
			}
			return err
		}
	}

	return nil
}

var limitOrderTypeTimeInForcePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["GTC","GTD","GFD","FOK","IOC"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		limitOrderTypeTimeInForcePropEnum = append(limitOrderTypeTimeInForcePropEnum, v)
	}
}

const (
	// LimitOrderTimeInForceGTC captures enum value "GTC"
	LimitOrderTimeInForceGTC string = "GTC"
	// LimitOrderTimeInForceGTD captures enum value "GTD"
	LimitOrderTimeInForceGTD string = "GTD"
	// LimitOrderTimeInForceGFD captures enum value "GFD"
	LimitOrderTimeInForceGFD string = "GFD"
	// LimitOrderTimeInForceFOK captures enum value "FOK"
	LimitOrderTimeInForceFOK string = "FOK"
	// LimitOrderTimeInForceIOC captures enum value "IOC"
	LimitOrderTimeInForceIOC string = "IOC"
)

// prop value enum
func (m *LimitOrder) validateTimeInForceEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, limitOrderTypeTimeInForcePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *LimitOrder) validateTimeInForce(formats strfmt.Registry) error {

	if swag.IsZero(m.TimeInForce) { // not required
		return nil
	}

	// value enum
	if err := m.validateTimeInForceEnum("timeInForce", "body", m.TimeInForce); err != nil {
		return err
	}

	return nil
}

func (m *LimitOrder) validateTradeClientExtensions(formats strfmt.Registry) error {

	if swag.IsZero(m.TradeClientExtensions) { // not required
		return nil
	}

	if m.TradeClientExtensions != nil {

		if err := m.TradeClientExtensions.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("tradeClientExtensions")
			}
			return err
		}
	}

	return nil
}

func (m *LimitOrder) validateTradeClosedIds(formats strfmt.Registry) error {

	if swag.IsZero(m.TradeClosedIds) { // not required
		return nil
	}

	return nil
}

func (m *LimitOrder) validateTrailingStopLossOnFill(formats strfmt.Registry) error {

	if swag.IsZero(m.TrailingStopLossOnFill) { // not required
		return nil
	}

	if m.TrailingStopLossOnFill != nil {

		if err := m.TrailingStopLossOnFill.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("trailingStopLossOnFill")
			}
			return err
		}
	}

	return nil
}

var limitOrderTypeTriggerConditionPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["DEFAULT","INVERSE","BID","ASK","MID"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		limitOrderTypeTriggerConditionPropEnum = append(limitOrderTypeTriggerConditionPropEnum, v)
	}
}

const (
	// LimitOrderTriggerConditionDEFAULT captures enum value "DEFAULT"
	LimitOrderTriggerConditionDEFAULT string = "DEFAULT"
	// LimitOrderTriggerConditionINVERSE captures enum value "INVERSE"
	LimitOrderTriggerConditionINVERSE string = "INVERSE"
	// LimitOrderTriggerConditionBID captures enum value "BID"
	LimitOrderTriggerConditionBID string = "BID"
	// LimitOrderTriggerConditionASK captures enum value "ASK"
	LimitOrderTriggerConditionASK string = "ASK"
	// LimitOrderTriggerConditionMID captures enum value "MID"
	LimitOrderTriggerConditionMID string = "MID"
)

// prop value enum
func (m *LimitOrder) validateTriggerConditionEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, limitOrderTypeTriggerConditionPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *LimitOrder) validateTriggerCondition(formats strfmt.Registry) error {

	if swag.IsZero(m.TriggerCondition) { // not required
		return nil
	}

	// value enum
	if err := m.validateTriggerConditionEnum("triggerCondition", "body", m.TriggerCondition); err != nil {
		return err
	}

	return nil
}

var limitOrderTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["MARKET","LIMIT","STOP","MARKET_IF_TOUCHED","TAKE_PROFIT","STOP_LOSS","TRAILING_STOP_LOSS"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		limitOrderTypeTypePropEnum = append(limitOrderTypeTypePropEnum, v)
	}
}

const (
	// LimitOrderTypeMARKET captures enum value "MARKET"
	LimitOrderTypeMARKET string = "MARKET"
	// LimitOrderTypeLIMIT captures enum value "LIMIT"
	LimitOrderTypeLIMIT string = "LIMIT"
	// LimitOrderTypeSTOP captures enum value "STOP"
	LimitOrderTypeSTOP string = "STOP"
	// LimitOrderTypeMARKETIFTOUCHED captures enum value "MARKET_IF_TOUCHED"
	LimitOrderTypeMARKETIFTOUCHED string = "MARKET_IF_TOUCHED"
	// LimitOrderTypeTAKEPROFIT captures enum value "TAKE_PROFIT"
	LimitOrderTypeTAKEPROFIT string = "TAKE_PROFIT"
	// LimitOrderTypeSTOPLOSS captures enum value "STOP_LOSS"
	LimitOrderTypeSTOPLOSS string = "STOP_LOSS"
	// LimitOrderTypeTRAILINGSTOPLOSS captures enum value "TRAILING_STOP_LOSS"
	LimitOrderTypeTRAILINGSTOPLOSS string = "TRAILING_STOP_LOSS"
)

// prop value enum
func (m *LimitOrder) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, limitOrderTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *LimitOrder) validateType(formats strfmt.Registry) error {

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
func (m *LimitOrder) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LimitOrder) UnmarshalBinary(b []byte) error {
	var res LimitOrder
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
