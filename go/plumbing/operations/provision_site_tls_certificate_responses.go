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

// ProvisionSiteTLSCertificateReader is a Reader for the ProvisionSiteTLSCertificate structure.
type ProvisionSiteTLSCertificateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ProvisionSiteTLSCertificateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewProvisionSiteTLSCertificateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewProvisionSiteTLSCertificateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewProvisionSiteTLSCertificateOK creates a ProvisionSiteTLSCertificateOK with default headers values
func NewProvisionSiteTLSCertificateOK() *ProvisionSiteTLSCertificateOK {
	return &ProvisionSiteTLSCertificateOK{}
}

/*
ProvisionSiteTLSCertificateOK handles this case with default header values.

OK
*/
type ProvisionSiteTLSCertificateOK struct {
	Payload *models.SniCertificate
}

func (o *ProvisionSiteTLSCertificateOK) Error() string {
	return fmt.Sprintf("[POST /sites/{site_id}/ssl][%d] provisionSiteTlsCertificateOK  %+v", 200, o.Payload)
}

func (o *ProvisionSiteTLSCertificateOK) GetPayload() *models.SniCertificate {
	return o.Payload
}

func (o *ProvisionSiteTLSCertificateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SniCertificate)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProvisionSiteTLSCertificateDefault creates a ProvisionSiteTLSCertificateDefault with default headers values
func NewProvisionSiteTLSCertificateDefault(code int) *ProvisionSiteTLSCertificateDefault {
	return &ProvisionSiteTLSCertificateDefault{
		_statusCode: code,
	}
}

/*
ProvisionSiteTLSCertificateDefault handles this case with default header values.

error
*/
type ProvisionSiteTLSCertificateDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the provision site TLS certificate default response
func (o *ProvisionSiteTLSCertificateDefault) Code() int {
	return o._statusCode
}

func (o *ProvisionSiteTLSCertificateDefault) Error() string {
	return fmt.Sprintf("[POST /sites/{site_id}/ssl][%d] provisionSiteTLSCertificate default  %+v", o._statusCode, o.Payload)
}

func (o *ProvisionSiteTLSCertificateDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *ProvisionSiteTLSCertificateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
