// Code generated by go-swagger; DO NOT EDIT.

package p_cloud_volumes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/IBM-Cloud/power-go-client/power/models"
)

// PcloudCloudinstancesVolumesFlashCopyMappingsGetReader is a Reader for the PcloudCloudinstancesVolumesFlashCopyMappingsGet structure.
type PcloudCloudinstancesVolumesFlashCopyMappingsGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PcloudCloudinstancesVolumesFlashCopyMappingsGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPcloudCloudinstancesVolumesFlashCopyMappingsGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPcloudCloudinstancesVolumesFlashCopyMappingsGetBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPcloudCloudinstancesVolumesFlashCopyMappingsGetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPcloudCloudinstancesVolumesFlashCopyMappingsGetForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPcloudCloudinstancesVolumesFlashCopyMappingsGetNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPcloudCloudinstancesVolumesFlashCopyMappingsGetTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPcloudCloudinstancesVolumesFlashCopyMappingsGetInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPcloudCloudinstancesVolumesFlashCopyMappingsGetOK creates a PcloudCloudinstancesVolumesFlashCopyMappingsGetOK with default headers values
func NewPcloudCloudinstancesVolumesFlashCopyMappingsGetOK() *PcloudCloudinstancesVolumesFlashCopyMappingsGetOK {
	return &PcloudCloudinstancesVolumesFlashCopyMappingsGetOK{}
}

/* PcloudCloudinstancesVolumesFlashCopyMappingsGetOK describes a response with status code 200, with default header values.

OK
*/
type PcloudCloudinstancesVolumesFlashCopyMappingsGetOK struct {
	Payload models.FlashCopyMappings
}

func (o *PcloudCloudinstancesVolumesFlashCopyMappingsGetOK) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/volumes/{volume_id}/flash-copy-mappings][%d] pcloudCloudinstancesVolumesFlashCopyMappingsGetOK  %+v", 200, o.Payload)
}
func (o *PcloudCloudinstancesVolumesFlashCopyMappingsGetOK) GetPayload() models.FlashCopyMappings {
	return o.Payload
}

func (o *PcloudCloudinstancesVolumesFlashCopyMappingsGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudCloudinstancesVolumesFlashCopyMappingsGetBadRequest creates a PcloudCloudinstancesVolumesFlashCopyMappingsGetBadRequest with default headers values
func NewPcloudCloudinstancesVolumesFlashCopyMappingsGetBadRequest() *PcloudCloudinstancesVolumesFlashCopyMappingsGetBadRequest {
	return &PcloudCloudinstancesVolumesFlashCopyMappingsGetBadRequest{}
}

/* PcloudCloudinstancesVolumesFlashCopyMappingsGetBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PcloudCloudinstancesVolumesFlashCopyMappingsGetBadRequest struct {
	Payload *models.Error
}

func (o *PcloudCloudinstancesVolumesFlashCopyMappingsGetBadRequest) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/volumes/{volume_id}/flash-copy-mappings][%d] pcloudCloudinstancesVolumesFlashCopyMappingsGetBadRequest  %+v", 400, o.Payload)
}
func (o *PcloudCloudinstancesVolumesFlashCopyMappingsGetBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudCloudinstancesVolumesFlashCopyMappingsGetBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudCloudinstancesVolumesFlashCopyMappingsGetUnauthorized creates a PcloudCloudinstancesVolumesFlashCopyMappingsGetUnauthorized with default headers values
func NewPcloudCloudinstancesVolumesFlashCopyMappingsGetUnauthorized() *PcloudCloudinstancesVolumesFlashCopyMappingsGetUnauthorized {
	return &PcloudCloudinstancesVolumesFlashCopyMappingsGetUnauthorized{}
}

/* PcloudCloudinstancesVolumesFlashCopyMappingsGetUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PcloudCloudinstancesVolumesFlashCopyMappingsGetUnauthorized struct {
	Payload *models.Error
}

func (o *PcloudCloudinstancesVolumesFlashCopyMappingsGetUnauthorized) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/volumes/{volume_id}/flash-copy-mappings][%d] pcloudCloudinstancesVolumesFlashCopyMappingsGetUnauthorized  %+v", 401, o.Payload)
}
func (o *PcloudCloudinstancesVolumesFlashCopyMappingsGetUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudCloudinstancesVolumesFlashCopyMappingsGetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudCloudinstancesVolumesFlashCopyMappingsGetForbidden creates a PcloudCloudinstancesVolumesFlashCopyMappingsGetForbidden with default headers values
func NewPcloudCloudinstancesVolumesFlashCopyMappingsGetForbidden() *PcloudCloudinstancesVolumesFlashCopyMappingsGetForbidden {
	return &PcloudCloudinstancesVolumesFlashCopyMappingsGetForbidden{}
}

/* PcloudCloudinstancesVolumesFlashCopyMappingsGetForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PcloudCloudinstancesVolumesFlashCopyMappingsGetForbidden struct {
	Payload *models.Error
}

func (o *PcloudCloudinstancesVolumesFlashCopyMappingsGetForbidden) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/volumes/{volume_id}/flash-copy-mappings][%d] pcloudCloudinstancesVolumesFlashCopyMappingsGetForbidden  %+v", 403, o.Payload)
}
func (o *PcloudCloudinstancesVolumesFlashCopyMappingsGetForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudCloudinstancesVolumesFlashCopyMappingsGetForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudCloudinstancesVolumesFlashCopyMappingsGetNotFound creates a PcloudCloudinstancesVolumesFlashCopyMappingsGetNotFound with default headers values
func NewPcloudCloudinstancesVolumesFlashCopyMappingsGetNotFound() *PcloudCloudinstancesVolumesFlashCopyMappingsGetNotFound {
	return &PcloudCloudinstancesVolumesFlashCopyMappingsGetNotFound{}
}

/* PcloudCloudinstancesVolumesFlashCopyMappingsGetNotFound describes a response with status code 404, with default header values.

Not Found
*/
type PcloudCloudinstancesVolumesFlashCopyMappingsGetNotFound struct {
	Payload *models.Error
}

func (o *PcloudCloudinstancesVolumesFlashCopyMappingsGetNotFound) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/volumes/{volume_id}/flash-copy-mappings][%d] pcloudCloudinstancesVolumesFlashCopyMappingsGetNotFound  %+v", 404, o.Payload)
}
func (o *PcloudCloudinstancesVolumesFlashCopyMappingsGetNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudCloudinstancesVolumesFlashCopyMappingsGetNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudCloudinstancesVolumesFlashCopyMappingsGetTooManyRequests creates a PcloudCloudinstancesVolumesFlashCopyMappingsGetTooManyRequests with default headers values
func NewPcloudCloudinstancesVolumesFlashCopyMappingsGetTooManyRequests() *PcloudCloudinstancesVolumesFlashCopyMappingsGetTooManyRequests {
	return &PcloudCloudinstancesVolumesFlashCopyMappingsGetTooManyRequests{}
}

/* PcloudCloudinstancesVolumesFlashCopyMappingsGetTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type PcloudCloudinstancesVolumesFlashCopyMappingsGetTooManyRequests struct {
	Payload *models.Error
}

func (o *PcloudCloudinstancesVolumesFlashCopyMappingsGetTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/volumes/{volume_id}/flash-copy-mappings][%d] pcloudCloudinstancesVolumesFlashCopyMappingsGetTooManyRequests  %+v", 429, o.Payload)
}
func (o *PcloudCloudinstancesVolumesFlashCopyMappingsGetTooManyRequests) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudCloudinstancesVolumesFlashCopyMappingsGetTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudCloudinstancesVolumesFlashCopyMappingsGetInternalServerError creates a PcloudCloudinstancesVolumesFlashCopyMappingsGetInternalServerError with default headers values
func NewPcloudCloudinstancesVolumesFlashCopyMappingsGetInternalServerError() *PcloudCloudinstancesVolumesFlashCopyMappingsGetInternalServerError {
	return &PcloudCloudinstancesVolumesFlashCopyMappingsGetInternalServerError{}
}

/* PcloudCloudinstancesVolumesFlashCopyMappingsGetInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PcloudCloudinstancesVolumesFlashCopyMappingsGetInternalServerError struct {
	Payload *models.Error
}

func (o *PcloudCloudinstancesVolumesFlashCopyMappingsGetInternalServerError) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/volumes/{volume_id}/flash-copy-mappings][%d] pcloudCloudinstancesVolumesFlashCopyMappingsGetInternalServerError  %+v", 500, o.Payload)
}
func (o *PcloudCloudinstancesVolumesFlashCopyMappingsGetInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudCloudinstancesVolumesFlashCopyMappingsGetInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
