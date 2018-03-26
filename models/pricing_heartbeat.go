// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// PricingHeartbeat A PricingHeartbeat object is injected into the Pricing stream to ensure that the HTTP connection remains active.
// swagger:model PricingHeartbeat
type PricingHeartbeat struct {

	// The date/time when the Heartbeat was created.
	Time string `json:"time,omitempty"`

	// The string "HEARTBEAT"
	Type string `json:"type,omitempty"`
}

// Validate validates this pricing heartbeat
func (m *PricingHeartbeat) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *PricingHeartbeat) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PricingHeartbeat) UnmarshalBinary(b []byte) error {
	var res PricingHeartbeat
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}