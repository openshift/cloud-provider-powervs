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

// SnapshotCreate snapshot create
//
// swagger:model SnapshotCreate
type SnapshotCreate struct {

	// Description of the PVM instance snapshot
	Description string `json:"description,omitempty"`

	// Name of the PVM instance snapshot to create
	// Required: true
	Name *string `json:"name"`

	// Enables optimized performance path for creating snapshots of PVM instances.
	PerformancePath *bool `json:"performancePath,omitempty"`

	// List of volumes to include in the PVM instance snapshot
	VolumeIDs []string `json:"volumeIDs"`
}

// Validate validates this snapshot create
func (m *SnapshotCreate) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SnapshotCreate) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this snapshot create based on context it is used
func (m *SnapshotCreate) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SnapshotCreate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SnapshotCreate) UnmarshalBinary(b []byte) error {
	var res SnapshotCreate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}