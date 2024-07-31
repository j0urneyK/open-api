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

// CreateTicketReader is a Reader for the CreateTicket structure.
type CreateTicketReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateTicketReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateTicketCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCreateTicketDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateTicketCreated creates a CreateTicketCreated with default headers values
func NewCreateTicketCreated() *CreateTicketCreated {
	return &CreateTicketCreated{}
}

/*
CreateTicketCreated handles this case with default header values.

ok
*/
type CreateTicketCreated struct {
	Payload *models.Ticket
}

func (o *CreateTicketCreated) Error() string {
	return fmt.Sprintf("[POST /oauth/tickets][%d] createTicketCreated  %+v", 201, o.Payload)
}

func (o *CreateTicketCreated) GetPayload() *models.Ticket {
	return o.Payload
}

func (o *CreateTicketCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Ticket)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateTicketDefault creates a CreateTicketDefault with default headers values
func NewCreateTicketDefault(code int) *CreateTicketDefault {
	return &CreateTicketDefault{
		_statusCode: code,
	}
}

/*
CreateTicketDefault handles this case with default header values.

error
*/
type CreateTicketDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the create ticket default response
func (o *CreateTicketDefault) Code() int {
	return o._statusCode
}

func (o *CreateTicketDefault) Error() string {
	return fmt.Sprintf("[POST /oauth/tickets][%d] createTicket default  %+v", o._statusCode, o.Payload)
}

func (o *CreateTicketDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateTicketDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
