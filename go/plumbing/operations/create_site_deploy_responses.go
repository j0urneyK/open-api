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

// CreateSiteDeployReader is a Reader for the CreateSiteDeploy structure.
type CreateSiteDeployReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateSiteDeployReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateSiteDeployOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCreateSiteDeployDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateSiteDeployOK creates a CreateSiteDeployOK with default headers values
func NewCreateSiteDeployOK() *CreateSiteDeployOK {
	return &CreateSiteDeployOK{}
}

/*
CreateSiteDeployOK handles this case with default header values.

OK
*/
type CreateSiteDeployOK struct {
	Payload *models.Deploy
}

func (o *CreateSiteDeployOK) Error() string {
	return fmt.Sprintf("[POST /sites/{site_id}/deploys][%d] createSiteDeployOK  %+v", 200, o.Payload)
}

func (o *CreateSiteDeployOK) GetPayload() *models.Deploy {
	return o.Payload
}

func (o *CreateSiteDeployOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Deploy)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateSiteDeployDefault creates a CreateSiteDeployDefault with default headers values
func NewCreateSiteDeployDefault(code int) *CreateSiteDeployDefault {
	return &CreateSiteDeployDefault{
		_statusCode: code,
	}
}

/*
CreateSiteDeployDefault handles this case with default header values.

error
*/
type CreateSiteDeployDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the create site deploy default response
func (o *CreateSiteDeployDefault) Code() int {
	return o._statusCode
}

func (o *CreateSiteDeployDefault) Error() string {
	return fmt.Sprintf("[POST /sites/{site_id}/deploys][%d] createSiteDeploy default  %+v", o._statusCode, o.Payload)
}

func (o *CreateSiteDeployDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateSiteDeployDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
