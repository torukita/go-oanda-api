// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// Account The full details of a client's Account. This includes full open Trade, open Position and pending Order representation.
// swagger:model Account
type Account struct {

	// The net asset value of the Account. Equal to Account balance + unrealizedPL. Represented in the Account's home currency.
	NAV string `json:"NAV,omitempty"`

	// Client-assigned alias for the Account. Only provided if the Account has an alias set
	Alias string `json:"alias,omitempty"`

	// The current balance of the Account. Represented in the Account's home currency.
	Balance string `json:"balance,omitempty"`

	// The total amount of commission paid over the lifetime of the Account. Represented in the Account's home currency.
	Commission string `json:"commission,omitempty"`

	// ID of the user that created the Account.
	CreatedByUserID int64 `json:"createdByUserID,omitempty"`

	// The date/time when the Account was created.
	CreatedTime string `json:"createdTime,omitempty"`

	// The home currency of the Account
	Currency string `json:"currency,omitempty"`

	// Flag indicating that the Account has hedging enabled.
	HedgingEnabled bool `json:"hedgingEnabled,omitempty"`

	// The Account's identifier
	ID string `json:"id,omitempty"`

	// The date/time of the Account's last margin call extension.
	LastMarginCallExtensionTime string `json:"lastMarginCallExtensionTime,omitempty"`

	// The ID of the last Transaction created for the Account.
	LastTransactionID string `json:"lastTransactionID,omitempty"`

	// Margin available for Account. Represented in the Account's home currency.
	MarginAvailable string `json:"marginAvailable,omitempty"`

	// The date/time when the Account entered a margin call state. Only provided if the Account is in a margin call.
	MarginCallEnterTime string `json:"marginCallEnterTime,omitempty"`

	// The number of times that the Account's current margin call was extended.
	MarginCallExtensionCount int64 `json:"marginCallExtensionCount,omitempty"`

	// The Account's margin call margin used.
	MarginCallMarginUsed string `json:"marginCallMarginUsed,omitempty"`

	// The Account's margin call percentage. When this value is 1.0 or above the Account is in a margin call situation.
	MarginCallPercent string `json:"marginCallPercent,omitempty"`

	// The Account's margin closeout margin used.
	MarginCloseoutMarginUsed string `json:"marginCloseoutMarginUsed,omitempty"`

	// The Account's margin closeout NAV.
	MarginCloseoutNAV string `json:"marginCloseoutNAV,omitempty"`

	// The Account's margin closeout percentage. When this value is 1.0 or above the Account is in a margin closeout situation.
	MarginCloseoutPercent string `json:"marginCloseoutPercent,omitempty"`

	// The value of the Account's open positions as used for margin closeout calculations represented in the Account's home currency.
	MarginCloseoutPositionValue string `json:"marginCloseoutPositionValue,omitempty"`

	// The Account's margin closeout unrealized PL.
	MarginCloseoutUnrealizedPL string `json:"marginCloseoutUnrealizedPL,omitempty"`

	// Client-provided margin rate override for the Account. The effective margin rate of the Account is the lesser of this value and the OANDA margin rate for the Account's division. This value is only provided if a margin rate override exists for the Account.
	MarginRate string `json:"marginRate,omitempty"`

	// Margin currently used for the Account. Represented in the Account's home currency.
	MarginUsed string `json:"marginUsed,omitempty"`

	// The number of Positions currently open in the Account.
	OpenPositionCount int64 `json:"openPositionCount,omitempty"`

	// The number of Trades currently open in the Account.
	OpenTradeCount int64 `json:"openTradeCount,omitempty"`

	// orders
	Orders AccountOrders `json:"orders"`

	// The number of Orders currently pending in the Account.
	PendingOrderCount int64 `json:"pendingOrderCount,omitempty"`

	// The total profit/loss realized over the lifetime of the Account. Represented in the Account's home currency.
	Pl string `json:"pl,omitempty"`

	// The value of the Account's open positions represented in the Account's home currency.
	PositionValue string `json:"positionValue,omitempty"`

	// positions
	Positions AccountPositions `json:"positions"`

	// The total realized profit/loss for the Account since it was last reset by the client. Represented in the Account's home currency.
	ResettablePL string `json:"resettablePL,omitempty"`

	// The date/time that the Account's resettablePL was last reset.
	ResettabledPLTime string `json:"resettabledPLTime,omitempty"`

	// trades
	Trades AccountTrades `json:"trades"`

	// The total unrealized profit/loss for all Trades currently open in the Account. Represented in the Account's home currency.
	UnrealizedPL string `json:"unrealizedPL,omitempty"`

	// The current WithdrawalLimit for the account which will be zero or a positive value indicating how much can be withdrawn from the account.
	WithdrawalLimit string `json:"withdrawalLimit,omitempty"`
}

// Validate validates this account
func (m *Account) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *Account) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Account) UnmarshalBinary(b []byte) error {
	var res Account
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
