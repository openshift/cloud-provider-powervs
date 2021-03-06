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

// Task task
//
// swagger:model Task
type Task struct {

	// Cloud Instance ID of task owner
	// Required: true
	CloudInstanceID *string `json:"cloudInstanceID"`

	// the component id of the task
	// Required: true
	ComponentID *string `json:"componentID"`

	// the component type of the task
	// Required: true
	ComponentType *string `json:"componentType"`

	// Creation Date
	// Required: true
	// Format: date-time
	CreationDate *strfmt.DateTime `json:"creationDate"`

	// Last Update Date
	// Required: true
	// Format: date-time
	LastUpdateDate *strfmt.DateTime `json:"lastUpdateDate"`

	// Task Operation
	// Required: true
	Operation *string `json:"operation"`

	// status code of the task
	// Required: true
	Status *string `json:"status"`

	// status detail of the task
	// Required: true
	StatusDetail *string `json:"statusDetail"`

	// Pcloud Task ID
	// Required: true
	TaskID *string `json:"taskID"`
}

// Validate validates this task
func (m *Task) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCloudInstanceID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateComponentID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateComponentType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreationDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastUpdateDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOperation(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatusDetail(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTaskID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Task) validateCloudInstanceID(formats strfmt.Registry) error {

	if err := validate.Required("cloudInstanceID", "body", m.CloudInstanceID); err != nil {
		return err
	}

	return nil
}

func (m *Task) validateComponentID(formats strfmt.Registry) error {

	if err := validate.Required("componentID", "body", m.ComponentID); err != nil {
		return err
	}

	return nil
}

func (m *Task) validateComponentType(formats strfmt.Registry) error {

	if err := validate.Required("componentType", "body", m.ComponentType); err != nil {
		return err
	}

	return nil
}

func (m *Task) validateCreationDate(formats strfmt.Registry) error {

	if err := validate.Required("creationDate", "body", m.CreationDate); err != nil {
		return err
	}

	if err := validate.FormatOf("creationDate", "body", "date-time", m.CreationDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Task) validateLastUpdateDate(formats strfmt.Registry) error {

	if err := validate.Required("lastUpdateDate", "body", m.LastUpdateDate); err != nil {
		return err
	}

	if err := validate.FormatOf("lastUpdateDate", "body", "date-time", m.LastUpdateDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Task) validateOperation(formats strfmt.Registry) error {

	if err := validate.Required("operation", "body", m.Operation); err != nil {
		return err
	}

	return nil
}

func (m *Task) validateStatus(formats strfmt.Registry) error {

	if err := validate.Required("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

func (m *Task) validateStatusDetail(formats strfmt.Registry) error {

	if err := validate.Required("statusDetail", "body", m.StatusDetail); err != nil {
		return err
	}

	return nil
}

func (m *Task) validateTaskID(formats strfmt.Registry) error {

	if err := validate.Required("taskID", "body", m.TaskID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this task based on context it is used
func (m *Task) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Task) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Task) UnmarshalBinary(b []byte) error {
	var res Task
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
