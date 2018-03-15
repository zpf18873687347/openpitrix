// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// ProtobufUint64Value Wrapper message for `uint64`.
//
// The JSON representation for `UInt64Value` is JSON string.
// swagger:model protobufUInt64Value
type ProtobufUint64Value struct {

	// The uint64 value.
	Value uint64 `json:"value,omitempty"`
}

// Validate validates this protobuf uint64 value
func (m *ProtobufUint64Value) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *ProtobufUint64Value) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ProtobufUint64Value) UnmarshalBinary(b []byte) error {
	var res ProtobufUint64Value
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}