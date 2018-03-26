// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// PositionBookBucket The position book data for a partition of the instrument's prices.
// swagger:model PositionBookBucket
type PositionBookBucket struct {

	// The percentage of the total number of positions represented by the long positions found in this bucket.
	LongCountPercent string `json:"longCountPercent,omitempty"`

	// The lowest price (inclusive) covered by the bucket. The bucket covers the price range from the price to price + the position book's bucketWidth.
	Price string `json:"price,omitempty"`

	// The percentage of the total number of positions represented by the short positions found in this bucket.
	ShortCountPercent string `json:"shortCountPercent,omitempty"`
}

// Validate validates this position book bucket
func (m *PositionBookBucket) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *PositionBookBucket) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PositionBookBucket) UnmarshalBinary(b []byte) error {
	var res PositionBookBucket
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}