// Code generated by go-swagger; DO NOT EDIT.

package config

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"strings"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/scylladb/scylla-operator/pkg/scyllaclient/internal/scylla_v2/models"
)

// FindConfigCPUSchedulerReader is a Reader for the FindConfigCPUScheduler structure.
type FindConfigCPUSchedulerReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FindConfigCPUSchedulerReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFindConfigCPUSchedulerOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewFindConfigCPUSchedulerDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewFindConfigCPUSchedulerOK creates a FindConfigCPUSchedulerOK with default headers values
func NewFindConfigCPUSchedulerOK() *FindConfigCPUSchedulerOK {
	return &FindConfigCPUSchedulerOK{}
}

/*
FindConfigCPUSchedulerOK handles this case with default header values.

Config value
*/
type FindConfigCPUSchedulerOK struct {
	Payload bool
}

func (o *FindConfigCPUSchedulerOK) GetPayload() bool {
	return o.Payload
}

func (o *FindConfigCPUSchedulerOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFindConfigCPUSchedulerDefault creates a FindConfigCPUSchedulerDefault with default headers values
func NewFindConfigCPUSchedulerDefault(code int) *FindConfigCPUSchedulerDefault {
	return &FindConfigCPUSchedulerDefault{
		_statusCode: code,
	}
}

/*
FindConfigCPUSchedulerDefault handles this case with default header values.

unexpected error
*/
type FindConfigCPUSchedulerDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the find config cpu scheduler default response
func (o *FindConfigCPUSchedulerDefault) Code() int {
	return o._statusCode
}

func (o *FindConfigCPUSchedulerDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *FindConfigCPUSchedulerDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *FindConfigCPUSchedulerDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}
