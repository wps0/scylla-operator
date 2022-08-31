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

	models "github.com/scylladb/scylla-operator/pkg/scyllaclient/internal/agent/models"
)

// NewOperationsMovefileParams creates a new OperationsMovefileParams object
// with the default values initialized.
func NewOperationsMovefileParams() *OperationsMovefileParams {
	var ()
	return &OperationsMovefileParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewOperationsMovefileParamsWithTimeout creates a new OperationsMovefileParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewOperationsMovefileParamsWithTimeout(timeout time.Duration) *OperationsMovefileParams {
	var ()
	return &OperationsMovefileParams{

		timeout: timeout,
	}
}

// NewOperationsMovefileParamsWithContext creates a new OperationsMovefileParams object
// with the default values initialized, and the ability to set a context for a request
func NewOperationsMovefileParamsWithContext(ctx context.Context) *OperationsMovefileParams {
	var ()
	return &OperationsMovefileParams{

		Context: ctx,
	}
}

// NewOperationsMovefileParamsWithHTTPClient creates a new OperationsMovefileParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewOperationsMovefileParamsWithHTTPClient(client *http.Client) *OperationsMovefileParams {
	var ()
	return &OperationsMovefileParams{
		HTTPClient: client,
	}
}

/*
OperationsMovefileParams contains all the parameters to send to the API endpoint
for the operations movefile operation typically these are written to a http.Request
*/
type OperationsMovefileParams struct {

	/*Group
	  Place this operation under this stat group

	*/
	Group string
	/*Copyfile
	  copyfile

	*/
	Copyfile *models.MoveOrCopyFileOptions

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the operations movefile params
func (o *OperationsMovefileParams) WithTimeout(timeout time.Duration) *OperationsMovefileParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the operations movefile params
func (o *OperationsMovefileParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the operations movefile params
func (o *OperationsMovefileParams) WithContext(ctx context.Context) *OperationsMovefileParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the operations movefile params
func (o *OperationsMovefileParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the operations movefile params
func (o *OperationsMovefileParams) WithHTTPClient(client *http.Client) *OperationsMovefileParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the operations movefile params
func (o *OperationsMovefileParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithGroup adds the group to the operations movefile params
func (o *OperationsMovefileParams) WithGroup(group string) *OperationsMovefileParams {
	o.SetGroup(group)
	return o
}

// SetGroup adds the group to the operations movefile params
func (o *OperationsMovefileParams) SetGroup(group string) {
	o.Group = group
}

// WithCopyfile adds the copyfile to the operations movefile params
func (o *OperationsMovefileParams) WithCopyfile(copyfile *models.MoveOrCopyFileOptions) *OperationsMovefileParams {
	o.SetCopyfile(copyfile)
	return o
}

// SetCopyfile adds the copyfile to the operations movefile params
func (o *OperationsMovefileParams) SetCopyfile(copyfile *models.MoveOrCopyFileOptions) {
	o.Copyfile = copyfile
}

// WriteToRequest writes these params to a swagger request
func (o *OperationsMovefileParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param _group
	qrGroup := o.Group
	qGroup := qrGroup
	if qGroup != "" {
		if err := r.SetQueryParam("_group", qGroup); err != nil {
			return err
		}
	}

	if o.Copyfile != nil {
		if err := r.SetBodyParam(o.Copyfile); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
