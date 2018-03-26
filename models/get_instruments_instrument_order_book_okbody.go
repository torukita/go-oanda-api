// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// GetInstrumentsInstrumentOrderBookOKBody get instruments instrument order book o k body
// swagger:model getInstrumentsInstrumentOrderBookOKBody
type GetInstrumentsInstrumentOrderBookOKBody struct {

	// order book
	OrderBook *OrderBook `json:"orderBook,omitempty"`
}

// Validate validates this get instruments instrument order book o k body
func (m *GetInstrumentsInstrumentOrderBookOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOrderBook(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetInstrumentsInstrumentOrderBookOKBody) validateOrderBook(formats strfmt.Registry) error {

	if swag.IsZero(m.OrderBook) { // not required
		return nil
	}

	if m.OrderBook != nil {

		if err := m.OrderBook.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("orderBook")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GetInstrumentsInstrumentOrderBookOKBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetInstrumentsInstrumentOrderBookOKBody) UnmarshalBinary(b []byte) error {
	var res GetInstrumentsInstrumentOrderBookOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
