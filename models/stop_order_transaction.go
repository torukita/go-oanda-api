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

// StopOrderTransaction A StopOrderTransaction represents the creation of a Stop Order in the user's Account.
// swagger:model StopOrderTransaction
type StopOrderTransaction struct {

	// The ID of the Account the Transaction was created for.
	AccountID string `json:"accountID,omitempty"`

	// The ID of the "batch" that the Transaction belongs to. Transactions in the same batch are applied to the Account simultaneously.
	BatchID string `json:"batchID,omitempty"`

	// The ID of the Transaction that cancels the replaced Order (only provided if this Order replaces an existing Order).
	CancellingTransactionID string `json:"cancellingTransactionID,omitempty"`

	// client extensions
	ClientExtensions *ClientExtensions `json:"clientExtensions,omitempty"`

	// The date/time when the Stop Order will be cancelled if its timeInForce is "GTD".
	GtdTime string `json:"gtdTime,omitempty"`

	// The Transaction's Identifier.
	ID string `json:"id,omitempty"`

	// The Stop Order's Instrument.
	Instrument string `json:"instrument,omitempty"`

	// Specification of how Positions in the Account are modified when the Order is filled.
	PositionFill string `json:"positionFill,omitempty"`

	// The price threshold specified for the Stop Order. The Stop Order will only be filled by a market price that is equal to or worse than this price.
	Price string `json:"price,omitempty"`

	// The worst market price that may be used to fill this Stop Order. If the market gaps and crosses through both the price and the priceBound, the Stop Order will be cancelled instead of being filled.
	PriceBound string `json:"priceBound,omitempty"`

	// The reason that the Stop Order was initiated
	Reason string `json:"reason,omitempty"`

	// The ID of the Order that this Order replaces (only provided if this Order replaces an existing Order).
	ReplacesOrderID string `json:"replacesOrderID,omitempty"`

	// The Request ID of the request which generated the transaction.
	RequestID string `json:"requestID,omitempty"`

	// stop loss on fill
	StopLossOnFill *StopLossDetails `json:"stopLossOnFill,omitempty"`

	// take profit on fill
	TakeProfitOnFill *TakeProfitDetails `json:"takeProfitOnFill,omitempty"`

	// The date/time when the Transaction was created.
	Time string `json:"time,omitempty"`

	// The time-in-force requested for the Stop Order.
	TimeInForce string `json:"timeInForce,omitempty"`

	// trade client extensions
	TradeClientExtensions *ClientExtensions `json:"tradeClientExtensions,omitempty"`

	// trailing stop loss on fill
	TrailingStopLossOnFill *TrailingStopLossDetails `json:"trailingStopLossOnFill,omitempty"`

	// Specification of what component of a price should be used for comparison when determining if the Order should be filled.
	TriggerCondition string `json:"triggerCondition,omitempty"`

	// The Type of the Transaction. Always set to "STOP_ORDER" in a StopOrderTransaction.
	Type string `json:"type,omitempty"`

	// The quantity requested to be filled by the Stop Order. A posititive number of units results in a long Order, and a negative number of units results in a short Order.
	Units string `json:"units,omitempty"`

	// The ID of the user that initiated the creation of the Transaction.
	UserID int64 `json:"userID,omitempty"`
}

// Validate validates this stop order transaction
func (m *StopOrderTransaction) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateClientExtensions(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validatePositionFill(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateReason(formats); err != nil {
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

func (m *StopOrderTransaction) validateClientExtensions(formats strfmt.Registry) error {

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

var stopOrderTransactionTypePositionFillPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["OPEN_ONLY","REDUCE_FIRST","REDUCE_ONLY","DEFAULT"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		stopOrderTransactionTypePositionFillPropEnum = append(stopOrderTransactionTypePositionFillPropEnum, v)
	}
}

const (
	// StopOrderTransactionPositionFillOPENONLY captures enum value "OPEN_ONLY"
	StopOrderTransactionPositionFillOPENONLY string = "OPEN_ONLY"
	// StopOrderTransactionPositionFillREDUCEFIRST captures enum value "REDUCE_FIRST"
	StopOrderTransactionPositionFillREDUCEFIRST string = "REDUCE_FIRST"
	// StopOrderTransactionPositionFillREDUCEONLY captures enum value "REDUCE_ONLY"
	StopOrderTransactionPositionFillREDUCEONLY string = "REDUCE_ONLY"
	// StopOrderTransactionPositionFillDEFAULT captures enum value "DEFAULT"
	StopOrderTransactionPositionFillDEFAULT string = "DEFAULT"
)

// prop value enum
func (m *StopOrderTransaction) validatePositionFillEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, stopOrderTransactionTypePositionFillPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *StopOrderTransaction) validatePositionFill(formats strfmt.Registry) error {

	if swag.IsZero(m.PositionFill) { // not required
		return nil
	}

	// value enum
	if err := m.validatePositionFillEnum("positionFill", "body", m.PositionFill); err != nil {
		return err
	}

	return nil
}

var stopOrderTransactionTypeReasonPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["CLIENT_ORDER","REPLACEMENT"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		stopOrderTransactionTypeReasonPropEnum = append(stopOrderTransactionTypeReasonPropEnum, v)
	}
}

const (
	// StopOrderTransactionReasonCLIENTORDER captures enum value "CLIENT_ORDER"
	StopOrderTransactionReasonCLIENTORDER string = "CLIENT_ORDER"
	// StopOrderTransactionReasonREPLACEMENT captures enum value "REPLACEMENT"
	StopOrderTransactionReasonREPLACEMENT string = "REPLACEMENT"
)

// prop value enum
func (m *StopOrderTransaction) validateReasonEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, stopOrderTransactionTypeReasonPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *StopOrderTransaction) validateReason(formats strfmt.Registry) error {

	if swag.IsZero(m.Reason) { // not required
		return nil
	}

	// value enum
	if err := m.validateReasonEnum("reason", "body", m.Reason); err != nil {
		return err
	}

	return nil
}

func (m *StopOrderTransaction) validateStopLossOnFill(formats strfmt.Registry) error {

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

func (m *StopOrderTransaction) validateTakeProfitOnFill(formats strfmt.Registry) error {

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

var stopOrderTransactionTypeTimeInForcePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["GTC","GTD","GFD","FOK","IOC"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		stopOrderTransactionTypeTimeInForcePropEnum = append(stopOrderTransactionTypeTimeInForcePropEnum, v)
	}
}

const (
	// StopOrderTransactionTimeInForceGTC captures enum value "GTC"
	StopOrderTransactionTimeInForceGTC string = "GTC"
	// StopOrderTransactionTimeInForceGTD captures enum value "GTD"
	StopOrderTransactionTimeInForceGTD string = "GTD"
	// StopOrderTransactionTimeInForceGFD captures enum value "GFD"
	StopOrderTransactionTimeInForceGFD string = "GFD"
	// StopOrderTransactionTimeInForceFOK captures enum value "FOK"
	StopOrderTransactionTimeInForceFOK string = "FOK"
	// StopOrderTransactionTimeInForceIOC captures enum value "IOC"
	StopOrderTransactionTimeInForceIOC string = "IOC"
)

// prop value enum
func (m *StopOrderTransaction) validateTimeInForceEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, stopOrderTransactionTypeTimeInForcePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *StopOrderTransaction) validateTimeInForce(formats strfmt.Registry) error {

	if swag.IsZero(m.TimeInForce) { // not required
		return nil
	}

	// value enum
	if err := m.validateTimeInForceEnum("timeInForce", "body", m.TimeInForce); err != nil {
		return err
	}

	return nil
}

func (m *StopOrderTransaction) validateTradeClientExtensions(formats strfmt.Registry) error {

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

func (m *StopOrderTransaction) validateTrailingStopLossOnFill(formats strfmt.Registry) error {

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

var stopOrderTransactionTypeTriggerConditionPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["DEFAULT","INVERSE","BID","ASK","MID"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		stopOrderTransactionTypeTriggerConditionPropEnum = append(stopOrderTransactionTypeTriggerConditionPropEnum, v)
	}
}

const (
	// StopOrderTransactionTriggerConditionDEFAULT captures enum value "DEFAULT"
	StopOrderTransactionTriggerConditionDEFAULT string = "DEFAULT"
	// StopOrderTransactionTriggerConditionINVERSE captures enum value "INVERSE"
	StopOrderTransactionTriggerConditionINVERSE string = "INVERSE"
	// StopOrderTransactionTriggerConditionBID captures enum value "BID"
	StopOrderTransactionTriggerConditionBID string = "BID"
	// StopOrderTransactionTriggerConditionASK captures enum value "ASK"
	StopOrderTransactionTriggerConditionASK string = "ASK"
	// StopOrderTransactionTriggerConditionMID captures enum value "MID"
	StopOrderTransactionTriggerConditionMID string = "MID"
)

// prop value enum
func (m *StopOrderTransaction) validateTriggerConditionEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, stopOrderTransactionTypeTriggerConditionPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *StopOrderTransaction) validateTriggerCondition(formats strfmt.Registry) error {

	if swag.IsZero(m.TriggerCondition) { // not required
		return nil
	}

	// value enum
	if err := m.validateTriggerConditionEnum("triggerCondition", "body", m.TriggerCondition); err != nil {
		return err
	}

	return nil
}

var stopOrderTransactionTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["CREATE","CLOSE","REOPEN","CLIENT_CONFIGURE","CLIENT_CONFIGURE_REJECT","TRANSFER_FUNDS","TRANSFER_FUNDS_REJECT","MARKET_ORDER","MARKET_ORDER_REJECT","LIMIT_ORDER","LIMIT_ORDER_REJECT","STOP_ORDER","STOP_ORDER_REJECT","MARKET_IF_TOUCHED_ORDER","MARKET_IF_TOUCHED_ORDER_REJECT","TAKE_PROFIT_ORDER","TAKE_PROFIT_ORDER_REJECT","STOP_LOSS_ORDER","STOP_LOSS_ORDER_REJECT","TRAILING_STOP_LOSS_ORDER","TRAILING_STOP_LOSS_ORDER_REJECT","ORDER_FILL","ORDER_CANCEL","ORDER_CANCEL_REJECT","ORDER_CLIENT_EXTENSIONS_MODIFY","ORDER_CLIENT_EXTENSIONS_MODIFY_REJECT","TRADE_CLIENT_EXTENSIONS_MODIFY","TRADE_CLIENT_EXTENSIONS_MODIFY_REJECT","MARGIN_CALL_ENTER","MARGIN_CALL_EXTEND","MARGIN_CALL_EXIT","DELAYED_TRADE_CLOSURE","DAILY_FINANCING","RESET_RESETTABLE_PL"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		stopOrderTransactionTypeTypePropEnum = append(stopOrderTransactionTypeTypePropEnum, v)
	}
}

const (
	// StopOrderTransactionTypeCREATE captures enum value "CREATE"
	StopOrderTransactionTypeCREATE string = "CREATE"
	// StopOrderTransactionTypeCLOSE captures enum value "CLOSE"
	StopOrderTransactionTypeCLOSE string = "CLOSE"
	// StopOrderTransactionTypeREOPEN captures enum value "REOPEN"
	StopOrderTransactionTypeREOPEN string = "REOPEN"
	// StopOrderTransactionTypeCLIENTCONFIGURE captures enum value "CLIENT_CONFIGURE"
	StopOrderTransactionTypeCLIENTCONFIGURE string = "CLIENT_CONFIGURE"
	// StopOrderTransactionTypeCLIENTCONFIGUREREJECT captures enum value "CLIENT_CONFIGURE_REJECT"
	StopOrderTransactionTypeCLIENTCONFIGUREREJECT string = "CLIENT_CONFIGURE_REJECT"
	// StopOrderTransactionTypeTRANSFERFUNDS captures enum value "TRANSFER_FUNDS"
	StopOrderTransactionTypeTRANSFERFUNDS string = "TRANSFER_FUNDS"
	// StopOrderTransactionTypeTRANSFERFUNDSREJECT captures enum value "TRANSFER_FUNDS_REJECT"
	StopOrderTransactionTypeTRANSFERFUNDSREJECT string = "TRANSFER_FUNDS_REJECT"
	// StopOrderTransactionTypeMARKETORDER captures enum value "MARKET_ORDER"
	StopOrderTransactionTypeMARKETORDER string = "MARKET_ORDER"
	// StopOrderTransactionTypeMARKETORDERREJECT captures enum value "MARKET_ORDER_REJECT"
	StopOrderTransactionTypeMARKETORDERREJECT string = "MARKET_ORDER_REJECT"
	// StopOrderTransactionTypeLIMITORDER captures enum value "LIMIT_ORDER"
	StopOrderTransactionTypeLIMITORDER string = "LIMIT_ORDER"
	// StopOrderTransactionTypeLIMITORDERREJECT captures enum value "LIMIT_ORDER_REJECT"
	StopOrderTransactionTypeLIMITORDERREJECT string = "LIMIT_ORDER_REJECT"
	// StopOrderTransactionTypeSTOPORDER captures enum value "STOP_ORDER"
	StopOrderTransactionTypeSTOPORDER string = "STOP_ORDER"
	// StopOrderTransactionTypeSTOPORDERREJECT captures enum value "STOP_ORDER_REJECT"
	StopOrderTransactionTypeSTOPORDERREJECT string = "STOP_ORDER_REJECT"
	// StopOrderTransactionTypeMARKETIFTOUCHEDORDER captures enum value "MARKET_IF_TOUCHED_ORDER"
	StopOrderTransactionTypeMARKETIFTOUCHEDORDER string = "MARKET_IF_TOUCHED_ORDER"
	// StopOrderTransactionTypeMARKETIFTOUCHEDORDERREJECT captures enum value "MARKET_IF_TOUCHED_ORDER_REJECT"
	StopOrderTransactionTypeMARKETIFTOUCHEDORDERREJECT string = "MARKET_IF_TOUCHED_ORDER_REJECT"
	// StopOrderTransactionTypeTAKEPROFITORDER captures enum value "TAKE_PROFIT_ORDER"
	StopOrderTransactionTypeTAKEPROFITORDER string = "TAKE_PROFIT_ORDER"
	// StopOrderTransactionTypeTAKEPROFITORDERREJECT captures enum value "TAKE_PROFIT_ORDER_REJECT"
	StopOrderTransactionTypeTAKEPROFITORDERREJECT string = "TAKE_PROFIT_ORDER_REJECT"
	// StopOrderTransactionTypeSTOPLOSSORDER captures enum value "STOP_LOSS_ORDER"
	StopOrderTransactionTypeSTOPLOSSORDER string = "STOP_LOSS_ORDER"
	// StopOrderTransactionTypeSTOPLOSSORDERREJECT captures enum value "STOP_LOSS_ORDER_REJECT"
	StopOrderTransactionTypeSTOPLOSSORDERREJECT string = "STOP_LOSS_ORDER_REJECT"
	// StopOrderTransactionTypeTRAILINGSTOPLOSSORDER captures enum value "TRAILING_STOP_LOSS_ORDER"
	StopOrderTransactionTypeTRAILINGSTOPLOSSORDER string = "TRAILING_STOP_LOSS_ORDER"
	// StopOrderTransactionTypeTRAILINGSTOPLOSSORDERREJECT captures enum value "TRAILING_STOP_LOSS_ORDER_REJECT"
	StopOrderTransactionTypeTRAILINGSTOPLOSSORDERREJECT string = "TRAILING_STOP_LOSS_ORDER_REJECT"
	// StopOrderTransactionTypeORDERFILL captures enum value "ORDER_FILL"
	StopOrderTransactionTypeORDERFILL string = "ORDER_FILL"
	// StopOrderTransactionTypeORDERCANCEL captures enum value "ORDER_CANCEL"
	StopOrderTransactionTypeORDERCANCEL string = "ORDER_CANCEL"
	// StopOrderTransactionTypeORDERCANCELREJECT captures enum value "ORDER_CANCEL_REJECT"
	StopOrderTransactionTypeORDERCANCELREJECT string = "ORDER_CANCEL_REJECT"
	// StopOrderTransactionTypeORDERCLIENTEXTENSIONSMODIFY captures enum value "ORDER_CLIENT_EXTENSIONS_MODIFY"
	StopOrderTransactionTypeORDERCLIENTEXTENSIONSMODIFY string = "ORDER_CLIENT_EXTENSIONS_MODIFY"
	// StopOrderTransactionTypeORDERCLIENTEXTENSIONSMODIFYREJECT captures enum value "ORDER_CLIENT_EXTENSIONS_MODIFY_REJECT"
	StopOrderTransactionTypeORDERCLIENTEXTENSIONSMODIFYREJECT string = "ORDER_CLIENT_EXTENSIONS_MODIFY_REJECT"
	// StopOrderTransactionTypeTRADECLIENTEXTENSIONSMODIFY captures enum value "TRADE_CLIENT_EXTENSIONS_MODIFY"
	StopOrderTransactionTypeTRADECLIENTEXTENSIONSMODIFY string = "TRADE_CLIENT_EXTENSIONS_MODIFY"
	// StopOrderTransactionTypeTRADECLIENTEXTENSIONSMODIFYREJECT captures enum value "TRADE_CLIENT_EXTENSIONS_MODIFY_REJECT"
	StopOrderTransactionTypeTRADECLIENTEXTENSIONSMODIFYREJECT string = "TRADE_CLIENT_EXTENSIONS_MODIFY_REJECT"
	// StopOrderTransactionTypeMARGINCALLENTER captures enum value "MARGIN_CALL_ENTER"
	StopOrderTransactionTypeMARGINCALLENTER string = "MARGIN_CALL_ENTER"
	// StopOrderTransactionTypeMARGINCALLEXTEND captures enum value "MARGIN_CALL_EXTEND"
	StopOrderTransactionTypeMARGINCALLEXTEND string = "MARGIN_CALL_EXTEND"
	// StopOrderTransactionTypeMARGINCALLEXIT captures enum value "MARGIN_CALL_EXIT"
	StopOrderTransactionTypeMARGINCALLEXIT string = "MARGIN_CALL_EXIT"
	// StopOrderTransactionTypeDELAYEDTRADECLOSURE captures enum value "DELAYED_TRADE_CLOSURE"
	StopOrderTransactionTypeDELAYEDTRADECLOSURE string = "DELAYED_TRADE_CLOSURE"
	// StopOrderTransactionTypeDAILYFINANCING captures enum value "DAILY_FINANCING"
	StopOrderTransactionTypeDAILYFINANCING string = "DAILY_FINANCING"
	// StopOrderTransactionTypeRESETRESETTABLEPL captures enum value "RESET_RESETTABLE_PL"
	StopOrderTransactionTypeRESETRESETTABLEPL string = "RESET_RESETTABLE_PL"
)

// prop value enum
func (m *StopOrderTransaction) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, stopOrderTransactionTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *StopOrderTransaction) validateType(formats strfmt.Registry) error {

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
func (m *StopOrderTransaction) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StopOrderTransaction) UnmarshalBinary(b []byte) error {
	var res StopOrderTransaction
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
