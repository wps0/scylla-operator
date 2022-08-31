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

// NewFindConfigLoadRingStateParams creates a new FindConfigLoadRingStateParams object
// with the default values initialized.
func NewFindConfigLoadRingStateParams() *FindConfigLoadRingStateParams {

	return &FindConfigLoadRingStateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewFindConfigLoadRingStateParamsWithTimeout creates a new FindConfigLoadRingStateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewFindConfigLoadRingStateParamsWithTimeout(timeout time.Duration) *FindConfigLoadRingStateParams {

	return &FindConfigLoadRingStateParams{

		timeout: timeout,
	}
}

// NewFindConfigLoadRingStateParamsWithContext creates a new FindConfigLoadRingStateParams object
// with the default values initialized, and the ability to set a context for a request
func NewFindConfigLoadRingStateParamsWithContext(ctx context.Context) *FindConfigLoadRingStateParams {

	return &FindConfigLoadRingStateParams{

		Context: ctx,
	}
}

// NewFindConfigLoadRingStateParamsWithHTTPClient creates a new FindConfigLoadRingStateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewFindConfigLoadRingStateParamsWithHTTPClient(client *http.Client) *FindConfigLoadRingStateParams {

	return &FindConfigLoadRingStateParams{
		HTTPClient: client,
	}
}

/*
FindConfigLoadRingStateParams contains all the parameters to send to the API endpoint
for the find config load ring state operation typically these are written to a http.Request
*/
type FindConfigLoadRingStateParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the find config load ring state params
func (o *FindConfigLoadRingStateParams) WithTimeout(timeout time.Duration) *FindConfigLoadRingStateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the find config load ring state params
func (o *FindConfigLoadRingStateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the find config load ring state params
func (o *FindConfigLoadRingStateParams) WithContext(ctx context.Context) *FindConfigLoadRingStateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the find config load ring state params
func (o *FindConfigLoadRingStateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the find config load ring state params
func (o *FindConfigLoadRingStateParams) WithHTTPClient(client *http.Client) *FindConfigLoadRingStateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the find config load ring state params
func (o *FindConfigLoadRingStateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *FindConfigLoadRingStateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
