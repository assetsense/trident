// Code generated by go-swagger; DO NOT EDIT.

package s_a_n

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/netapp/trident/storage_drivers/ontap/api/rest/models"
)

// LunAttributeModifyReader is a Reader for the LunAttributeModify structure.
type LunAttributeModifyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *LunAttributeModifyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewLunAttributeModifyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewLunAttributeModifyDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewLunAttributeModifyOK creates a LunAttributeModifyOK with default headers values
func NewLunAttributeModifyOK() *LunAttributeModifyOK {
	return &LunAttributeModifyOK{}
}

/* LunAttributeModifyOK describes a response with status code 200, with default header values.

OK
*/
type LunAttributeModifyOK struct {
}

func (o *LunAttributeModifyOK) Error() string {
	return fmt.Sprintf("[PATCH /storage/luns/{lun.uuid}/attributes/{name}][%d] lunAttributeModifyOK ", 200)
}

func (o *LunAttributeModifyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewLunAttributeModifyDefault creates a LunAttributeModifyDefault with default headers values
func NewLunAttributeModifyDefault(code int) *LunAttributeModifyDefault {
	return &LunAttributeModifyDefault{
		_statusCode: code,
	}
}

/* LunAttributeModifyDefault describes a response with status code -1, with default header values.

 ONTAP Error Response Codes
| Error Code | Description |
| ---------- | ----------- |
| 5374875 | The specified LUN was not found. |
| 5374929 | The combined sizes of an attribute name and value are too large. |
| 5374931 | The specified attribute was not found. |

*/
type LunAttributeModifyDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the lun attribute modify default response
func (o *LunAttributeModifyDefault) Code() int {
	return o._statusCode
}

func (o *LunAttributeModifyDefault) Error() string {
	return fmt.Sprintf("[PATCH /storage/luns/{lun.uuid}/attributes/{name}][%d] lun_attribute_modify default  %+v", o._statusCode, o.Payload)
}
func (o *LunAttributeModifyDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *LunAttributeModifyDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}