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

// StopOrder A StopOrder is an order that is created with a price threshold, and will only be filled by a price that is equal to or worse than the threshold.
// swagger:model StopOrder
type StopOrder struct {

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

	// The date/time when the Stop Order will be cancelled if its timeInForce is "GTD".
	GtdTime string `json:"gtdTime,omitempty"`

	// The Order's identifier, unique within the Order's Account.
	ID string `json:"id,omitempty"`

	// The Stop Order's Instrument.
	Instrument string `json:"instrument,omitempty"`

	// Specification of how Positions in the Account are modified when the Order is filled.
	PositionFill string `json:"positionFill,omitempty"`

	// The price threshold specified for the Stop Order. The Stop Order will only be filled by a market price that is equal to or worse than this price.
	Price string `json:"price,omitempty"`

	// The worst market price that may be used to fill this Stop Order. If the market gaps and crosses through both the price and the priceBound, the Stop Order will be cancelled instead of being filled.
	PriceBound string `json:"priceBound,omitempty"`

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

	// The time-in-force requested for the Stop Order.
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

	// The type of the Order. Always set to "STOP" for Stop Orders.
	Type string `json:"type,omitempty"`

	// The quantity requested to be filled by the Stop Order. A posititive number of units results in a long Order, and a negative number of units results in a short Order.
	Units string `json:"units,omitempty"`
}

// Validate validates this stop order
func (m *StopOrder) Validate(formats strfmt.Registry) error {
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

func (m *StopOrder) validateClientExtensions(formats strfmt.Registry) error {

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

var stopOrderTypePositionFillPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["OPEN_ONLY","REDUCE_FIRST","REDUCE_ONLY","DEFAULT"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		stopOrderTypePositionFillPropEnum = append(stopOrderTypePositionFillPropEnum, v)
	}
}

const (
	// StopOrderPositionFillOPENONLY captures enum value "OPEN_ONLY"
	StopOrderPositionFillOPENONLY string = "OPEN_ONLY"
	// StopOrderPositionFillREDUCEFIRST captures enum value "REDUCE_FIRST"
	StopOrderPositionFillREDUCEFIRST string = "REDUCE_FIRST"
	// StopOrderPositionFillREDUCEONLY captures enum value "REDUCE_ONLY"
	StopOrderPositionFillREDUCEONLY string = "REDUCE_ONLY"
	// StopOrderPositionFillDEFAULT captures enum value "DEFAULT"
	StopOrderPositionFillDEFAULT string = "DEFAULT"
)

// prop value enum
func (m *StopOrder) validatePositionFillEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, stopOrderTypePositionFillPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *StopOrder) validatePositionFill(formats strfmt.Registry) error {

	if swag.IsZero(m.PositionFill) { // not required
		return nil
	}

	// value enum
	if err := m.validatePositionFillEnum("positionFill", "body", m.PositionFill); err != nil {
		return err
	}

	return nil
}

var stopOrderTypeStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["PENDING","FILLED","TRIGGERED","CANCELLED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		stopOrderTypeStatePropEnum = append(stopOrderTypeStatePropEnum, v)
	}
}

const (
	// StopOrderStatePENDING captures enum value "PENDING"
	StopOrderStatePENDING string = "PENDING"
	// StopOrderStateFILLED captures enum value "FILLED"
	StopOrderStateFILLED string = "FILLED"
	// StopOrderStateTRIGGERED captures enum value "TRIGGERED"
	StopOrderStateTRIGGERED string = "TRIGGERED"
	// StopOrderStateCANCELLED captures enum value "CANCELLED"
	StopOrderStateCANCELLED string = "CANCELLED"
)

// prop value enum
func (m *StopOrder) validateStateEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, stopOrderTypeStatePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *StopOrder) validateState(formats strfmt.Registry) error {

	if swag.IsZero(m.State) { // not required
		return nil
	}

	// value enum
	if err := m.validateStateEnum("state", "body", m.State); err != nil {
		return err
	}

	return nil
}

func (m *StopOrder) validateStopLossOnFill(formats strfmt.Registry) error {

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

func (m *StopOrder) validateTakeProfitOnFill(formats strfmt.Registry) error {

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

var stopOrderTypeTimeInForcePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["GTC","GTD","GFD","FOK","IOC"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		stopOrderTypeTimeInForcePropEnum = append(stopOrderTypeTimeInForcePropEnum, v)
	}
}

const (
	// StopOrderTimeInForceGTC captures enum value "GTC"
	StopOrderTimeInForceGTC string = "GTC"
	// StopOrderTimeInForceGTD captures enum value "GTD"
	StopOrderTimeInForceGTD string = "GTD"
	// StopOrderTimeInForceGFD captures enum value "GFD"
	StopOrderTimeInForceGFD string = "GFD"
	// StopOrderTimeInForceFOK captures enum value "FOK"
	StopOrderTimeInForceFOK string = "FOK"
	// StopOrderTimeInForceIOC captures enum value "IOC"
	StopOrderTimeInForceIOC string = "IOC"
)

// prop value enum
func (m *StopOrder) validateTimeInForceEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, stopOrderTypeTimeInForcePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *StopOrder) validateTimeInForce(formats strfmt.Registry) error {

	if swag.IsZero(m.TimeInForce) { // not required
		return nil
	}

	// value enum
	if err := m.validateTimeInForceEnum("timeInForce", "body", m.TimeInForce); err != nil {
		return err
	}

	return nil
}

func (m *StopOrder) validateTradeClientExtensions(formats strfmt.Registry) error {

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

func (m *StopOrder) validateTradeClosedIds(formats strfmt.Registry) error {

	if swag.IsZero(m.TradeClosedIds) { // not required
		return nil
	}

	return nil
}

func (m *StopOrder) validateTrailingStopLossOnFill(formats strfmt.Registry) error {

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

var stopOrderTypeTriggerConditionPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["DEFAULT","INVERSE","BID","ASK","MID"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		stopOrderTypeTriggerConditionPropEnum = append(stopOrderTypeTriggerConditionPropEnum, v)
	}
}

const (
	// StopOrderTriggerConditionDEFAULT captures enum value "DEFAULT"
	StopOrderTriggerConditionDEFAULT string = "DEFAULT"
	// StopOrderTriggerConditionINVERSE captures enum value "INVERSE"
	StopOrderTriggerConditionINVERSE string = "INVERSE"
	// StopOrderTriggerConditionBID captures enum value "BID"
	StopOrderTriggerConditionBID string = "BID"
	// StopOrderTriggerConditionASK captures enum value "ASK"
	StopOrderTriggerConditionASK string = "ASK"
	// StopOrderTriggerConditionMID captures enum value "MID"
	StopOrderTriggerConditionMID string = "MID"
)

// prop value enum
func (m *StopOrder) validateTriggerConditionEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, stopOrderTypeTriggerConditionPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *StopOrder) validateTriggerCondition(formats strfmt.Registry) error {

	if swag.IsZero(m.TriggerCondition) { // not required
		return nil
	}

	// value enum
	if err := m.validateTriggerConditionEnum("triggerCondition", "body", m.TriggerCondition); err != nil {
		return err
	}

	return nil
}

var stopOrderTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["MARKET","LIMIT","STOP","MARKET_IF_TOUCHED","TAKE_PROFIT","STOP_LOSS","TRAILING_STOP_LOSS"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		stopOrderTypeTypePropEnum = append(stopOrderTypeTypePropEnum, v)
	}
}

const (
	// StopOrderTypeMARKET captures enum value "MARKET"
	StopOrderTypeMARKET string = "MARKET"
	// StopOrderTypeLIMIT captures enum value "LIMIT"
	StopOrderTypeLIMIT string = "LIMIT"
	// StopOrderTypeSTOP captures enum value "STOP"
	StopOrderTypeSTOP string = "STOP"
	// StopOrderTypeMARKETIFTOUCHED captures enum value "MARKET_IF_TOUCHED"
	StopOrderTypeMARKETIFTOUCHED string = "MARKET_IF_TOUCHED"
	// StopOrderTypeTAKEPROFIT captures enum value "TAKE_PROFIT"
	StopOrderTypeTAKEPROFIT string = "TAKE_PROFIT"
	// StopOrderTypeSTOPLOSS captures enum value "STOP_LOSS"
	StopOrderTypeSTOPLOSS string = "STOP_LOSS"
	// StopOrderTypeTRAILINGSTOPLOSS captures enum value "TRAILING_STOP_LOSS"
	StopOrderTypeTRAILINGSTOPLOSS string = "TRAILING_STOP_LOSS"
)

// prop value enum
func (m *StopOrder) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, stopOrderTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *StopOrder) validateType(formats strfmt.Registry) error {

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
func (m *StopOrder) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StopOrder) UnmarshalBinary(b []byte) error {
	var res StopOrder
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}