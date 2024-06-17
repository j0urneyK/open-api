// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ExcludedFunctionRoute excluded function route
//
// swagger:model excludedFunctionRoute
type ExcludedFunctionRoute struct {

	// expression
	Expression string `json:"expression,omitempty"`

	// literal
	Literal string `json:"literal,omitempty"`

	// pattern
	Pattern string `json:"pattern,omitempty"`
}

// Validate validates this excluded function route
func (m *ExcludedFunctionRoute) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ExcludedFunctionRoute) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ExcludedFunctionRoute) UnmarshalBinary(b []byte) error {
	var res ExcludedFunctionRoute
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
