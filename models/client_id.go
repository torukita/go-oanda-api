// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
)

// ClientID A client-provided identifier, used by clients to refer to their Orders or Trades with an identifier that they have provided.
// swagger:model ClientID
type ClientID string

// Validate validates this client ID
func (m ClientID) Validate(formats strfmt.Registry) error {
	return nil
}
