// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// SetOrderClientExtensionsParamsBody set order client extensions params body
// swagger:model setOrderClientExtensionsParamsBody
type SetOrderClientExtensionsParamsBody struct {

	// client extensions
	ClientExtensions *ClientExtensions `json:"clientExtensions,omitempty"`

	// trade client extensions
	TradeClientExtensions *ClientExtensions `json:"tradeClientExtensions,omitempty"`
}

// Validate validates this set order client extensions params body
func (m *SetOrderClientExtensionsParamsBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateClientExtensions(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateTradeClientExtensions(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SetOrderClientExtensionsParamsBody) validateClientExtensions(formats strfmt.Registry) error {

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

func (m *SetOrderClientExtensionsParamsBody) validateTradeClientExtensions(formats strfmt.Registry) error {

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

// MarshalBinary interface implementation
func (m *SetOrderClientExtensionsParamsBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SetOrderClientExtensionsParamsBody) UnmarshalBinary(b []byte) error {
	var res SetOrderClientExtensionsParamsBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}