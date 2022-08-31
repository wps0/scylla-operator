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

// NewFindConfigSstablePreemptiveOpenIntervalInMbParams creates a new FindConfigSstablePreemptiveOpenIntervalInMbParams object
// with the default values initialized.
func NewFindConfigSstablePreemptiveOpenIntervalInMbParams() *FindConfigSstablePreemptiveOpenIntervalInMbParams {

	return &FindConfigSstablePreemptiveOpenIntervalInMbParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewFindConfigSstablePreemptiveOpenIntervalInMbParamsWithTimeout creates a new FindConfigSstablePreemptiveOpenIntervalInMbParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewFindConfigSstablePreemptiveOpenIntervalInMbParamsWithTimeout(timeout time.Duration) *FindConfigSstablePreemptiveOpenIntervalInMbParams {

	return &FindConfigSstablePreemptiveOpenIntervalInMbParams{

		timeout: timeout,
	}
}

// NewFindConfigSstablePreemptiveOpenIntervalInMbParamsWithContext creates a new FindConfigSstablePreemptiveOpenIntervalInMbParams object
// with the default values initialized, and the ability to set a context for a request
func NewFindConfigSstablePreemptiveOpenIntervalInMbParamsWithContext(ctx context.Context) *FindConfigSstablePreemptiveOpenIntervalInMbParams {

	return &FindConfigSstablePreemptiveOpenIntervalInMbParams{

		Context: ctx,
	}
}

// NewFindConfigSstablePreemptiveOpenIntervalInMbParamsWithHTTPClient creates a new FindConfigSstablePreemptiveOpenIntervalInMbParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewFindConfigSstablePreemptiveOpenIntervalInMbParamsWithHTTPClient(client *http.Client) *FindConfigSstablePreemptiveOpenIntervalInMbParams {

	return &FindConfigSstablePreemptiveOpenIntervalInMbParams{
		HTTPClient: client,
	}
}

/*
FindConfigSstablePreemptiveOpenIntervalInMbParams contains all the parameters to send to the API endpoint
for the find config sstable preemptive open interval in mb operation typically these are written to a http.Request
*/
type FindConfigSstablePreemptiveOpenIntervalInMbParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the find config sstable preemptive open interval in mb params
func (o *FindConfigSstablePreemptiveOpenIntervalInMbParams) WithTimeout(timeout time.Duration) *FindConfigSstablePreemptiveOpenIntervalInMbParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the find config sstable preemptive open interval in mb params
func (o *FindConfigSstablePreemptiveOpenIntervalInMbParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the find config sstable preemptive open interval in mb params
func (o *FindConfigSstablePreemptiveOpenIntervalInMbParams) WithContext(ctx context.Context) *FindConfigSstablePreemptiveOpenIntervalInMbParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the find config sstable preemptive open interval in mb params
func (o *FindConfigSstablePreemptiveOpenIntervalInMbParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the find config sstable preemptive open interval in mb params
func (o *FindConfigSstablePreemptiveOpenIntervalInMbParams) WithHTTPClient(client *http.Client) *FindConfigSstablePreemptiveOpenIntervalInMbParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the find config sstable preemptive open interval in mb params
func (o *FindConfigSstablePreemptiveOpenIntervalInMbParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *FindConfigSstablePreemptiveOpenIntervalInMbParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
