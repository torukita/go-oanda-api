// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// GetPositionOKBody get position o k body
// swagger:model getPositionOKBody
type GetPositionOKBody struct {

	// The ID of the most recent Transaction created for the Account
	LastTransactionID string `json:"lastTransactionID,omitempty"`

	// position
	Position *Position `json:"position,omitempty"`
}

// Validate validates this get position o k body
func (m *GetPositionOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePosition(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetPositionOKBody) validatePosition(formats strfmt.Registry) error {

	if swag.IsZero(m.Position) { // not required
		return nil
	}

	if m.Position != nil {

		if err := m.Position.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("position")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GetPositionOKBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetPositionOKBody) UnmarshalBinary(b []byte) error {
	var res GetPositionOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
