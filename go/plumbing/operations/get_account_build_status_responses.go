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

// GetAccountBuildStatusReader is a Reader for the GetAccountBuildStatus structure.
type GetAccountBuildStatusReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAccountBuildStatusReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAccountBuildStatusOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetAccountBuildStatusDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetAccountBuildStatusOK creates a GetAccountBuildStatusOK with default headers values
func NewGetAccountBuildStatusOK() *GetAccountBuildStatusOK {
	return &GetAccountBuildStatusOK{}
}

/*
GetAccountBuildStatusOK handles this case with default header values.

OK
*/
type GetAccountBuildStatusOK struct {
	Payload []*models.BuildStatus
}

func (o *GetAccountBuildStatusOK) Error() string {
	return fmt.Sprintf("[GET /{account_id}/builds/status][%d] getAccountBuildStatusOK  %+v", 200, o.Payload)
}

func (o *GetAccountBuildStatusOK) GetPayload() []*models.BuildStatus {
	return o.Payload
}

func (o *GetAccountBuildStatusOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAccountBuildStatusDefault creates a GetAccountBuildStatusDefault with default headers values
func NewGetAccountBuildStatusDefault(code int) *GetAccountBuildStatusDefault {
	return &GetAccountBuildStatusDefault{
		_statusCode: code,
	}
}

/*
GetAccountBuildStatusDefault handles this case with default header values.

error
*/
type GetAccountBuildStatusDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get account build status default response
func (o *GetAccountBuildStatusDefault) Code() int {
	return o._statusCode
}

func (o *GetAccountBuildStatusDefault) Error() string {
	return fmt.Sprintf("[GET /{account_id}/builds/status][%d] getAccountBuildStatus default  %+v", o._statusCode, o.Payload)
}

func (o *GetAccountBuildStatusDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetAccountBuildStatusDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
