// Code generated by go-swagger; DO NOT EDIT.

package config

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewFindConfigReplaceAddressParams creates a new FindConfigReplaceAddressParams object
// with the default values initialized.
func NewFindConfigReplaceAddressParams() *FindConfigReplaceAddressParams {

	return &FindConfigReplaceAddressParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewFindConfigReplaceAddressParamsWithTimeout creates a new FindConfigReplaceAddressParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewFindConfigReplaceAddressParamsWithTimeout(timeout time.Duration) *FindConfigReplaceAddressParams {

	return &FindConfigReplaceAddressParams{

		timeout: timeout,
	}
}

// NewFindConfigReplaceAddressParamsWithContext creates a new FindConfigReplaceAddressParams object
// with the default values initialized, and the ability to set a context for a request
func NewFindConfigReplaceAddressParamsWithContext(ctx context.Context) *FindConfigReplaceAddressParams {

	return &FindConfigReplaceAddressParams{

		Context: ctx,
	}
}

// NewFindConfigReplaceAddressParamsWithHTTPClient creates a new FindConfigReplaceAddressParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewFindConfigReplaceAddressParamsWithHTTPClient(client *http.Client) *FindConfigReplaceAddressParams {

	return &FindConfigReplaceAddressParams{
		HTTPClient: client,
	}
}

/*
FindConfigReplaceAddressParams contains all the parameters to send to the API endpoint
for the find config replace address operation typically these are written to a http.Request
*/
type FindConfigReplaceAddressParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the find config replace address params
func (o *FindConfigReplaceAddressParams) WithTimeout(timeout time.Duration) *FindConfigReplaceAddressParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the find config replace address params
func (o *FindConfigReplaceAddressParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the find config replace address params
func (o *FindConfigReplaceAddressParams) WithContext(ctx context.Context) *FindConfigReplaceAddressParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the find config replace address params
func (o *FindConfigReplaceAddressParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the find config replace address params
func (o *FindConfigReplaceAddressParams) WithHTTPClient(client *http.Client) *FindConfigReplaceAddressParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the find config replace address params
func (o *FindConfigReplaceAddressParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *FindConfigReplaceAddressParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
