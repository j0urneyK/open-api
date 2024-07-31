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

// UpdateHookReader is a Reader for the UpdateHook structure.
type UpdateHookReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateHookReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateHookOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewUpdateHookDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateHookOK creates a UpdateHookOK with default headers values
func NewUpdateHookOK() *UpdateHookOK {
	return &UpdateHookOK{}
}

/*
UpdateHookOK handles this case with default header values.

OK
*/
type UpdateHookOK struct {
	Payload *models.Hook
}

func (o *UpdateHookOK) Error() string {
	return fmt.Sprintf("[PUT /hooks/{hook_id}][%d] updateHookOK  %+v", 200, o.Payload)
}

func (o *UpdateHookOK) GetPayload() *models.Hook {
	return o.Payload
}

func (o *UpdateHookOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Hook)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateHookDefault creates a UpdateHookDefault with default headers values
func NewUpdateHookDefault(code int) *UpdateHookDefault {
	return &UpdateHookDefault{
		_statusCode: code,
	}
}

/*
UpdateHookDefault handles this case with default header values.

error
*/
type UpdateHookDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the update hook default response
func (o *UpdateHookDefault) Code() int {
	return o._statusCode
}

func (o *UpdateHookDefault) Error() string {
	return fmt.Sprintf("[PUT /hooks/{hook_id}][%d] updateHook default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateHookDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateHookDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
