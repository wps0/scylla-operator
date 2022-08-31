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

	strfmt "github.com/go-openapi/strfmt"
)

// NewColumnFamilyMetricsRecentBloomFilterFalsePositivesByNameGetParams creates a new ColumnFamilyMetricsRecentBloomFilterFalsePositivesByNameGetParams object
// with the default values initialized.
func NewColumnFamilyMetricsRecentBloomFilterFalsePositivesByNameGetParams() *ColumnFamilyMetricsRecentBloomFilterFalsePositivesByNameGetParams {
	var ()
	return &ColumnFamilyMetricsRecentBloomFilterFalsePositivesByNameGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewColumnFamilyMetricsRecentBloomFilterFalsePositivesByNameGetParamsWithTimeout creates a new ColumnFamilyMetricsRecentBloomFilterFalsePositivesByNameGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewColumnFamilyMetricsRecentBloomFilterFalsePositivesByNameGetParamsWithTimeout(timeout time.Duration) *ColumnFamilyMetricsRecentBloomFilterFalsePositivesByNameGetParams {
	var ()
	return &ColumnFamilyMetricsRecentBloomFilterFalsePositivesByNameGetParams{

		timeout: timeout,
	}
}

// NewColumnFamilyMetricsRecentBloomFilterFalsePositivesByNameGetParamsWithContext creates a new ColumnFamilyMetricsRecentBloomFilterFalsePositivesByNameGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewColumnFamilyMetricsRecentBloomFilterFalsePositivesByNameGetParamsWithContext(ctx context.Context) *ColumnFamilyMetricsRecentBloomFilterFalsePositivesByNameGetParams {
	var ()
	return &ColumnFamilyMetricsRecentBloomFilterFalsePositivesByNameGetParams{

		Context: ctx,
	}
}

// NewColumnFamilyMetricsRecentBloomFilterFalsePositivesByNameGetParamsWithHTTPClient creates a new ColumnFamilyMetricsRecentBloomFilterFalsePositivesByNameGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewColumnFamilyMetricsRecentBloomFilterFalsePositivesByNameGetParamsWithHTTPClient(client *http.Client) *ColumnFamilyMetricsRecentBloomFilterFalsePositivesByNameGetParams {
	var ()
	return &ColumnFamilyMetricsRecentBloomFilterFalsePositivesByNameGetParams{
		HTTPClient: client,
	}
}

/*
ColumnFamilyMetricsRecentBloomFilterFalsePositivesByNameGetParams contains all the parameters to send to the API endpoint
for the column family metrics recent bloom filter false positives by name get operation typically these are written to a http.Request
*/
type ColumnFamilyMetricsRecentBloomFilterFalsePositivesByNameGetParams struct {

	/*Name
	  The column family name in keyspace:name format

	*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the column family metrics recent bloom filter false positives by name get params
func (o *ColumnFamilyMetricsRecentBloomFilterFalsePositivesByNameGetParams) WithTimeout(timeout time.Duration) *ColumnFamilyMetricsRecentBloomFilterFalsePositivesByNameGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the column family metrics recent bloom filter false positives by name get params
func (o *ColumnFamilyMetricsRecentBloomFilterFalsePositivesByNameGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the column family metrics recent bloom filter false positives by name get params
func (o *ColumnFamilyMetricsRecentBloomFilterFalsePositivesByNameGetParams) WithContext(ctx context.Context) *ColumnFamilyMetricsRecentBloomFilterFalsePositivesByNameGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the column family metrics recent bloom filter false positives by name get params
func (o *ColumnFamilyMetricsRecentBloomFilterFalsePositivesByNameGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the column family metrics recent bloom filter false positives by name get params
func (o *ColumnFamilyMetricsRecentBloomFilterFalsePositivesByNameGetParams) WithHTTPClient(client *http.Client) *ColumnFamilyMetricsRecentBloomFilterFalsePositivesByNameGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the column family metrics recent bloom filter false positives by name get params
func (o *ColumnFamilyMetricsRecentBloomFilterFalsePositivesByNameGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the column family metrics recent bloom filter false positives by name get params
func (o *ColumnFamilyMetricsRecentBloomFilterFalsePositivesByNameGetParams) WithName(name string) *ColumnFamilyMetricsRecentBloomFilterFalsePositivesByNameGetParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the column family metrics recent bloom filter false positives by name get params
func (o *ColumnFamilyMetricsRecentBloomFilterFalsePositivesByNameGetParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *ColumnFamilyMetricsRecentBloomFilterFalsePositivesByNameGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
