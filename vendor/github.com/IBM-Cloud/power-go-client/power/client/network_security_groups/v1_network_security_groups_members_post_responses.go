// Code generated by go-swagger; DO NOT EDIT.

package network_security_groups

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/IBM-Cloud/power-go-client/power/models"
)

// V1NetworkSecurityGroupsMembersPostReader is a Reader for the V1NetworkSecurityGroupsMembersPost structure.
type V1NetworkSecurityGroupsMembersPostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1NetworkSecurityGroupsMembersPostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewV1NetworkSecurityGroupsMembersPostOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewV1NetworkSecurityGroupsMembersPostBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewV1NetworkSecurityGroupsMembersPostUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewV1NetworkSecurityGroupsMembersPostForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewV1NetworkSecurityGroupsMembersPostNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewV1NetworkSecurityGroupsMembersPostConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewV1NetworkSecurityGroupsMembersPostUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewV1NetworkSecurityGroupsMembersPostInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /v1/network-security-groups/{network_security_group_id}/members] v1.networkSecurityGroups.members.post", response, response.Code())
	}
}

// NewV1NetworkSecurityGroupsMembersPostOK creates a V1NetworkSecurityGroupsMembersPostOK with default headers values
func NewV1NetworkSecurityGroupsMembersPostOK() *V1NetworkSecurityGroupsMembersPostOK {
	return &V1NetworkSecurityGroupsMembersPostOK{}
}

/*
V1NetworkSecurityGroupsMembersPostOK describes a response with status code 200, with default header values.

OK
*/
type V1NetworkSecurityGroupsMembersPostOK struct {
	Payload *models.NetworkSecurityGroupMember
}

// IsSuccess returns true when this v1 network security groups members post o k response has a 2xx status code
func (o *V1NetworkSecurityGroupsMembersPostOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this v1 network security groups members post o k response has a 3xx status code
func (o *V1NetworkSecurityGroupsMembersPostOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this v1 network security groups members post o k response has a 4xx status code
func (o *V1NetworkSecurityGroupsMembersPostOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this v1 network security groups members post o k response has a 5xx status code
func (o *V1NetworkSecurityGroupsMembersPostOK) IsServerError() bool {
	return false
}

// IsCode returns true when this v1 network security groups members post o k response a status code equal to that given
func (o *V1NetworkSecurityGroupsMembersPostOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the v1 network security groups members post o k response
func (o *V1NetworkSecurityGroupsMembersPostOK) Code() int {
	return 200
}

func (o *V1NetworkSecurityGroupsMembersPostOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/network-security-groups/{network_security_group_id}/members][%d] v1NetworkSecurityGroupsMembersPostOK %s", 200, payload)
}

func (o *V1NetworkSecurityGroupsMembersPostOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/network-security-groups/{network_security_group_id}/members][%d] v1NetworkSecurityGroupsMembersPostOK %s", 200, payload)
}

func (o *V1NetworkSecurityGroupsMembersPostOK) GetPayload() *models.NetworkSecurityGroupMember {
	return o.Payload
}

func (o *V1NetworkSecurityGroupsMembersPostOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NetworkSecurityGroupMember)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV1NetworkSecurityGroupsMembersPostBadRequest creates a V1NetworkSecurityGroupsMembersPostBadRequest with default headers values
func NewV1NetworkSecurityGroupsMembersPostBadRequest() *V1NetworkSecurityGroupsMembersPostBadRequest {
	return &V1NetworkSecurityGroupsMembersPostBadRequest{}
}

/*
V1NetworkSecurityGroupsMembersPostBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type V1NetworkSecurityGroupsMembersPostBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this v1 network security groups members post bad request response has a 2xx status code
func (o *V1NetworkSecurityGroupsMembersPostBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this v1 network security groups members post bad request response has a 3xx status code
func (o *V1NetworkSecurityGroupsMembersPostBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this v1 network security groups members post bad request response has a 4xx status code
func (o *V1NetworkSecurityGroupsMembersPostBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this v1 network security groups members post bad request response has a 5xx status code
func (o *V1NetworkSecurityGroupsMembersPostBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this v1 network security groups members post bad request response a status code equal to that given
func (o *V1NetworkSecurityGroupsMembersPostBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the v1 network security groups members post bad request response
func (o *V1NetworkSecurityGroupsMembersPostBadRequest) Code() int {
	return 400
}

func (o *V1NetworkSecurityGroupsMembersPostBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/network-security-groups/{network_security_group_id}/members][%d] v1NetworkSecurityGroupsMembersPostBadRequest %s", 400, payload)
}

func (o *V1NetworkSecurityGroupsMembersPostBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/network-security-groups/{network_security_group_id}/members][%d] v1NetworkSecurityGroupsMembersPostBadRequest %s", 400, payload)
}

func (o *V1NetworkSecurityGroupsMembersPostBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *V1NetworkSecurityGroupsMembersPostBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV1NetworkSecurityGroupsMembersPostUnauthorized creates a V1NetworkSecurityGroupsMembersPostUnauthorized with default headers values
func NewV1NetworkSecurityGroupsMembersPostUnauthorized() *V1NetworkSecurityGroupsMembersPostUnauthorized {
	return &V1NetworkSecurityGroupsMembersPostUnauthorized{}
}

/*
V1NetworkSecurityGroupsMembersPostUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type V1NetworkSecurityGroupsMembersPostUnauthorized struct {
	Payload *models.Error
}

// IsSuccess returns true when this v1 network security groups members post unauthorized response has a 2xx status code
func (o *V1NetworkSecurityGroupsMembersPostUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this v1 network security groups members post unauthorized response has a 3xx status code
func (o *V1NetworkSecurityGroupsMembersPostUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this v1 network security groups members post unauthorized response has a 4xx status code
func (o *V1NetworkSecurityGroupsMembersPostUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this v1 network security groups members post unauthorized response has a 5xx status code
func (o *V1NetworkSecurityGroupsMembersPostUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this v1 network security groups members post unauthorized response a status code equal to that given
func (o *V1NetworkSecurityGroupsMembersPostUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the v1 network security groups members post unauthorized response
func (o *V1NetworkSecurityGroupsMembersPostUnauthorized) Code() int {
	return 401
}

func (o *V1NetworkSecurityGroupsMembersPostUnauthorized) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/network-security-groups/{network_security_group_id}/members][%d] v1NetworkSecurityGroupsMembersPostUnauthorized %s", 401, payload)
}

func (o *V1NetworkSecurityGroupsMembersPostUnauthorized) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/network-security-groups/{network_security_group_id}/members][%d] v1NetworkSecurityGroupsMembersPostUnauthorized %s", 401, payload)
}

func (o *V1NetworkSecurityGroupsMembersPostUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *V1NetworkSecurityGroupsMembersPostUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV1NetworkSecurityGroupsMembersPostForbidden creates a V1NetworkSecurityGroupsMembersPostForbidden with default headers values
func NewV1NetworkSecurityGroupsMembersPostForbidden() *V1NetworkSecurityGroupsMembersPostForbidden {
	return &V1NetworkSecurityGroupsMembersPostForbidden{}
}

/*
V1NetworkSecurityGroupsMembersPostForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type V1NetworkSecurityGroupsMembersPostForbidden struct {
	Payload *models.Error
}

// IsSuccess returns true when this v1 network security groups members post forbidden response has a 2xx status code
func (o *V1NetworkSecurityGroupsMembersPostForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this v1 network security groups members post forbidden response has a 3xx status code
func (o *V1NetworkSecurityGroupsMembersPostForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this v1 network security groups members post forbidden response has a 4xx status code
func (o *V1NetworkSecurityGroupsMembersPostForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this v1 network security groups members post forbidden response has a 5xx status code
func (o *V1NetworkSecurityGroupsMembersPostForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this v1 network security groups members post forbidden response a status code equal to that given
func (o *V1NetworkSecurityGroupsMembersPostForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the v1 network security groups members post forbidden response
func (o *V1NetworkSecurityGroupsMembersPostForbidden) Code() int {
	return 403
}

func (o *V1NetworkSecurityGroupsMembersPostForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/network-security-groups/{network_security_group_id}/members][%d] v1NetworkSecurityGroupsMembersPostForbidden %s", 403, payload)
}

func (o *V1NetworkSecurityGroupsMembersPostForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/network-security-groups/{network_security_group_id}/members][%d] v1NetworkSecurityGroupsMembersPostForbidden %s", 403, payload)
}

func (o *V1NetworkSecurityGroupsMembersPostForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *V1NetworkSecurityGroupsMembersPostForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV1NetworkSecurityGroupsMembersPostNotFound creates a V1NetworkSecurityGroupsMembersPostNotFound with default headers values
func NewV1NetworkSecurityGroupsMembersPostNotFound() *V1NetworkSecurityGroupsMembersPostNotFound {
	return &V1NetworkSecurityGroupsMembersPostNotFound{}
}

/*
V1NetworkSecurityGroupsMembersPostNotFound describes a response with status code 404, with default header values.

Not Found
*/
type V1NetworkSecurityGroupsMembersPostNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this v1 network security groups members post not found response has a 2xx status code
func (o *V1NetworkSecurityGroupsMembersPostNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this v1 network security groups members post not found response has a 3xx status code
func (o *V1NetworkSecurityGroupsMembersPostNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this v1 network security groups members post not found response has a 4xx status code
func (o *V1NetworkSecurityGroupsMembersPostNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this v1 network security groups members post not found response has a 5xx status code
func (o *V1NetworkSecurityGroupsMembersPostNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this v1 network security groups members post not found response a status code equal to that given
func (o *V1NetworkSecurityGroupsMembersPostNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the v1 network security groups members post not found response
func (o *V1NetworkSecurityGroupsMembersPostNotFound) Code() int {
	return 404
}

func (o *V1NetworkSecurityGroupsMembersPostNotFound) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/network-security-groups/{network_security_group_id}/members][%d] v1NetworkSecurityGroupsMembersPostNotFound %s", 404, payload)
}

func (o *V1NetworkSecurityGroupsMembersPostNotFound) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/network-security-groups/{network_security_group_id}/members][%d] v1NetworkSecurityGroupsMembersPostNotFound %s", 404, payload)
}

func (o *V1NetworkSecurityGroupsMembersPostNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *V1NetworkSecurityGroupsMembersPostNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV1NetworkSecurityGroupsMembersPostConflict creates a V1NetworkSecurityGroupsMembersPostConflict with default headers values
func NewV1NetworkSecurityGroupsMembersPostConflict() *V1NetworkSecurityGroupsMembersPostConflict {
	return &V1NetworkSecurityGroupsMembersPostConflict{}
}

/*
V1NetworkSecurityGroupsMembersPostConflict describes a response with status code 409, with default header values.

Conflict
*/
type V1NetworkSecurityGroupsMembersPostConflict struct {
	Payload *models.Error
}

// IsSuccess returns true when this v1 network security groups members post conflict response has a 2xx status code
func (o *V1NetworkSecurityGroupsMembersPostConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this v1 network security groups members post conflict response has a 3xx status code
func (o *V1NetworkSecurityGroupsMembersPostConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this v1 network security groups members post conflict response has a 4xx status code
func (o *V1NetworkSecurityGroupsMembersPostConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this v1 network security groups members post conflict response has a 5xx status code
func (o *V1NetworkSecurityGroupsMembersPostConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this v1 network security groups members post conflict response a status code equal to that given
func (o *V1NetworkSecurityGroupsMembersPostConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the v1 network security groups members post conflict response
func (o *V1NetworkSecurityGroupsMembersPostConflict) Code() int {
	return 409
}

func (o *V1NetworkSecurityGroupsMembersPostConflict) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/network-security-groups/{network_security_group_id}/members][%d] v1NetworkSecurityGroupsMembersPostConflict %s", 409, payload)
}

func (o *V1NetworkSecurityGroupsMembersPostConflict) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/network-security-groups/{network_security_group_id}/members][%d] v1NetworkSecurityGroupsMembersPostConflict %s", 409, payload)
}

func (o *V1NetworkSecurityGroupsMembersPostConflict) GetPayload() *models.Error {
	return o.Payload
}

func (o *V1NetworkSecurityGroupsMembersPostConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV1NetworkSecurityGroupsMembersPostUnprocessableEntity creates a V1NetworkSecurityGroupsMembersPostUnprocessableEntity with default headers values
func NewV1NetworkSecurityGroupsMembersPostUnprocessableEntity() *V1NetworkSecurityGroupsMembersPostUnprocessableEntity {
	return &V1NetworkSecurityGroupsMembersPostUnprocessableEntity{}
}

/*
V1NetworkSecurityGroupsMembersPostUnprocessableEntity describes a response with status code 422, with default header values.

Unprocessable Entity
*/
type V1NetworkSecurityGroupsMembersPostUnprocessableEntity struct {
	Payload *models.Error
}

// IsSuccess returns true when this v1 network security groups members post unprocessable entity response has a 2xx status code
func (o *V1NetworkSecurityGroupsMembersPostUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this v1 network security groups members post unprocessable entity response has a 3xx status code
func (o *V1NetworkSecurityGroupsMembersPostUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this v1 network security groups members post unprocessable entity response has a 4xx status code
func (o *V1NetworkSecurityGroupsMembersPostUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this v1 network security groups members post unprocessable entity response has a 5xx status code
func (o *V1NetworkSecurityGroupsMembersPostUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this v1 network security groups members post unprocessable entity response a status code equal to that given
func (o *V1NetworkSecurityGroupsMembersPostUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the v1 network security groups members post unprocessable entity response
func (o *V1NetworkSecurityGroupsMembersPostUnprocessableEntity) Code() int {
	return 422
}

func (o *V1NetworkSecurityGroupsMembersPostUnprocessableEntity) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/network-security-groups/{network_security_group_id}/members][%d] v1NetworkSecurityGroupsMembersPostUnprocessableEntity %s", 422, payload)
}

func (o *V1NetworkSecurityGroupsMembersPostUnprocessableEntity) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/network-security-groups/{network_security_group_id}/members][%d] v1NetworkSecurityGroupsMembersPostUnprocessableEntity %s", 422, payload)
}

func (o *V1NetworkSecurityGroupsMembersPostUnprocessableEntity) GetPayload() *models.Error {
	return o.Payload
}

func (o *V1NetworkSecurityGroupsMembersPostUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV1NetworkSecurityGroupsMembersPostInternalServerError creates a V1NetworkSecurityGroupsMembersPostInternalServerError with default headers values
func NewV1NetworkSecurityGroupsMembersPostInternalServerError() *V1NetworkSecurityGroupsMembersPostInternalServerError {
	return &V1NetworkSecurityGroupsMembersPostInternalServerError{}
}

/*
V1NetworkSecurityGroupsMembersPostInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type V1NetworkSecurityGroupsMembersPostInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this v1 network security groups members post internal server error response has a 2xx status code
func (o *V1NetworkSecurityGroupsMembersPostInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this v1 network security groups members post internal server error response has a 3xx status code
func (o *V1NetworkSecurityGroupsMembersPostInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this v1 network security groups members post internal server error response has a 4xx status code
func (o *V1NetworkSecurityGroupsMembersPostInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this v1 network security groups members post internal server error response has a 5xx status code
func (o *V1NetworkSecurityGroupsMembersPostInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this v1 network security groups members post internal server error response a status code equal to that given
func (o *V1NetworkSecurityGroupsMembersPostInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the v1 network security groups members post internal server error response
func (o *V1NetworkSecurityGroupsMembersPostInternalServerError) Code() int {
	return 500
}

func (o *V1NetworkSecurityGroupsMembersPostInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/network-security-groups/{network_security_group_id}/members][%d] v1NetworkSecurityGroupsMembersPostInternalServerError %s", 500, payload)
}

func (o *V1NetworkSecurityGroupsMembersPostInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/network-security-groups/{network_security_group_id}/members][%d] v1NetworkSecurityGroupsMembersPostInternalServerError %s", 500, payload)
}

func (o *V1NetworkSecurityGroupsMembersPostInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *V1NetworkSecurityGroupsMembersPostInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}