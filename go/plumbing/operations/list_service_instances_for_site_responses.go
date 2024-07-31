// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/j0urneyK/open-api/v2/go/models"
)

// ListServiceInstancesForSiteReader is a Reader for the ListServiceInstancesForSite structure.
type ListServiceInstancesForSiteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListServiceInstancesForSiteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListServiceInstancesForSiteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewListServiceInstancesForSiteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListServiceInstancesForSiteOK creates a ListServiceInstancesForSiteOK with default headers values
func NewListServiceInstancesForSiteOK() *ListServiceInstancesForSiteOK {
	return &ListServiceInstancesForSiteOK{}
}

/*
ListServiceInstancesForSiteOK handles this case with default header values.

OK
*/
type ListServiceInstancesForSiteOK struct {
	Payload []*models.ServiceInstance
}

func (o *ListServiceInstancesForSiteOK) Error() string {
	return fmt.Sprintf("[GET /sites/{site_id}/service-instances][%d] listServiceInstancesForSiteOK  %+v", 200, o.Payload)
}

func (o *ListServiceInstancesForSiteOK) GetPayload() []*models.ServiceInstance {
	return o.Payload
}

func (o *ListServiceInstancesForSiteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListServiceInstancesForSiteDefault creates a ListServiceInstancesForSiteDefault with default headers values
func NewListServiceInstancesForSiteDefault(code int) *ListServiceInstancesForSiteDefault {
	return &ListServiceInstancesForSiteDefault{
		_statusCode: code,
	}
}

/*
ListServiceInstancesForSiteDefault handles this case with default header values.

error
*/
type ListServiceInstancesForSiteDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the list service instances for site default response
func (o *ListServiceInstancesForSiteDefault) Code() int {
	return o._statusCode
}

func (o *ListServiceInstancesForSiteDefault) Error() string {
	return fmt.Sprintf("[GET /sites/{site_id}/service-instances][%d] listServiceInstancesForSite default  %+v", o._statusCode, o.Payload)
}

func (o *ListServiceInstancesForSiteDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *ListServiceInstancesForSiteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
