// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PVMInstanceFault Fault information (if occurred)
//
// swagger:model PVMInstanceFault
type PVMInstanceFault struct {

	// The fault status of the server, if any
	Code float64 `json:"code,omitempty"`

	// The date and time the fault occurred
	// Format: date-time
	Created strfmt.DateTime `json:"created,omitempty"`

	// The fault details of the server, if any
	Details string `json:"details,omitempty"`

	// The fault message of the server, if any
	Message string `json:"message,omitempty"`
}

// Validate validates this p VM instance fault
func (m *PVMInstanceFault) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreated(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PVMInstanceFault) validateCreated(formats strfmt.Registry) error {
	if swag.IsZero(m.Created) { // not required
		return nil
	}

	if err := validate.FormatOf("created", "body", "date-time", m.Created.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this p VM instance fault based on context it is used
func (m *PVMInstanceFault) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PVMInstanceFault) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PVMInstanceFault) UnmarshalBinary(b []byte) error {
	var res PVMInstanceFault
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
