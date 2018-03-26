// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// GetUserInfoOKBody get user info o k body
// swagger:model getUserInfoOKBody
type GetUserInfoOKBody struct {

	// user info
	UserInfo *UserInfo `json:"userInfo,omitempty"`
}

// Validate validates this get user info o k body
func (m *GetUserInfoOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateUserInfo(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetUserInfoOKBody) validateUserInfo(formats strfmt.Registry) error {

	if swag.IsZero(m.UserInfo) { // not required
		return nil
	}

	if m.UserInfo != nil {

		if err := m.UserInfo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("userInfo")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GetUserInfoOKBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetUserInfoOKBody) UnmarshalBinary(b []byte) error {
	var res GetUserInfoOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}