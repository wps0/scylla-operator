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
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewCollectdPostParams creates a new CollectdPostParams object
// with the default values initialized.
func NewCollectdPostParams() *CollectdPostParams {
	var ()
	return &CollectdPostParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCollectdPostParamsWithTimeout creates a new CollectdPostParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCollectdPostParamsWithTimeout(timeout time.Duration) *CollectdPostParams {
	var ()
	return &CollectdPostParams{

		timeout: timeout,
	}
}

// NewCollectdPostParamsWithContext creates a new CollectdPostParams object
// with the default values initialized, and the ability to set a context for a request
func NewCollectdPostParamsWithContext(ctx context.Context) *CollectdPostParams {
	var ()
	return &CollectdPostParams{

		Context: ctx,
	}
}

// NewCollectdPostParamsWithHTTPClient creates a new CollectdPostParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCollectdPostParamsWithHTTPClient(client *http.Client) *CollectdPostParams {
	var ()
	return &CollectdPostParams{
		HTTPClient: client,
	}
}

/*
CollectdPostParams contains all the parameters to send to the API endpoint
for the collectd post operation typically these are written to a http.Request
*/
type CollectdPostParams struct {

	/*Enable
	  set to true to enable all, anything else or omit to disable

	*/
	Enable *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the collectd post params
func (o *CollectdPostParams) WithTimeout(timeout time.Duration) *CollectdPostParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the collectd post params
func (o *CollectdPostParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the collectd post params
func (o *CollectdPostParams) WithContext(ctx context.Context) *CollectdPostParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the collectd post params
func (o *CollectdPostParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the collectd post params
func (o *CollectdPostParams) WithHTTPClient(client *http.Client) *CollectdPostParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the collectd post params
func (o *CollectdPostParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEnable adds the enable to the collectd post params
func (o *CollectdPostParams) WithEnable(enable *bool) *CollectdPostParams {
	o.SetEnable(enable)
	return o
}

// SetEnable adds the enable to the collectd post params
func (o *CollectdPostParams) SetEnable(enable *bool) {
	o.Enable = enable
}

// WriteToRequest writes these params to a swagger request
func (o *CollectdPostParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Enable != nil {

		// query param enable
		var qrEnable bool
		if o.Enable != nil {
			qrEnable = *o.Enable
		}
		qEnable := swag.FormatBool(qrEnable)
		if qEnable != "" {
			if err := r.SetQueryParam("enable", qEnable); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
