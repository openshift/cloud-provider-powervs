// Code generated by go-swagger; DO NOT EDIT.

package p_cloud_cloud_connections

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/IBM-Cloud/power-go-client/power/models"
)

// PcloudCloudconnectionsNetworksPutReader is a Reader for the PcloudCloudconnectionsNetworksPut structure.
type PcloudCloudconnectionsNetworksPutReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PcloudCloudconnectionsNetworksPutReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPcloudCloudconnectionsNetworksPutOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewPcloudCloudconnectionsNetworksPutAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPcloudCloudconnectionsNetworksPutBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPcloudCloudconnectionsNetworksPutUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPcloudCloudconnectionsNetworksPutNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPcloudCloudconnectionsNetworksPutRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewPcloudCloudconnectionsNetworksPutUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPcloudCloudconnectionsNetworksPutInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPcloudCloudconnectionsNetworksPutOK creates a PcloudCloudconnectionsNetworksPutOK with default headers values
func NewPcloudCloudconnectionsNetworksPutOK() *PcloudCloudconnectionsNetworksPutOK {
	return &PcloudCloudconnectionsNetworksPutOK{}
}

/*
PcloudCloudconnectionsNetworksPutOK describes a response with status code 200, with default header values.

OK
*/
type PcloudCloudconnectionsNetworksPutOK struct {
	Payload models.Object
}

// IsSuccess returns true when this pcloud cloudconnections networks put o k response has a 2xx status code
func (o *PcloudCloudconnectionsNetworksPutOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this pcloud cloudconnections networks put o k response has a 3xx status code
func (o *PcloudCloudconnectionsNetworksPutOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud cloudconnections networks put o k response has a 4xx status code
func (o *PcloudCloudconnectionsNetworksPutOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this pcloud cloudconnections networks put o k response has a 5xx status code
func (o *PcloudCloudconnectionsNetworksPutOK) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud cloudconnections networks put o k response a status code equal to that given
func (o *PcloudCloudconnectionsNetworksPutOK) IsCode(code int) bool {
	return code == 200
}

func (o *PcloudCloudconnectionsNetworksPutOK) Error() string {
	return fmt.Sprintf("[PUT /pcloud/v1/cloud-instances/{cloud_instance_id}/cloud-connections/{cloud_connection_id}/networks/{network_id}][%d] pcloudCloudconnectionsNetworksPutOK  %+v", 200, o.Payload)
}

func (o *PcloudCloudconnectionsNetworksPutOK) String() string {
	return fmt.Sprintf("[PUT /pcloud/v1/cloud-instances/{cloud_instance_id}/cloud-connections/{cloud_connection_id}/networks/{network_id}][%d] pcloudCloudconnectionsNetworksPutOK  %+v", 200, o.Payload)
}

func (o *PcloudCloudconnectionsNetworksPutOK) GetPayload() models.Object {
	return o.Payload
}

func (o *PcloudCloudconnectionsNetworksPutOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudCloudconnectionsNetworksPutAccepted creates a PcloudCloudconnectionsNetworksPutAccepted with default headers values
func NewPcloudCloudconnectionsNetworksPutAccepted() *PcloudCloudconnectionsNetworksPutAccepted {
	return &PcloudCloudconnectionsNetworksPutAccepted{}
}

/*
PcloudCloudconnectionsNetworksPutAccepted describes a response with status code 202, with default header values.

Accepted
*/
type PcloudCloudconnectionsNetworksPutAccepted struct {
	Payload *models.JobReference
}

// IsSuccess returns true when this pcloud cloudconnections networks put accepted response has a 2xx status code
func (o *PcloudCloudconnectionsNetworksPutAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this pcloud cloudconnections networks put accepted response has a 3xx status code
func (o *PcloudCloudconnectionsNetworksPutAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud cloudconnections networks put accepted response has a 4xx status code
func (o *PcloudCloudconnectionsNetworksPutAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this pcloud cloudconnections networks put accepted response has a 5xx status code
func (o *PcloudCloudconnectionsNetworksPutAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud cloudconnections networks put accepted response a status code equal to that given
func (o *PcloudCloudconnectionsNetworksPutAccepted) IsCode(code int) bool {
	return code == 202
}

func (o *PcloudCloudconnectionsNetworksPutAccepted) Error() string {
	return fmt.Sprintf("[PUT /pcloud/v1/cloud-instances/{cloud_instance_id}/cloud-connections/{cloud_connection_id}/networks/{network_id}][%d] pcloudCloudconnectionsNetworksPutAccepted  %+v", 202, o.Payload)
}

func (o *PcloudCloudconnectionsNetworksPutAccepted) String() string {
	return fmt.Sprintf("[PUT /pcloud/v1/cloud-instances/{cloud_instance_id}/cloud-connections/{cloud_connection_id}/networks/{network_id}][%d] pcloudCloudconnectionsNetworksPutAccepted  %+v", 202, o.Payload)
}

func (o *PcloudCloudconnectionsNetworksPutAccepted) GetPayload() *models.JobReference {
	return o.Payload
}

func (o *PcloudCloudconnectionsNetworksPutAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JobReference)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudCloudconnectionsNetworksPutBadRequest creates a PcloudCloudconnectionsNetworksPutBadRequest with default headers values
func NewPcloudCloudconnectionsNetworksPutBadRequest() *PcloudCloudconnectionsNetworksPutBadRequest {
	return &PcloudCloudconnectionsNetworksPutBadRequest{}
}

/*
PcloudCloudconnectionsNetworksPutBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PcloudCloudconnectionsNetworksPutBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud cloudconnections networks put bad request response has a 2xx status code
func (o *PcloudCloudconnectionsNetworksPutBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud cloudconnections networks put bad request response has a 3xx status code
func (o *PcloudCloudconnectionsNetworksPutBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud cloudconnections networks put bad request response has a 4xx status code
func (o *PcloudCloudconnectionsNetworksPutBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud cloudconnections networks put bad request response has a 5xx status code
func (o *PcloudCloudconnectionsNetworksPutBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud cloudconnections networks put bad request response a status code equal to that given
func (o *PcloudCloudconnectionsNetworksPutBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PcloudCloudconnectionsNetworksPutBadRequest) Error() string {
	return fmt.Sprintf("[PUT /pcloud/v1/cloud-instances/{cloud_instance_id}/cloud-connections/{cloud_connection_id}/networks/{network_id}][%d] pcloudCloudconnectionsNetworksPutBadRequest  %+v", 400, o.Payload)
}

func (o *PcloudCloudconnectionsNetworksPutBadRequest) String() string {
	return fmt.Sprintf("[PUT /pcloud/v1/cloud-instances/{cloud_instance_id}/cloud-connections/{cloud_connection_id}/networks/{network_id}][%d] pcloudCloudconnectionsNetworksPutBadRequest  %+v", 400, o.Payload)
}

func (o *PcloudCloudconnectionsNetworksPutBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudCloudconnectionsNetworksPutBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudCloudconnectionsNetworksPutUnauthorized creates a PcloudCloudconnectionsNetworksPutUnauthorized with default headers values
func NewPcloudCloudconnectionsNetworksPutUnauthorized() *PcloudCloudconnectionsNetworksPutUnauthorized {
	return &PcloudCloudconnectionsNetworksPutUnauthorized{}
}

/*
PcloudCloudconnectionsNetworksPutUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PcloudCloudconnectionsNetworksPutUnauthorized struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud cloudconnections networks put unauthorized response has a 2xx status code
func (o *PcloudCloudconnectionsNetworksPutUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud cloudconnections networks put unauthorized response has a 3xx status code
func (o *PcloudCloudconnectionsNetworksPutUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud cloudconnections networks put unauthorized response has a 4xx status code
func (o *PcloudCloudconnectionsNetworksPutUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud cloudconnections networks put unauthorized response has a 5xx status code
func (o *PcloudCloudconnectionsNetworksPutUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud cloudconnections networks put unauthorized response a status code equal to that given
func (o *PcloudCloudconnectionsNetworksPutUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *PcloudCloudconnectionsNetworksPutUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /pcloud/v1/cloud-instances/{cloud_instance_id}/cloud-connections/{cloud_connection_id}/networks/{network_id}][%d] pcloudCloudconnectionsNetworksPutUnauthorized  %+v", 401, o.Payload)
}

func (o *PcloudCloudconnectionsNetworksPutUnauthorized) String() string {
	return fmt.Sprintf("[PUT /pcloud/v1/cloud-instances/{cloud_instance_id}/cloud-connections/{cloud_connection_id}/networks/{network_id}][%d] pcloudCloudconnectionsNetworksPutUnauthorized  %+v", 401, o.Payload)
}

func (o *PcloudCloudconnectionsNetworksPutUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudCloudconnectionsNetworksPutUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudCloudconnectionsNetworksPutNotFound creates a PcloudCloudconnectionsNetworksPutNotFound with default headers values
func NewPcloudCloudconnectionsNetworksPutNotFound() *PcloudCloudconnectionsNetworksPutNotFound {
	return &PcloudCloudconnectionsNetworksPutNotFound{}
}

/*
PcloudCloudconnectionsNetworksPutNotFound describes a response with status code 404, with default header values.

Not Found
*/
type PcloudCloudconnectionsNetworksPutNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud cloudconnections networks put not found response has a 2xx status code
func (o *PcloudCloudconnectionsNetworksPutNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud cloudconnections networks put not found response has a 3xx status code
func (o *PcloudCloudconnectionsNetworksPutNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud cloudconnections networks put not found response has a 4xx status code
func (o *PcloudCloudconnectionsNetworksPutNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud cloudconnections networks put not found response has a 5xx status code
func (o *PcloudCloudconnectionsNetworksPutNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud cloudconnections networks put not found response a status code equal to that given
func (o *PcloudCloudconnectionsNetworksPutNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *PcloudCloudconnectionsNetworksPutNotFound) Error() string {
	return fmt.Sprintf("[PUT /pcloud/v1/cloud-instances/{cloud_instance_id}/cloud-connections/{cloud_connection_id}/networks/{network_id}][%d] pcloudCloudconnectionsNetworksPutNotFound  %+v", 404, o.Payload)
}

func (o *PcloudCloudconnectionsNetworksPutNotFound) String() string {
	return fmt.Sprintf("[PUT /pcloud/v1/cloud-instances/{cloud_instance_id}/cloud-connections/{cloud_connection_id}/networks/{network_id}][%d] pcloudCloudconnectionsNetworksPutNotFound  %+v", 404, o.Payload)
}

func (o *PcloudCloudconnectionsNetworksPutNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudCloudconnectionsNetworksPutNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudCloudconnectionsNetworksPutRequestTimeout creates a PcloudCloudconnectionsNetworksPutRequestTimeout with default headers values
func NewPcloudCloudconnectionsNetworksPutRequestTimeout() *PcloudCloudconnectionsNetworksPutRequestTimeout {
	return &PcloudCloudconnectionsNetworksPutRequestTimeout{}
}

/*
PcloudCloudconnectionsNetworksPutRequestTimeout describes a response with status code 408, with default header values.

Request Timeout
*/
type PcloudCloudconnectionsNetworksPutRequestTimeout struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud cloudconnections networks put request timeout response has a 2xx status code
func (o *PcloudCloudconnectionsNetworksPutRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud cloudconnections networks put request timeout response has a 3xx status code
func (o *PcloudCloudconnectionsNetworksPutRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud cloudconnections networks put request timeout response has a 4xx status code
func (o *PcloudCloudconnectionsNetworksPutRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud cloudconnections networks put request timeout response has a 5xx status code
func (o *PcloudCloudconnectionsNetworksPutRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud cloudconnections networks put request timeout response a status code equal to that given
func (o *PcloudCloudconnectionsNetworksPutRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *PcloudCloudconnectionsNetworksPutRequestTimeout) Error() string {
	return fmt.Sprintf("[PUT /pcloud/v1/cloud-instances/{cloud_instance_id}/cloud-connections/{cloud_connection_id}/networks/{network_id}][%d] pcloudCloudconnectionsNetworksPutRequestTimeout  %+v", 408, o.Payload)
}

func (o *PcloudCloudconnectionsNetworksPutRequestTimeout) String() string {
	return fmt.Sprintf("[PUT /pcloud/v1/cloud-instances/{cloud_instance_id}/cloud-connections/{cloud_connection_id}/networks/{network_id}][%d] pcloudCloudconnectionsNetworksPutRequestTimeout  %+v", 408, o.Payload)
}

func (o *PcloudCloudconnectionsNetworksPutRequestTimeout) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudCloudconnectionsNetworksPutRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudCloudconnectionsNetworksPutUnprocessableEntity creates a PcloudCloudconnectionsNetworksPutUnprocessableEntity with default headers values
func NewPcloudCloudconnectionsNetworksPutUnprocessableEntity() *PcloudCloudconnectionsNetworksPutUnprocessableEntity {
	return &PcloudCloudconnectionsNetworksPutUnprocessableEntity{}
}

/*
PcloudCloudconnectionsNetworksPutUnprocessableEntity describes a response with status code 422, with default header values.

Unprocessable Entity
*/
type PcloudCloudconnectionsNetworksPutUnprocessableEntity struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud cloudconnections networks put unprocessable entity response has a 2xx status code
func (o *PcloudCloudconnectionsNetworksPutUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud cloudconnections networks put unprocessable entity response has a 3xx status code
func (o *PcloudCloudconnectionsNetworksPutUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud cloudconnections networks put unprocessable entity response has a 4xx status code
func (o *PcloudCloudconnectionsNetworksPutUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud cloudconnections networks put unprocessable entity response has a 5xx status code
func (o *PcloudCloudconnectionsNetworksPutUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud cloudconnections networks put unprocessable entity response a status code equal to that given
func (o *PcloudCloudconnectionsNetworksPutUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

func (o *PcloudCloudconnectionsNetworksPutUnprocessableEntity) Error() string {
	return fmt.Sprintf("[PUT /pcloud/v1/cloud-instances/{cloud_instance_id}/cloud-connections/{cloud_connection_id}/networks/{network_id}][%d] pcloudCloudconnectionsNetworksPutUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *PcloudCloudconnectionsNetworksPutUnprocessableEntity) String() string {
	return fmt.Sprintf("[PUT /pcloud/v1/cloud-instances/{cloud_instance_id}/cloud-connections/{cloud_connection_id}/networks/{network_id}][%d] pcloudCloudconnectionsNetworksPutUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *PcloudCloudconnectionsNetworksPutUnprocessableEntity) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudCloudconnectionsNetworksPutUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudCloudconnectionsNetworksPutInternalServerError creates a PcloudCloudconnectionsNetworksPutInternalServerError with default headers values
func NewPcloudCloudconnectionsNetworksPutInternalServerError() *PcloudCloudconnectionsNetworksPutInternalServerError {
	return &PcloudCloudconnectionsNetworksPutInternalServerError{}
}

/*
PcloudCloudconnectionsNetworksPutInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PcloudCloudconnectionsNetworksPutInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud cloudconnections networks put internal server error response has a 2xx status code
func (o *PcloudCloudconnectionsNetworksPutInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud cloudconnections networks put internal server error response has a 3xx status code
func (o *PcloudCloudconnectionsNetworksPutInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud cloudconnections networks put internal server error response has a 4xx status code
func (o *PcloudCloudconnectionsNetworksPutInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this pcloud cloudconnections networks put internal server error response has a 5xx status code
func (o *PcloudCloudconnectionsNetworksPutInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this pcloud cloudconnections networks put internal server error response a status code equal to that given
func (o *PcloudCloudconnectionsNetworksPutInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PcloudCloudconnectionsNetworksPutInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /pcloud/v1/cloud-instances/{cloud_instance_id}/cloud-connections/{cloud_connection_id}/networks/{network_id}][%d] pcloudCloudconnectionsNetworksPutInternalServerError  %+v", 500, o.Payload)
}

func (o *PcloudCloudconnectionsNetworksPutInternalServerError) String() string {
	return fmt.Sprintf("[PUT /pcloud/v1/cloud-instances/{cloud_instance_id}/cloud-connections/{cloud_connection_id}/networks/{network_id}][%d] pcloudCloudconnectionsNetworksPutInternalServerError  %+v", 500, o.Payload)
}

func (o *PcloudCloudconnectionsNetworksPutInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudCloudconnectionsNetworksPutInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}