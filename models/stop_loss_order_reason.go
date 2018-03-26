// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// StopLossOrderReason The reason that the Stop Loss Order was initiated
// swagger:model StopLossOrderReason
type StopLossOrderReason string

const (
	// StopLossOrderReasonCLIENTORDER captures enum value "CLIENT_ORDER"
	StopLossOrderReasonCLIENTORDER StopLossOrderReason = "CLIENT_ORDER"
	// StopLossOrderReasonREPLACEMENT captures enum value "REPLACEMENT"
	StopLossOrderReasonREPLACEMENT StopLossOrderReason = "REPLACEMENT"
	// StopLossOrderReasonONFILL captures enum value "ON_FILL"
	StopLossOrderReasonONFILL StopLossOrderReason = "ON_FILL"
)

// for schema
var stopLossOrderReasonEnum []interface{}

func init() {
	var res []StopLossOrderReason
	if err := json.Unmarshal([]byte(`["CLIENT_ORDER","REPLACEMENT","ON_FILL"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		stopLossOrderReasonEnum = append(stopLossOrderReasonEnum, v)
	}
}

func (m StopLossOrderReason) validateStopLossOrderReasonEnum(path, location string, value StopLossOrderReason) error {
	if err := validate.Enum(path, location, value, stopLossOrderReasonEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this stop loss order reason
func (m StopLossOrderReason) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateStopLossOrderReasonEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
