// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/j0urneyK/open-api/v2/go/models"
)

// NewUpdateSiteMetadataParams creates a new UpdateSiteMetadataParams object
// with the default values initialized.
func NewUpdateSiteMetadataParams() *UpdateSiteMetadataParams {
	var ()
	return &UpdateSiteMetadataParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateSiteMetadataParamsWithTimeout creates a new UpdateSiteMetadataParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateSiteMetadataParamsWithTimeout(timeout time.Duration) *UpdateSiteMetadataParams {
	var ()
	return &UpdateSiteMetadataParams{

		timeout: timeout,
	}
}

// NewUpdateSiteMetadataParamsWithContext creates a new UpdateSiteMetadataParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateSiteMetadataParamsWithContext(ctx context.Context) *UpdateSiteMetadataParams {
	var ()
	return &UpdateSiteMetadataParams{

		Context: ctx,
	}
}

// NewUpdateSiteMetadataParamsWithHTTPClient creates a new UpdateSiteMetadataParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateSiteMetadataParamsWithHTTPClient(client *http.Client) *UpdateSiteMetadataParams {
	var ()
	return &UpdateSiteMetadataParams{
		HTTPClient: client,
	}
}

/*
UpdateSiteMetadataParams contains all the parameters to send to the API endpoint
for the update site metadata operation typically these are written to a http.Request
*/
type UpdateSiteMetadataParams struct {

	/*Metadata*/
	Metadata models.Metadata
	/*SiteID*/
	SiteID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update site metadata params
func (o *UpdateSiteMetadataParams) WithTimeout(timeout time.Duration) *UpdateSiteMetadataParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update site metadata params
func (o *UpdateSiteMetadataParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update site metadata params
func (o *UpdateSiteMetadataParams) WithContext(ctx context.Context) *UpdateSiteMetadataParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update site metadata params
func (o *UpdateSiteMetadataParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update site metadata params
func (o *UpdateSiteMetadataParams) WithHTTPClient(client *http.Client) *UpdateSiteMetadataParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update site metadata params
func (o *UpdateSiteMetadataParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithMetadata adds the metadata to the update site metadata params
func (o *UpdateSiteMetadataParams) WithMetadata(metadata models.Metadata) *UpdateSiteMetadataParams {
	o.SetMetadata(metadata)
	return o
}

// SetMetadata adds the metadata to the update site metadata params
func (o *UpdateSiteMetadataParams) SetMetadata(metadata models.Metadata) {
	o.Metadata = metadata
}

// WithSiteID adds the siteID to the update site metadata params
func (o *UpdateSiteMetadataParams) WithSiteID(siteID string) *UpdateSiteMetadataParams {
	o.SetSiteID(siteID)
	return o
}

// SetSiteID adds the siteId to the update site metadata params
func (o *UpdateSiteMetadataParams) SetSiteID(siteID string) {
	o.SiteID = siteID
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateSiteMetadataParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Metadata != nil {
		if err := r.SetBodyParam(o.Metadata); err != nil {
			return err
		}
	}

	// path param site_id
	if err := r.SetPathParam("site_id", o.SiteID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
