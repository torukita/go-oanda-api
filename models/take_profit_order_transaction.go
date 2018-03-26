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

// TakeProfitOrderTransaction A TakeProfitOrderTransaction represents the creation of a TakeProfit Order in the user's Account.
// swagger:model TakeProfitOrderTransaction
type TakeProfitOrderTransaction struct {

	// The ID of the Account the Transaction was created for.
	AccountID string `json:"accountID,omitempty"`

	// The ID of the "batch" that the Transaction belongs to. Transactions in the same batch are applied to the Account simultaneously.
	BatchID string `json:"batchID,omitempty"`

	// The ID of the Transaction that cancels the replaced Order (only provided if this Order replaces an existing Order).
	CancellingTransactionID string `json:"cancellingTransactionID,omitempty"`

	// client extensions
	ClientExtensions *ClientExtensions `json:"clientExtensions,omitempty"`

	// The client ID of the Trade to be closed when the price threshold is breached.
	ClientTradeID string `json:"clientTradeID,omitempty"`

	// The date/time when the TakeProfit Order will be cancelled if its timeInForce is "GTD".
	GtdTime string `json:"gtdTime,omitempty"`

	// The Transaction's Identifier.
	ID string `json:"id,omitempty"`

	// The ID of the OrderFill Transaction that caused this Order to be created (only provided if this Order was created automatically when another Order was filled).
	OrderFillTransactionID string `json:"orderFillTransactionID,omitempty"`

	// The price threshold specified for the TakeProfit Order. The associated Trade will be closed by a market price that is equal to or better than this threshold.
	Price string `json:"price,omitempty"`

	// The reason that the Take Profit Order was initiated
	Reason string `json:"reason,omitempty"`

	// The ID of the Order that this Order replaces (only provided if this Order replaces an existing Order).
	ReplacesOrderID string `json:"replacesOrderID,omitempty"`

	// The Request ID of the request which generated the transaction.
	RequestID string `json:"requestID,omitempty"`

	// The date/time when the Transaction was created.
	Time string `json:"time,omitempty"`

	// The time-in-force requested for the TakeProfit Order. Restricted to "GTC", "GFD" and "GTD" for TakeProfit Orders.
	TimeInForce string `json:"timeInForce,omitempty"`

	// The ID of the Trade to close when the price threshold is breached.
	TradeID string `json:"tradeID,omitempty"`

	// Specification of what component of a price should be used for comparison when determining if the Order should be filled.
	TriggerCondition string `json:"triggerCondition,omitempty"`

	// The Type of the Transaction. Always set to "TAKE_PROFIT_ORDER" in a TakeProfitOrderTransaction.
	Type string `json:"type,omitempty"`

	// The ID of the user that initiated the creation of the Transaction.
	UserID int64 `json:"userID,omitempty"`
}

// Validate validates this take profit order transaction
func (m *TakeProfitOrderTransaction) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateClientExtensions(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateReason(formats); err != nil {
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

func (m *TakeProfitOrderTransaction) validateClientExtensions(formats strfmt.Registry) error {

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

var takeProfitOrderTransactionTypeReasonPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["CLIENT_ORDER","REPLACEMENT","ON_FILL"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		takeProfitOrderTransactionTypeReasonPropEnum = append(takeProfitOrderTransactionTypeReasonPropEnum, v)
	}
}

const (
	// TakeProfitOrderTransactionReasonCLIENTORDER captures enum value "CLIENT_ORDER"
	TakeProfitOrderTransactionReasonCLIENTORDER string = "CLIENT_ORDER"
	// TakeProfitOrderTransactionReasonREPLACEMENT captures enum value "REPLACEMENT"
	TakeProfitOrderTransactionReasonREPLACEMENT string = "REPLACEMENT"
	// TakeProfitOrderTransactionReasonONFILL captures enum value "ON_FILL"
	TakeProfitOrderTransactionReasonONFILL string = "ON_FILL"
)

// prop value enum
func (m *TakeProfitOrderTransaction) validateReasonEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, takeProfitOrderTransactionTypeReasonPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *TakeProfitOrderTransaction) validateReason(formats strfmt.Registry) error {

	if swag.IsZero(m.Reason) { // not required
		return nil
	}

	// value enum
	if err := m.validateReasonEnum("reason", "body", m.Reason); err != nil {
		return err
	}

	return nil
}

var takeProfitOrderTransactionTypeTimeInForcePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["GTC","GTD","GFD","FOK","IOC"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		takeProfitOrderTransactionTypeTimeInForcePropEnum = append(takeProfitOrderTransactionTypeTimeInForcePropEnum, v)
	}
}

const (
	// TakeProfitOrderTransactionTimeInForceGTC captures enum value "GTC"
	TakeProfitOrderTransactionTimeInForceGTC string = "GTC"
	// TakeProfitOrderTransactionTimeInForceGTD captures enum value "GTD"
	TakeProfitOrderTransactionTimeInForceGTD string = "GTD"
	// TakeProfitOrderTransactionTimeInForceGFD captures enum value "GFD"
	TakeProfitOrderTransactionTimeInForceGFD string = "GFD"
	// TakeProfitOrderTransactionTimeInForceFOK captures enum value "FOK"
	TakeProfitOrderTransactionTimeInForceFOK string = "FOK"
	// TakeProfitOrderTransactionTimeInForceIOC captures enum value "IOC"
	TakeProfitOrderTransactionTimeInForceIOC string = "IOC"
)

// prop value enum
func (m *TakeProfitOrderTransaction) validateTimeInForceEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, takeProfitOrderTransactionTypeTimeInForcePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *TakeProfitOrderTransaction) validateTimeInForce(formats strfmt.Registry) error {

	if swag.IsZero(m.TimeInForce) { // not required
		return nil
	}

	// value enum
	if err := m.validateTimeInForceEnum("timeInForce", "body", m.TimeInForce); err != nil {
		return err
	}

	return nil
}

var takeProfitOrderTransactionTypeTriggerConditionPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["DEFAULT","INVERSE","BID","ASK","MID"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		takeProfitOrderTransactionTypeTriggerConditionPropEnum = append(takeProfitOrderTransactionTypeTriggerConditionPropEnum, v)
	}
}

const (
	// TakeProfitOrderTransactionTriggerConditionDEFAULT captures enum value "DEFAULT"
	TakeProfitOrderTransactionTriggerConditionDEFAULT string = "DEFAULT"
	// TakeProfitOrderTransactionTriggerConditionINVERSE captures enum value "INVERSE"
	TakeProfitOrderTransactionTriggerConditionINVERSE string = "INVERSE"
	// TakeProfitOrderTransactionTriggerConditionBID captures enum value "BID"
	TakeProfitOrderTransactionTriggerConditionBID string = "BID"
	// TakeProfitOrderTransactionTriggerConditionASK captures enum value "ASK"
	TakeProfitOrderTransactionTriggerConditionASK string = "ASK"
	// TakeProfitOrderTransactionTriggerConditionMID captures enum value "MID"
	TakeProfitOrderTransactionTriggerConditionMID string = "MID"
)

// prop value enum
func (m *TakeProfitOrderTransaction) validateTriggerConditionEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, takeProfitOrderTransactionTypeTriggerConditionPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *TakeProfitOrderTransaction) validateTriggerCondition(formats strfmt.Registry) error {

	if swag.IsZero(m.TriggerCondition) { // not required
		return nil
	}

	// value enum
	if err := m.validateTriggerConditionEnum("triggerCondition", "body", m.TriggerCondition); err != nil {
		return err
	}

	return nil
}

var takeProfitOrderTransactionTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["CREATE","CLOSE","REOPEN","CLIENT_CONFIGURE","CLIENT_CONFIGURE_REJECT","TRANSFER_FUNDS","TRANSFER_FUNDS_REJECT","MARKET_ORDER","MARKET_ORDER_REJECT","LIMIT_ORDER","LIMIT_ORDER_REJECT","STOP_ORDER","STOP_ORDER_REJECT","MARKET_IF_TOUCHED_ORDER","MARKET_IF_TOUCHED_ORDER_REJECT","TAKE_PROFIT_ORDER","TAKE_PROFIT_ORDER_REJECT","STOP_LOSS_ORDER","STOP_LOSS_ORDER_REJECT","TRAILING_STOP_LOSS_ORDER","TRAILING_STOP_LOSS_ORDER_REJECT","ORDER_FILL","ORDER_CANCEL","ORDER_CANCEL_REJECT","ORDER_CLIENT_EXTENSIONS_MODIFY","ORDER_CLIENT_EXTENSIONS_MODIFY_REJECT","TRADE_CLIENT_EXTENSIONS_MODIFY","TRADE_CLIENT_EXTENSIONS_MODIFY_REJECT","MARGIN_CALL_ENTER","MARGIN_CALL_EXTEND","MARGIN_CALL_EXIT","DELAYED_TRADE_CLOSURE","DAILY_FINANCING","RESET_RESETTABLE_PL"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		takeProfitOrderTransactionTypeTypePropEnum = append(takeProfitOrderTransactionTypeTypePropEnum, v)
	}
}

const (
	// TakeProfitOrderTransactionTypeCREATE captures enum value "CREATE"
	TakeProfitOrderTransactionTypeCREATE string = "CREATE"
	// TakeProfitOrderTransactionTypeCLOSE captures enum value "CLOSE"
	TakeProfitOrderTransactionTypeCLOSE string = "CLOSE"
	// TakeProfitOrderTransactionTypeREOPEN captures enum value "REOPEN"
	TakeProfitOrderTransactionTypeREOPEN string = "REOPEN"
	// TakeProfitOrderTransactionTypeCLIENTCONFIGURE captures enum value "CLIENT_CONFIGURE"
	TakeProfitOrderTransactionTypeCLIENTCONFIGURE string = "CLIENT_CONFIGURE"
	// TakeProfitOrderTransactionTypeCLIENTCONFIGUREREJECT captures enum value "CLIENT_CONFIGURE_REJECT"
	TakeProfitOrderTransactionTypeCLIENTCONFIGUREREJECT string = "CLIENT_CONFIGURE_REJECT"
	// TakeProfitOrderTransactionTypeTRANSFERFUNDS captures enum value "TRANSFER_FUNDS"
	TakeProfitOrderTransactionTypeTRANSFERFUNDS string = "TRANSFER_FUNDS"
	// TakeProfitOrderTransactionTypeTRANSFERFUNDSREJECT captures enum value "TRANSFER_FUNDS_REJECT"
	TakeProfitOrderTransactionTypeTRANSFERFUNDSREJECT string = "TRANSFER_FUNDS_REJECT"
	// TakeProfitOrderTransactionTypeMARKETORDER captures enum value "MARKET_ORDER"
	TakeProfitOrderTransactionTypeMARKETORDER string = "MARKET_ORDER"
	// TakeProfitOrderTransactionTypeMARKETORDERREJECT captures enum value "MARKET_ORDER_REJECT"
	TakeProfitOrderTransactionTypeMARKETORDERREJECT string = "MARKET_ORDER_REJECT"
	// TakeProfitOrderTransactionTypeLIMITORDER captures enum value "LIMIT_ORDER"
	TakeProfitOrderTransactionTypeLIMITORDER string = "LIMIT_ORDER"
	// TakeProfitOrderTransactionTypeLIMITORDERREJECT captures enum value "LIMIT_ORDER_REJECT"
	TakeProfitOrderTransactionTypeLIMITORDERREJECT string = "LIMIT_ORDER_REJECT"
	// TakeProfitOrderTransactionTypeSTOPORDER captures enum value "STOP_ORDER"
	TakeProfitOrderTransactionTypeSTOPORDER string = "STOP_ORDER"
	// TakeProfitOrderTransactionTypeSTOPORDERREJECT captures enum value "STOP_ORDER_REJECT"
	TakeProfitOrderTransactionTypeSTOPORDERREJECT string = "STOP_ORDER_REJECT"
	// TakeProfitOrderTransactionTypeMARKETIFTOUCHEDORDER captures enum value "MARKET_IF_TOUCHED_ORDER"
	TakeProfitOrderTransactionTypeMARKETIFTOUCHEDORDER string = "MARKET_IF_TOUCHED_ORDER"
	// TakeProfitOrderTransactionTypeMARKETIFTOUCHEDORDERREJECT captures enum value "MARKET_IF_TOUCHED_ORDER_REJECT"
	TakeProfitOrderTransactionTypeMARKETIFTOUCHEDORDERREJECT string = "MARKET_IF_TOUCHED_ORDER_REJECT"
	// TakeProfitOrderTransactionTypeTAKEPROFITORDER captures enum value "TAKE_PROFIT_ORDER"
	TakeProfitOrderTransactionTypeTAKEPROFITORDER string = "TAKE_PROFIT_ORDER"
	// TakeProfitOrderTransactionTypeTAKEPROFITORDERREJECT captures enum value "TAKE_PROFIT_ORDER_REJECT"
	TakeProfitOrderTransactionTypeTAKEPROFITORDERREJECT string = "TAKE_PROFIT_ORDER_REJECT"
	// TakeProfitOrderTransactionTypeSTOPLOSSORDER captures enum value "STOP_LOSS_ORDER"
	TakeProfitOrderTransactionTypeSTOPLOSSORDER string = "STOP_LOSS_ORDER"
	// TakeProfitOrderTransactionTypeSTOPLOSSORDERREJECT captures enum value "STOP_LOSS_ORDER_REJECT"
	TakeProfitOrderTransactionTypeSTOPLOSSORDERREJECT string = "STOP_LOSS_ORDER_REJECT"
	// TakeProfitOrderTransactionTypeTRAILINGSTOPLOSSORDER captures enum value "TRAILING_STOP_LOSS_ORDER"
	TakeProfitOrderTransactionTypeTRAILINGSTOPLOSSORDER string = "TRAILING_STOP_LOSS_ORDER"
	// TakeProfitOrderTransactionTypeTRAILINGSTOPLOSSORDERREJECT captures enum value "TRAILING_STOP_LOSS_ORDER_REJECT"
	TakeProfitOrderTransactionTypeTRAILINGSTOPLOSSORDERREJECT string = "TRAILING_STOP_LOSS_ORDER_REJECT"
	// TakeProfitOrderTransactionTypeORDERFILL captures enum value "ORDER_FILL"
	TakeProfitOrderTransactionTypeORDERFILL string = "ORDER_FILL"
	// TakeProfitOrderTransactionTypeORDERCANCEL captures enum value "ORDER_CANCEL"
	TakeProfitOrderTransactionTypeORDERCANCEL string = "ORDER_CANCEL"
	// TakeProfitOrderTransactionTypeORDERCANCELREJECT captures enum value "ORDER_CANCEL_REJECT"
	TakeProfitOrderTransactionTypeORDERCANCELREJECT string = "ORDER_CANCEL_REJECT"
	// TakeProfitOrderTransactionTypeORDERCLIENTEXTENSIONSMODIFY captures enum value "ORDER_CLIENT_EXTENSIONS_MODIFY"
	TakeProfitOrderTransactionTypeORDERCLIENTEXTENSIONSMODIFY string = "ORDER_CLIENT_EXTENSIONS_MODIFY"
	// TakeProfitOrderTransactionTypeORDERCLIENTEXTENSIONSMODIFYREJECT captures enum value "ORDER_CLIENT_EXTENSIONS_MODIFY_REJECT"
	TakeProfitOrderTransactionTypeORDERCLIENTEXTENSIONSMODIFYREJECT string = "ORDER_CLIENT_EXTENSIONS_MODIFY_REJECT"
	// TakeProfitOrderTransactionTypeTRADECLIENTEXTENSIONSMODIFY captures enum value "TRADE_CLIENT_EXTENSIONS_MODIFY"
	TakeProfitOrderTransactionTypeTRADECLIENTEXTENSIONSMODIFY string = "TRADE_CLIENT_EXTENSIONS_MODIFY"
	// TakeProfitOrderTransactionTypeTRADECLIENTEXTENSIONSMODIFYREJECT captures enum value "TRADE_CLIENT_EXTENSIONS_MODIFY_REJECT"
	TakeProfitOrderTransactionTypeTRADECLIENTEXTENSIONSMODIFYREJECT string = "TRADE_CLIENT_EXTENSIONS_MODIFY_REJECT"
	// TakeProfitOrderTransactionTypeMARGINCALLENTER captures enum value "MARGIN_CALL_ENTER"
	TakeProfitOrderTransactionTypeMARGINCALLENTER string = "MARGIN_CALL_ENTER"
	// TakeProfitOrderTransactionTypeMARGINCALLEXTEND captures enum value "MARGIN_CALL_EXTEND"
	TakeProfitOrderTransactionTypeMARGINCALLEXTEND string = "MARGIN_CALL_EXTEND"
	// TakeProfitOrderTransactionTypeMARGINCALLEXIT captures enum value "MARGIN_CALL_EXIT"
	TakeProfitOrderTransactionTypeMARGINCALLEXIT string = "MARGIN_CALL_EXIT"
	// TakeProfitOrderTransactionTypeDELAYEDTRADECLOSURE captures enum value "DELAYED_TRADE_CLOSURE"
	TakeProfitOrderTransactionTypeDELAYEDTRADECLOSURE string = "DELAYED_TRADE_CLOSURE"
	// TakeProfitOrderTransactionTypeDAILYFINANCING captures enum value "DAILY_FINANCING"
	TakeProfitOrderTransactionTypeDAILYFINANCING string = "DAILY_FINANCING"
	// TakeProfitOrderTransactionTypeRESETRESETTABLEPL captures enum value "RESET_RESETTABLE_PL"
	TakeProfitOrderTransactionTypeRESETRESETTABLEPL string = "RESET_RESETTABLE_PL"
)

// prop value enum
func (m *TakeProfitOrderTransaction) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, takeProfitOrderTransactionTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *TakeProfitOrderTransaction) validateType(formats strfmt.Registry) error {

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
func (m *TakeProfitOrderTransaction) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TakeProfitOrderTransaction) UnmarshalBinary(b []byte) error {
	var res TakeProfitOrderTransaction
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}