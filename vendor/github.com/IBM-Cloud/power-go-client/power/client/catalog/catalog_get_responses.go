// Code generated by go-swagger; DO NOT EDIT.

package catalog

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/IBM-Cloud/power-go-client/power/models"
)

// CatalogGetReader is a Reader for the CatalogGet structure.
type CatalogGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CatalogGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCatalogGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCatalogGetOK creates a CatalogGetOK with default headers values
func NewCatalogGetOK() *CatalogGetOK {
	return &CatalogGetOK{}
}

/* CatalogGetOK describes a response with status code 200, with default header values.

catalog response
*/
type CatalogGetOK struct {
	Payload *models.Catalog
}

func (o *CatalogGetOK) Error() string {
	return fmt.Sprintf("[GET /v2/catalog][%d] catalogGetOK  %+v", 200, o.Payload)
}
func (o *CatalogGetOK) GetPayload() *models.Catalog {
	return o.Payload
}

func (o *CatalogGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Catalog)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
