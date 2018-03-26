// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// SetTradeClientExtensionsMethodNotAllowedBody set trade client extensions method not allowed body
// swagger:model setTradeClientExtensionsMethodNotAllowedBody
type SetTradeClientExtensionsMethodNotAllowedBody struct {

	// The code of the error that has occurred. This field may not be returned for some errors.
	ErrorCode string `json:"errorCode,omitempty"`

	// The human-readable description of the error that has occurred.
	ErrorMessage string `json:"errorMessage,omitempty"`
}

// Validate validates this set trade client extensions method not allowed body
func (m *SetTradeClientExtensionsMethodNotAllowedBody) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *SetTradeClientExtensionsMethodNotAllowedBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SetTradeClientExtensionsMethodNotAllowedBody) UnmarshalBinary(b []byte) error {
	var res SetTradeClientExtensionsMethodNotAllowedBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
