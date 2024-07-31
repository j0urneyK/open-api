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

// GetEnvVarReader is a Reader for the GetEnvVar structure.
type GetEnvVarReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetEnvVarReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetEnvVarOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetEnvVarDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetEnvVarOK creates a GetEnvVarOK with default headers values
func NewGetEnvVarOK() *GetEnvVarOK {
	return &GetEnvVarOK{}
}

/*
GetEnvVarOK handles this case with default header values.

OK
*/
type GetEnvVarOK struct {
	Payload *models.EnvVar
}

func (o *GetEnvVarOK) Error() string {
	return fmt.Sprintf("[GET /accounts/{account_id}/env/{key}][%d] getEnvVarOK  %+v", 200, o.Payload)
}

func (o *GetEnvVarOK) GetPayload() *models.EnvVar {
	return o.Payload
}

func (o *GetEnvVarOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EnvVar)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetEnvVarDefault creates a GetEnvVarDefault with default headers values
func NewGetEnvVarDefault(code int) *GetEnvVarDefault {
	return &GetEnvVarDefault{
		_statusCode: code,
	}
}

/*
GetEnvVarDefault handles this case with default header values.

error
*/
type GetEnvVarDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get env var default response
func (o *GetEnvVarDefault) Code() int {
	return o._statusCode
}

func (o *GetEnvVarDefault) Error() string {
	return fmt.Sprintf("[GET /accounts/{account_id}/env/{key}][%d] getEnvVar default  %+v", o._statusCode, o.Payload)
}

func (o *GetEnvVarDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetEnvVarDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
