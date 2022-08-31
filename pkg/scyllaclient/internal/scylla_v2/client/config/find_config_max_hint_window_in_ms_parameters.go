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

// NewFindConfigMaxHintWindowInMsParams creates a new FindConfigMaxHintWindowInMsParams object
// with the default values initialized.
func NewFindConfigMaxHintWindowInMsParams() *FindConfigMaxHintWindowInMsParams {

	return &FindConfigMaxHintWindowInMsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewFindConfigMaxHintWindowInMsParamsWithTimeout creates a new FindConfigMaxHintWindowInMsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewFindConfigMaxHintWindowInMsParamsWithTimeout(timeout time.Duration) *FindConfigMaxHintWindowInMsParams {

	return &FindConfigMaxHintWindowInMsParams{

		timeout: timeout,
	}
}

// NewFindConfigMaxHintWindowInMsParamsWithContext creates a new FindConfigMaxHintWindowInMsParams object
// with the default values initialized, and the ability to set a context for a request
func NewFindConfigMaxHintWindowInMsParamsWithContext(ctx context.Context) *FindConfigMaxHintWindowInMsParams {

	return &FindConfigMaxHintWindowInMsParams{

		Context: ctx,
	}
}

// NewFindConfigMaxHintWindowInMsParamsWithHTTPClient creates a new FindConfigMaxHintWindowInMsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewFindConfigMaxHintWindowInMsParamsWithHTTPClient(client *http.Client) *FindConfigMaxHintWindowInMsParams {

	return &FindConfigMaxHintWindowInMsParams{
		HTTPClient: client,
	}
}

/*
FindConfigMaxHintWindowInMsParams contains all the parameters to send to the API endpoint
for the find config max hint window in ms operation typically these are written to a http.Request
*/
type FindConfigMaxHintWindowInMsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the find config max hint window in ms params
func (o *FindConfigMaxHintWindowInMsParams) WithTimeout(timeout time.Duration) *FindConfigMaxHintWindowInMsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the find config max hint window in ms params
func (o *FindConfigMaxHintWindowInMsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the find config max hint window in ms params
func (o *FindConfigMaxHintWindowInMsParams) WithContext(ctx context.Context) *FindConfigMaxHintWindowInMsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the find config max hint window in ms params
func (o *FindConfigMaxHintWindowInMsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the find config max hint window in ms params
func (o *FindConfigMaxHintWindowInMsParams) WithHTTPClient(client *http.Client) *FindConfigMaxHintWindowInMsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the find config max hint window in ms params
func (o *FindConfigMaxHintWindowInMsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *FindConfigMaxHintWindowInMsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
