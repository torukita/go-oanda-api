// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// ClosePositionParamsBody close position params body
// swagger:model closePositionParamsBody
type ClosePositionParamsBody struct {

	// long client extensions
	LongClientExtensions *ClientExtensions `json:"longClientExtensions,omitempty"`

	// Indication of how much of the long Position to closeout. Either the string "ALL", the string "NONE", or a DecimalNumber representing how many units of the long position to close using a PositionCloseout MarketOrder. The units specified must always be positive.
	LongUnits string `json:"longUnits,omitempty"`

	// short client extensions
	ShortClientExtensions *ClientExtensions `json:"shortClientExtensions,omitempty"`

	// Indication of how much of the short Position to closeout. Either the string "ALL", the string "NONE", or a DecimalNumber representing how many units of the short position to close using a PositionCloseout MarketOrder. The units specified must always be positive.
	ShortUnits string `json:"shortUnits,omitempty"`
}

// Validate validates this close position params body
func (m *ClosePositionParamsBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLongClientExtensions(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateShortClientExtensions(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClosePositionParamsBody) validateLongClientExtensions(formats strfmt.Registry) error {

	if swag.IsZero(m.LongClientExtensions) { // not required
		return nil
	}

	if m.LongClientExtensions != nil {

		if err := m.LongClientExtensions.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("longClientExtensions")
			}
			return err
		}
	}

	return nil
}

func (m *ClosePositionParamsBody) validateShortClientExtensions(formats strfmt.Registry) error {

	if swag.IsZero(m.ShortClientExtensions) { // not required
		return nil
	}

	if m.ShortClientExtensions != nil {

		if err := m.ShortClientExtensions.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("shortClientExtensions")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ClosePositionParamsBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClosePositionParamsBody) UnmarshalBinary(b []byte) error {
	var res ClosePositionParamsBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
