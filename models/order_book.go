// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// OrderBook The representation of an instrument's order book at a point in time
// swagger:model OrderBook
type OrderBook struct {

	// The price width for each bucket. Each bucket covers the price range from the bucket's price to the bucket's price + bucketWidth.
	BucketWidth string `json:"bucketWidth,omitempty"`

	// buckets
	Buckets OrderBookBuckets `json:"buckets"`

	// The order book's instrument
	Instrument string `json:"instrument,omitempty"`

	// The price (midpoint) for the order book's instrument at the time of the order book snapshot
	Price string `json:"price,omitempty"`

	// The time when the order book snapshot was created.
	Time string `json:"time,omitempty"`
}

// Validate validates this order book
func (m *OrderBook) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *OrderBook) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OrderBook) UnmarshalBinary(b []byte) error {
	var res OrderBook
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}