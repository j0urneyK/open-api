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

// CreateServiceInstanceReader is a Reader for the CreateServiceInstance structure.
type CreateServiceInstanceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateServiceInstanceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateServiceInstanceCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCreateServiceInstanceDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateServiceInstanceCreated creates a CreateServiceInstanceCreated with default headers values
func NewCreateServiceInstanceCreated() *CreateServiceInstanceCreated {
	return &CreateServiceInstanceCreated{}
}

/*
CreateServiceInstanceCreated handles this case with default header values.

Created
*/
type CreateServiceInstanceCreated struct {
	Payload *models.ServiceInstance
}

func (o *CreateServiceInstanceCreated) Error() string {
	return fmt.Sprintf("[POST /sites/{site_id}/services/{addon}/instances][%d] createServiceInstanceCreated  %+v", 201, o.Payload)
}

func (o *CreateServiceInstanceCreated) GetPayload() *models.ServiceInstance {
	return o.Payload
}

func (o *CreateServiceInstanceCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceInstance)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateServiceInstanceDefault creates a CreateServiceInstanceDefault with default headers values
func NewCreateServiceInstanceDefault(code int) *CreateServiceInstanceDefault {
	return &CreateServiceInstanceDefault{
		_statusCode: code,
	}
}

/*
CreateServiceInstanceDefault handles this case with default header values.

error
*/
type CreateServiceInstanceDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the create service instance default response
func (o *CreateServiceInstanceDefault) Code() int {
	return o._statusCode
}

func (o *CreateServiceInstanceDefault) Error() string {
	return fmt.Sprintf("[POST /sites/{site_id}/services/{addon}/instances][%d] createServiceInstance default  %+v", o._statusCode, o.Payload)
}

func (o *CreateServiceInstanceDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateServiceInstanceDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
