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

// Order The base Order definition specifies the properties that are common to all Orders.
// swagger:model Order
type Order struct {

	// client extensions
	ClientExtensions *ClientExtensions `json:"clientExtensions,omitempty"`

	// The time when the Order was created.
	CreateTime string `json:"createTime,omitempty"`

	// The Order's identifier, unique within the Order's Account.
	ID string `json:"id,omitempty"`

	// The current state of the Order.
	State string `json:"state,omitempty"`
}

// Validate validates this order
func (m *Order) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateClientExtensions(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateState(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Order) validateClientExtensions(formats strfmt.Registry) error {

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

var orderTypeStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["PENDING","FILLED","TRIGGERED","CANCELLED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		orderTypeStatePropEnum = append(orderTypeStatePropEnum, v)
	}
}

const (
	// OrderStatePENDING captures enum value "PENDING"
	OrderStatePENDING string = "PENDING"
	// OrderStateFILLED captures enum value "FILLED"
	OrderStateFILLED string = "FILLED"
	// OrderStateTRIGGERED captures enum value "TRIGGERED"
	OrderStateTRIGGERED string = "TRIGGERED"
	// OrderStateCANCELLED captures enum value "CANCELLED"
	OrderStateCANCELLED string = "CANCELLED"
)

// prop value enum
func (m *Order) validateStateEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, orderTypeStatePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Order) validateState(formats strfmt.Registry) error {

	if swag.IsZero(m.State) { // not required
		return nil
	}

	// value enum
	if err := m.validateStateEnum("state", "body", m.State); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Order) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Order) UnmarshalBinary(b []byte) error {
	var res Order
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}