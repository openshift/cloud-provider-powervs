// Code generated by go-swagger; DO NOT EDIT.

package p_cloud_p_vm_instances

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/IBM-Cloud/power-go-client/power/models"
)

// PcloudPvminstancesGetallReader is a Reader for the PcloudPvminstancesGetall structure.
type PcloudPvminstancesGetallReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PcloudPvminstancesGetallReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPcloudPvminstancesGetallOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPcloudPvminstancesGetallBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPcloudPvminstancesGetallUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPcloudPvminstancesGetallInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPcloudPvminstancesGetallOK creates a PcloudPvminstancesGetallOK with default headers values
func NewPcloudPvminstancesGetallOK() *PcloudPvminstancesGetallOK {
	return &PcloudPvminstancesGetallOK{}
}

/* PcloudPvminstancesGetallOK describes a response with status code 200, with default header values.

OK
*/
type PcloudPvminstancesGetallOK struct {
	Payload *models.PVMInstances
}

func (o *PcloudPvminstancesGetallOK) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances][%d] pcloudPvminstancesGetallOK  %+v", 200, o.Payload)
}
func (o *PcloudPvminstancesGetallOK) GetPayload() *models.PVMInstances {
	return o.Payload
}

func (o *PcloudPvminstancesGetallOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PVMInstances)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudPvminstancesGetallBadRequest creates a PcloudPvminstancesGetallBadRequest with default headers values
func NewPcloudPvminstancesGetallBadRequest() *PcloudPvminstancesGetallBadRequest {
	return &PcloudPvminstancesGetallBadRequest{}
}

/* PcloudPvminstancesGetallBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PcloudPvminstancesGetallBadRequest struct {
	Payload *models.Error
}

func (o *PcloudPvminstancesGetallBadRequest) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances][%d] pcloudPvminstancesGetallBadRequest  %+v", 400, o.Payload)
}
func (o *PcloudPvminstancesGetallBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudPvminstancesGetallBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudPvminstancesGetallUnauthorized creates a PcloudPvminstancesGetallUnauthorized with default headers values
func NewPcloudPvminstancesGetallUnauthorized() *PcloudPvminstancesGetallUnauthorized {
	return &PcloudPvminstancesGetallUnauthorized{}
}

/* PcloudPvminstancesGetallUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PcloudPvminstancesGetallUnauthorized struct {
	Payload *models.Error
}

func (o *PcloudPvminstancesGetallUnauthorized) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances][%d] pcloudPvminstancesGetallUnauthorized  %+v", 401, o.Payload)
}
func (o *PcloudPvminstancesGetallUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudPvminstancesGetallUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudPvminstancesGetallInternalServerError creates a PcloudPvminstancesGetallInternalServerError with default headers values
func NewPcloudPvminstancesGetallInternalServerError() *PcloudPvminstancesGetallInternalServerError {
	return &PcloudPvminstancesGetallInternalServerError{}
}

/* PcloudPvminstancesGetallInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PcloudPvminstancesGetallInternalServerError struct {
	Payload *models.Error
}

func (o *PcloudPvminstancesGetallInternalServerError) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances][%d] pcloudPvminstancesGetallInternalServerError  %+v", 500, o.Payload)
}
func (o *PcloudPvminstancesGetallInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudPvminstancesGetallInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
