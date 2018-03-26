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

// WeeklyAlignment The day of the week to use for candlestick granularities with weekly alignment.
// swagger:model WeeklyAlignment
type WeeklyAlignment string

const (
	// WeeklyAlignmentMonday captures enum value "Monday"
	WeeklyAlignmentMonday WeeklyAlignment = "Monday"
	// WeeklyAlignmentTuesday captures enum value "Tuesday"
	WeeklyAlignmentTuesday WeeklyAlignment = "Tuesday"
	// WeeklyAlignmentWednesday captures enum value "Wednesday"
	WeeklyAlignmentWednesday WeeklyAlignment = "Wednesday"
	// WeeklyAlignmentThursday captures enum value "Thursday"
	WeeklyAlignmentThursday WeeklyAlignment = "Thursday"
	// WeeklyAlignmentFriday captures enum value "Friday"
	WeeklyAlignmentFriday WeeklyAlignment = "Friday"
	// WeeklyAlignmentSaturday captures enum value "Saturday"
	WeeklyAlignmentSaturday WeeklyAlignment = "Saturday"
	// WeeklyAlignmentSunday captures enum value "Sunday"
	WeeklyAlignmentSunday WeeklyAlignment = "Sunday"
)

// for schema
var weeklyAlignmentEnum []interface{}

func init() {
	var res []WeeklyAlignment
	if err := json.Unmarshal([]byte(`["Monday","Tuesday","Wednesday","Thursday","Friday","Saturday","Sunday"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		weeklyAlignmentEnum = append(weeklyAlignmentEnum, v)
	}
}

func (m WeeklyAlignment) validateWeeklyAlignmentEnum(path, location string, value WeeklyAlignment) error {
	if err := validate.Enum(path, location, value, weeklyAlignmentEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this weekly alignment
func (m WeeklyAlignment) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateWeeklyAlignmentEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
