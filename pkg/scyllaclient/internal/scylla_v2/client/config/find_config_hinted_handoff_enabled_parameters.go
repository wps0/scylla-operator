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

// NewFindConfigHintedHandoffEnabledParams creates a new FindConfigHintedHandoffEnabledParams object
// with the default values initialized.
func NewFindConfigHintedHandoffEnabledParams() *FindConfigHintedHandoffEnabledParams {

	return &FindConfigHintedHandoffEnabledParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewFindConfigHintedHandoffEnabledParamsWithTimeout creates a new FindConfigHintedHandoffEnabledParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewFindConfigHintedHandoffEnabledParamsWithTimeout(timeout time.Duration) *FindConfigHintedHandoffEnabledParams {

	return &FindConfigHintedHandoffEnabledParams{

		timeout: timeout,
	}
}

// NewFindConfigHintedHandoffEnabledParamsWithContext creates a new FindConfigHintedHandoffEnabledParams object
// with the default values initialized, and the ability to set a context for a request
func NewFindConfigHintedHandoffEnabledParamsWithContext(ctx context.Context) *FindConfigHintedHandoffEnabledParams {

	return &FindConfigHintedHandoffEnabledParams{

		Context: ctx,
	}
}

// NewFindConfigHintedHandoffEnabledParamsWithHTTPClient creates a new FindConfigHintedHandoffEnabledParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewFindConfigHintedHandoffEnabledParamsWithHTTPClient(client *http.Client) *FindConfigHintedHandoffEnabledParams {

	return &FindConfigHintedHandoffEnabledParams{
		HTTPClient: client,
	}
}

/*
FindConfigHintedHandoffEnabledParams contains all the parameters to send to the API endpoint
for the find config hinted handoff enabled operation typically these are written to a http.Request
*/
type FindConfigHintedHandoffEnabledParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the find config hinted handoff enabled params
func (o *FindConfigHintedHandoffEnabledParams) WithTimeout(timeout time.Duration) *FindConfigHintedHandoffEnabledParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the find config hinted handoff enabled params
func (o *FindConfigHintedHandoffEnabledParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the find config hinted handoff enabled params
func (o *FindConfigHintedHandoffEnabledParams) WithContext(ctx context.Context) *FindConfigHintedHandoffEnabledParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the find config hinted handoff enabled params
func (o *FindConfigHintedHandoffEnabledParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the find config hinted handoff enabled params
func (o *FindConfigHintedHandoffEnabledParams) WithHTTPClient(client *http.Client) *FindConfigHintedHandoffEnabledParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the find config hinted handoff enabled params
func (o *FindConfigHintedHandoffEnabledParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *FindConfigHintedHandoffEnabledParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
