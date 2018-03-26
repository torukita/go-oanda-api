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

// StopLossDetails StopLossDetails specifies the details of a Stop Loss Order to be created on behalf of a client. This may happen when an Order is filled that opens a Trade requiring a Stop Loss, or when a Trade's dependent Stop Loss Order is modified directly through the Trade.
// swagger:model StopLossDetails
type StopLossDetails struct {

	// client extensions
	ClientExtensions *ClientExtensions `json:"clientExtensions,omitempty"`

	// The date when the Stop Loss Order will be cancelled on if timeInForce is GTD.
	GtdTime string `json:"gtdTime,omitempty"`

	// The price that the Stop Loss Order will be triggered at.
	Price string `json:"price,omitempty"`

	// The time in force for the created Stop Loss Order. This may only be GTC, GTD or GFD.
	TimeInForce string `json:"timeInForce,omitempty"`
}

// Validate validates this stop loss details
func (m *StopLossDetails) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateClientExtensions(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateTimeInForce(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *StopLossDetails) validateClientExtensions(formats strfmt.Registry) error {

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

var stopLossDetailsTypeTimeInForcePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["GTC","GTD","GFD","FOK","IOC"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		stopLossDetailsTypeTimeInForcePropEnum = append(stopLossDetailsTypeTimeInForcePropEnum, v)
	}
}

const (
	// StopLossDetailsTimeInForceGTC captures enum value "GTC"
	StopLossDetailsTimeInForceGTC string = "GTC"
	// StopLossDetailsTimeInForceGTD captures enum value "GTD"
	StopLossDetailsTimeInForceGTD string = "GTD"
	// StopLossDetailsTimeInForceGFD captures enum value "GFD"
	StopLossDetailsTimeInForceGFD string = "GFD"
	// StopLossDetailsTimeInForceFOK captures enum value "FOK"
	StopLossDetailsTimeInForceFOK string = "FOK"
	// StopLossDetailsTimeInForceIOC captures enum value "IOC"
	StopLossDetailsTimeInForceIOC string = "IOC"
)

// prop value enum
func (m *StopLossDetails) validateTimeInForceEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, stopLossDetailsTypeTimeInForcePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *StopLossDetails) validateTimeInForce(formats strfmt.Registry) error {

	if swag.IsZero(m.TimeInForce) { // not required
		return nil
	}

	// value enum
	if err := m.validateTimeInForceEnum("timeInForce", "body", m.TimeInForce); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *StopLossDetails) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StopLossDetails) UnmarshalBinary(b []byte) error {
	var res StopLossDetails
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}