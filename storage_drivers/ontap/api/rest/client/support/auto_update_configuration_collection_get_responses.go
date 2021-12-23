// Code generated by go-swagger; DO NOT EDIT.

package support

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/netapp/trident/storage_drivers/ontap/api/rest/models"
)

// AutoUpdateConfigurationCollectionGetReader is a Reader for the AutoUpdateConfigurationCollectionGet structure.
type AutoUpdateConfigurationCollectionGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AutoUpdateConfigurationCollectionGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAutoUpdateConfigurationCollectionGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewAutoUpdateConfigurationCollectionGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAutoUpdateConfigurationCollectionGetOK creates a AutoUpdateConfigurationCollectionGetOK with default headers values
func NewAutoUpdateConfigurationCollectionGetOK() *AutoUpdateConfigurationCollectionGetOK {
	return &AutoUpdateConfigurationCollectionGetOK{}
}

/* AutoUpdateConfigurationCollectionGetOK describes a response with status code 200, with default header values.

OK
*/
type AutoUpdateConfigurationCollectionGetOK struct {
	Payload *models.AutoUpdateConfigurationResponse
}

func (o *AutoUpdateConfigurationCollectionGetOK) Error() string {
	return fmt.Sprintf("[GET /support/auto-update/configurations][%d] autoUpdateConfigurationCollectionGetOK  %+v", 200, o.Payload)
}
func (o *AutoUpdateConfigurationCollectionGetOK) GetPayload() *models.AutoUpdateConfigurationResponse {
	return o.Payload
}

func (o *AutoUpdateConfigurationCollectionGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AutoUpdateConfigurationResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAutoUpdateConfigurationCollectionGetDefault creates a AutoUpdateConfigurationCollectionGetDefault with default headers values
func NewAutoUpdateConfigurationCollectionGetDefault(code int) *AutoUpdateConfigurationCollectionGetDefault {
	return &AutoUpdateConfigurationCollectionGetDefault{
		_statusCode: code,
	}
}

/* AutoUpdateConfigurationCollectionGetDefault describes a response with status code -1, with default header values.

Error
*/
type AutoUpdateConfigurationCollectionGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the auto update configuration collection get default response
func (o *AutoUpdateConfigurationCollectionGetDefault) Code() int {
	return o._statusCode
}

func (o *AutoUpdateConfigurationCollectionGetDefault) Error() string {
	return fmt.Sprintf("[GET /support/auto-update/configurations][%d] auto_update_configuration_collection_get default  %+v", o._statusCode, o.Payload)
}
func (o *AutoUpdateConfigurationCollectionGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *AutoUpdateConfigurationCollectionGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}