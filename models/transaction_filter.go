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

// TransactionFilter A filter that can be used when fetching Transactions
// swagger:model TransactionFilter
type TransactionFilter string

const (
	// TransactionFilterORDER captures enum value "ORDER"
	TransactionFilterORDER TransactionFilter = "ORDER"
	// TransactionFilterFUNDING captures enum value "FUNDING"
	TransactionFilterFUNDING TransactionFilter = "FUNDING"
	// TransactionFilterADMIN captures enum value "ADMIN"
	TransactionFilterADMIN TransactionFilter = "ADMIN"
	// TransactionFilterCREATE captures enum value "CREATE"
	TransactionFilterCREATE TransactionFilter = "CREATE"
	// TransactionFilterCLOSE captures enum value "CLOSE"
	TransactionFilterCLOSE TransactionFilter = "CLOSE"
	// TransactionFilterREOPEN captures enum value "REOPEN"
	TransactionFilterREOPEN TransactionFilter = "REOPEN"
	// TransactionFilterCLIENTCONFIGURE captures enum value "CLIENT_CONFIGURE"
	TransactionFilterCLIENTCONFIGURE TransactionFilter = "CLIENT_CONFIGURE"
	// TransactionFilterCLIENTCONFIGUREREJECT captures enum value "CLIENT_CONFIGURE_REJECT"
	TransactionFilterCLIENTCONFIGUREREJECT TransactionFilter = "CLIENT_CONFIGURE_REJECT"
	// TransactionFilterTRANSFERFUNDS captures enum value "TRANSFER_FUNDS"
	TransactionFilterTRANSFERFUNDS TransactionFilter = "TRANSFER_FUNDS"
	// TransactionFilterTRANSFERFUNDSREJECT captures enum value "TRANSFER_FUNDS_REJECT"
	TransactionFilterTRANSFERFUNDSREJECT TransactionFilter = "TRANSFER_FUNDS_REJECT"
	// TransactionFilterMARKETORDER captures enum value "MARKET_ORDER"
	TransactionFilterMARKETORDER TransactionFilter = "MARKET_ORDER"
	// TransactionFilterMARKETORDERREJECT captures enum value "MARKET_ORDER_REJECT"
	TransactionFilterMARKETORDERREJECT TransactionFilter = "MARKET_ORDER_REJECT"
	// TransactionFilterLIMITORDER captures enum value "LIMIT_ORDER"
	TransactionFilterLIMITORDER TransactionFilter = "LIMIT_ORDER"
	// TransactionFilterLIMITORDERREJECT captures enum value "LIMIT_ORDER_REJECT"
	TransactionFilterLIMITORDERREJECT TransactionFilter = "LIMIT_ORDER_REJECT"
	// TransactionFilterSTOPORDER captures enum value "STOP_ORDER"
	TransactionFilterSTOPORDER TransactionFilter = "STOP_ORDER"
	// TransactionFilterSTOPORDERREJECT captures enum value "STOP_ORDER_REJECT"
	TransactionFilterSTOPORDERREJECT TransactionFilter = "STOP_ORDER_REJECT"
	// TransactionFilterMARKETIFTOUCHEDORDER captures enum value "MARKET_IF_TOUCHED_ORDER"
	TransactionFilterMARKETIFTOUCHEDORDER TransactionFilter = "MARKET_IF_TOUCHED_ORDER"
	// TransactionFilterMARKETIFTOUCHEDORDERREJECT captures enum value "MARKET_IF_TOUCHED_ORDER_REJECT"
	TransactionFilterMARKETIFTOUCHEDORDERREJECT TransactionFilter = "MARKET_IF_TOUCHED_ORDER_REJECT"
	// TransactionFilterTAKEPROFITORDER captures enum value "TAKE_PROFIT_ORDER"
	TransactionFilterTAKEPROFITORDER TransactionFilter = "TAKE_PROFIT_ORDER"
	// TransactionFilterTAKEPROFITORDERREJECT captures enum value "TAKE_PROFIT_ORDER_REJECT"
	TransactionFilterTAKEPROFITORDERREJECT TransactionFilter = "TAKE_PROFIT_ORDER_REJECT"
	// TransactionFilterSTOPLOSSORDER captures enum value "STOP_LOSS_ORDER"
	TransactionFilterSTOPLOSSORDER TransactionFilter = "STOP_LOSS_ORDER"
	// TransactionFilterSTOPLOSSORDERREJECT captures enum value "STOP_LOSS_ORDER_REJECT"
	TransactionFilterSTOPLOSSORDERREJECT TransactionFilter = "STOP_LOSS_ORDER_REJECT"
	// TransactionFilterTRAILINGSTOPLOSSORDER captures enum value "TRAILING_STOP_LOSS_ORDER"
	TransactionFilterTRAILINGSTOPLOSSORDER TransactionFilter = "TRAILING_STOP_LOSS_ORDER"
	// TransactionFilterTRAILINGSTOPLOSSORDERREJECT captures enum value "TRAILING_STOP_LOSS_ORDER_REJECT"
	TransactionFilterTRAILINGSTOPLOSSORDERREJECT TransactionFilter = "TRAILING_STOP_LOSS_ORDER_REJECT"
	// TransactionFilterONECANCELSALLORDER captures enum value "ONE_CANCELS_ALL_ORDER"
	TransactionFilterONECANCELSALLORDER TransactionFilter = "ONE_CANCELS_ALL_ORDER"
	// TransactionFilterONECANCELSALLORDERREJECT captures enum value "ONE_CANCELS_ALL_ORDER_REJECT"
	TransactionFilterONECANCELSALLORDERREJECT TransactionFilter = "ONE_CANCELS_ALL_ORDER_REJECT"
	// TransactionFilterONECANCELSALLORDERTRIGGERED captures enum value "ONE_CANCELS_ALL_ORDER_TRIGGERED"
	TransactionFilterONECANCELSALLORDERTRIGGERED TransactionFilter = "ONE_CANCELS_ALL_ORDER_TRIGGERED"
	// TransactionFilterORDERFILL captures enum value "ORDER_FILL"
	TransactionFilterORDERFILL TransactionFilter = "ORDER_FILL"
	// TransactionFilterORDERCANCEL captures enum value "ORDER_CANCEL"
	TransactionFilterORDERCANCEL TransactionFilter = "ORDER_CANCEL"
	// TransactionFilterORDERCANCELREJECT captures enum value "ORDER_CANCEL_REJECT"
	TransactionFilterORDERCANCELREJECT TransactionFilter = "ORDER_CANCEL_REJECT"
	// TransactionFilterORDERCLIENTEXTENSIONSMODIFY captures enum value "ORDER_CLIENT_EXTENSIONS_MODIFY"
	TransactionFilterORDERCLIENTEXTENSIONSMODIFY TransactionFilter = "ORDER_CLIENT_EXTENSIONS_MODIFY"
	// TransactionFilterORDERCLIENTEXTENSIONSMODIFYREJECT captures enum value "ORDER_CLIENT_EXTENSIONS_MODIFY_REJECT"
	TransactionFilterORDERCLIENTEXTENSIONSMODIFYREJECT TransactionFilter = "ORDER_CLIENT_EXTENSIONS_MODIFY_REJECT"
	// TransactionFilterTRADECLIENTEXTENSIONSMODIFY captures enum value "TRADE_CLIENT_EXTENSIONS_MODIFY"
	TransactionFilterTRADECLIENTEXTENSIONSMODIFY TransactionFilter = "TRADE_CLIENT_EXTENSIONS_MODIFY"
	// TransactionFilterTRADECLIENTEXTENSIONSMODIFYREJECT captures enum value "TRADE_CLIENT_EXTENSIONS_MODIFY_REJECT"
	TransactionFilterTRADECLIENTEXTENSIONSMODIFYREJECT TransactionFilter = "TRADE_CLIENT_EXTENSIONS_MODIFY_REJECT"
	// TransactionFilterMARGINCALLENTER captures enum value "MARGIN_CALL_ENTER"
	TransactionFilterMARGINCALLENTER TransactionFilter = "MARGIN_CALL_ENTER"
	// TransactionFilterMARGINCALLEXTEND captures enum value "MARGIN_CALL_EXTEND"
	TransactionFilterMARGINCALLEXTEND TransactionFilter = "MARGIN_CALL_EXTEND"
	// TransactionFilterMARGINCALLEXIT captures enum value "MARGIN_CALL_EXIT"
	TransactionFilterMARGINCALLEXIT TransactionFilter = "MARGIN_CALL_EXIT"
	// TransactionFilterDELAYEDTRADECLOSURE captures enum value "DELAYED_TRADE_CLOSURE"
	TransactionFilterDELAYEDTRADECLOSURE TransactionFilter = "DELAYED_TRADE_CLOSURE"
	// TransactionFilterDAILYFINANCING captures enum value "DAILY_FINANCING"
	TransactionFilterDAILYFINANCING TransactionFilter = "DAILY_FINANCING"
	// TransactionFilterRESETRESETTABLEPL captures enum value "RESET_RESETTABLE_PL"
	TransactionFilterRESETRESETTABLEPL TransactionFilter = "RESET_RESETTABLE_PL"
)

// for schema
var transactionFilterEnum []interface{}

func init() {
	var res []TransactionFilter
	if err := json.Unmarshal([]byte(`["ORDER","FUNDING","ADMIN","CREATE","CLOSE","REOPEN","CLIENT_CONFIGURE","CLIENT_CONFIGURE_REJECT","TRANSFER_FUNDS","TRANSFER_FUNDS_REJECT","MARKET_ORDER","MARKET_ORDER_REJECT","LIMIT_ORDER","LIMIT_ORDER_REJECT","STOP_ORDER","STOP_ORDER_REJECT","MARKET_IF_TOUCHED_ORDER","MARKET_IF_TOUCHED_ORDER_REJECT","TAKE_PROFIT_ORDER","TAKE_PROFIT_ORDER_REJECT","STOP_LOSS_ORDER","STOP_LOSS_ORDER_REJECT","TRAILING_STOP_LOSS_ORDER","TRAILING_STOP_LOSS_ORDER_REJECT","ONE_CANCELS_ALL_ORDER","ONE_CANCELS_ALL_ORDER_REJECT","ONE_CANCELS_ALL_ORDER_TRIGGERED","ORDER_FILL","ORDER_CANCEL","ORDER_CANCEL_REJECT","ORDER_CLIENT_EXTENSIONS_MODIFY","ORDER_CLIENT_EXTENSIONS_MODIFY_REJECT","TRADE_CLIENT_EXTENSIONS_MODIFY","TRADE_CLIENT_EXTENSIONS_MODIFY_REJECT","MARGIN_CALL_ENTER","MARGIN_CALL_EXTEND","MARGIN_CALL_EXIT","DELAYED_TRADE_CLOSURE","DAILY_FINANCING","RESET_RESETTABLE_PL"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		transactionFilterEnum = append(transactionFilterEnum, v)
	}
}

func (m TransactionFilter) validateTransactionFilterEnum(path, location string, value TransactionFilter) error {
	if err := validate.Enum(path, location, value, transactionFilterEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this transaction filter
func (m TransactionFilter) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateTransactionFilterEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
