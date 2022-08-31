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

// NewColumnFamilyMetricsCasPrepareByNameGetParams creates a new ColumnFamilyMetricsCasPrepareByNameGetParams object
// with the default values initialized.
func NewColumnFamilyMetricsCasPrepareByNameGetParams() *ColumnFamilyMetricsCasPrepareByNameGetParams {
	var ()
	return &ColumnFamilyMetricsCasPrepareByNameGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewColumnFamilyMetricsCasPrepareByNameGetParamsWithTimeout creates a new ColumnFamilyMetricsCasPrepareByNameGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewColumnFamilyMetricsCasPrepareByNameGetParamsWithTimeout(timeout time.Duration) *ColumnFamilyMetricsCasPrepareByNameGetParams {
	var ()
	return &ColumnFamilyMetricsCasPrepareByNameGetParams{

		timeout: timeout,
	}
}

// NewColumnFamilyMetricsCasPrepareByNameGetParamsWithContext creates a new ColumnFamilyMetricsCasPrepareByNameGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewColumnFamilyMetricsCasPrepareByNameGetParamsWithContext(ctx context.Context) *ColumnFamilyMetricsCasPrepareByNameGetParams {
	var ()
	return &ColumnFamilyMetricsCasPrepareByNameGetParams{

		Context: ctx,
	}
}

// NewColumnFamilyMetricsCasPrepareByNameGetParamsWithHTTPClient creates a new ColumnFamilyMetricsCasPrepareByNameGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewColumnFamilyMetricsCasPrepareByNameGetParamsWithHTTPClient(client *http.Client) *ColumnFamilyMetricsCasPrepareByNameGetParams {
	var ()
	return &ColumnFamilyMetricsCasPrepareByNameGetParams{
		HTTPClient: client,
	}
}

/*
ColumnFamilyMetricsCasPrepareByNameGetParams contains all the parameters to send to the API endpoint
for the column family metrics cas prepare by name get operation typically these are written to a http.Request
*/
type ColumnFamilyMetricsCasPrepareByNameGetParams struct {

	/*Name
	  The column family name in keyspace:name format

	*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the column family metrics cas prepare by name get params
func (o *ColumnFamilyMetricsCasPrepareByNameGetParams) WithTimeout(timeout time.Duration) *ColumnFamilyMetricsCasPrepareByNameGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the column family metrics cas prepare by name get params
func (o *ColumnFamilyMetricsCasPrepareByNameGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the column family metrics cas prepare by name get params
func (o *ColumnFamilyMetricsCasPrepareByNameGetParams) WithContext(ctx context.Context) *ColumnFamilyMetricsCasPrepareByNameGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the column family metrics cas prepare by name get params
func (o *ColumnFamilyMetricsCasPrepareByNameGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the column family metrics cas prepare by name get params
func (o *ColumnFamilyMetricsCasPrepareByNameGetParams) WithHTTPClient(client *http.Client) *ColumnFamilyMetricsCasPrepareByNameGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the column family metrics cas prepare by name get params
func (o *ColumnFamilyMetricsCasPrepareByNameGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the column family metrics cas prepare by name get params
func (o *ColumnFamilyMetricsCasPrepareByNameGetParams) WithName(name string) *ColumnFamilyMetricsCasPrepareByNameGetParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the column family metrics cas prepare by name get params
func (o *ColumnFamilyMetricsCasPrepareByNameGetParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *ColumnFamilyMetricsCasPrepareByNameGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
