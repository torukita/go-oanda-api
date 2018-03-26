// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// InstrumentCommission An InstrumentCommission represents an instrument-specific commission
// swagger:model InstrumentCommission
type InstrumentCommission struct {

	// The commission amount (in the Account's home currency) charged per unitsTraded of the instrument
	Commission string `json:"commission,omitempty"`

	// The name of the instrument
	Instrument string `json:"instrument,omitempty"`

	// The minimum commission amount (in the Account's home currency) that is charged when an Order is filled for this instrument.
	MinimumCommission string `json:"minimumCommission,omitempty"`

	// The number of units traded that the commission amount is based on.
	UnitsTraded string `json:"unitsTraded,omitempty"`
}

// Validate validates this instrument commission
func (m *InstrumentCommission) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *InstrumentCommission) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InstrumentCommission) UnmarshalBinary(b []byte) error {
	var res InstrumentCommission
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}