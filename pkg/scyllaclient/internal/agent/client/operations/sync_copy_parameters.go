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

// NewSyncCopyParams creates a new SyncCopyParams object
// with the default values initialized.
func NewSyncCopyParams() *SyncCopyParams {
	var (
		asyncDefault = bool(true)
	)
	return &SyncCopyParams{
		Async: asyncDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewSyncCopyParamsWithTimeout creates a new SyncCopyParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSyncCopyParamsWithTimeout(timeout time.Duration) *SyncCopyParams {
	var (
		asyncDefault = bool(true)
	)
	return &SyncCopyParams{
		Async: asyncDefault,

		timeout: timeout,
	}
}

// NewSyncCopyParamsWithContext creates a new SyncCopyParams object
// with the default values initialized, and the ability to set a context for a request
func NewSyncCopyParamsWithContext(ctx context.Context) *SyncCopyParams {
	var (
		asyncDefault = bool(true)
	)
	return &SyncCopyParams{
		Async: asyncDefault,

		Context: ctx,
	}
}

// NewSyncCopyParamsWithHTTPClient creates a new SyncCopyParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSyncCopyParamsWithHTTPClient(client *http.Client) *SyncCopyParams {
	var (
		asyncDefault = bool(true)
	)
	return &SyncCopyParams{
		Async:      asyncDefault,
		HTTPClient: client,
	}
}

/*
SyncCopyParams contains all the parameters to send to the API endpoint
for the sync copy operation typically these are written to a http.Request
*/
type SyncCopyParams struct {

	/*Async
	  Async request

	*/
	Async bool
	/*Group
	  Place this operation under this stat group

	*/
	Group string
	/*Copydir
	  copydir

	*/
	Copydir SyncCopyBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the sync copy params
func (o *SyncCopyParams) WithTimeout(timeout time.Duration) *SyncCopyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the sync copy params
func (o *SyncCopyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the sync copy params
func (o *SyncCopyParams) WithContext(ctx context.Context) *SyncCopyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the sync copy params
func (o *SyncCopyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the sync copy params
func (o *SyncCopyParams) WithHTTPClient(client *http.Client) *SyncCopyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the sync copy params
func (o *SyncCopyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAsync adds the async to the sync copy params
func (o *SyncCopyParams) WithAsync(async bool) *SyncCopyParams {
	o.SetAsync(async)
	return o
}

// SetAsync adds the async to the sync copy params
func (o *SyncCopyParams) SetAsync(async bool) {
	o.Async = async
}

// WithGroup adds the group to the sync copy params
func (o *SyncCopyParams) WithGroup(group string) *SyncCopyParams {
	o.SetGroup(group)
	return o
}

// SetGroup adds the group to the sync copy params
func (o *SyncCopyParams) SetGroup(group string) {
	o.Group = group
}

// WithCopydir adds the copydir to the sync copy params
func (o *SyncCopyParams) WithCopydir(copydir SyncCopyBody) *SyncCopyParams {
	o.SetCopydir(copydir)
	return o
}

// SetCopydir adds the copydir to the sync copy params
func (o *SyncCopyParams) SetCopydir(copydir SyncCopyBody) {
	o.Copydir = copydir
}

// WriteToRequest writes these params to a swagger request
func (o *SyncCopyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param _async
	qrAsync := o.Async
	qAsync := swag.FormatBool(qrAsync)
	if qAsync != "" {
		if err := r.SetQueryParam("_async", qAsync); err != nil {
			return err
		}
	}

	// query param _group
	qrGroup := o.Group
	qGroup := qrGroup
	if qGroup != "" {
		if err := r.SetQueryParam("_group", qGroup); err != nil {
			return err
		}
	}

	if err := r.SetBodyParam(o.Copydir); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
