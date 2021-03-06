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

// ServiceInstanceRequest service instance request
//
// swagger:model ServiceInstanceRequest
type ServiceInstanceRequest struct {

	// Indicates the current state of the service instance.
	// Required: true
	Enabled *bool `json:"enabled"`

	// Optional string stating the reason code for the service instance state change. Valid values are BMX_ACCT_ACTIVATE, BMX_SERVICE_INSTANCE_BELOW_CAP for enable calls, and BMX_ACCT_SUSPEND, BMX_SERVICE_INSTANCE_ABOVE_CAP for disable calls.
	InitiatorID string `json:"initiator_id,omitempty"`

	// Optional string showing the user id initiating the call
	ReasonCode string `json:"reason_code,omitempty"`
}

// Validate validates this service instance request
func (m *ServiceInstanceRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEnabled(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ServiceInstanceRequest) validateEnabled(formats strfmt.Registry) error {

	if err := validate.Required("enabled", "body", m.Enabled); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this service instance request based on context it is used
func (m *ServiceInstanceRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ServiceInstanceRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServiceInstanceRequest) UnmarshalBinary(b []byte) error {
	var res ServiceInstanceRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
