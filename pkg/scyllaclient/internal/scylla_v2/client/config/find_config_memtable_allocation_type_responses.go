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

// FindConfigMemtableAllocationTypeReader is a Reader for the FindConfigMemtableAllocationType structure.
type FindConfigMemtableAllocationTypeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FindConfigMemtableAllocationTypeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFindConfigMemtableAllocationTypeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewFindConfigMemtableAllocationTypeDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewFindConfigMemtableAllocationTypeOK creates a FindConfigMemtableAllocationTypeOK with default headers values
func NewFindConfigMemtableAllocationTypeOK() *FindConfigMemtableAllocationTypeOK {
	return &FindConfigMemtableAllocationTypeOK{}
}

/*
FindConfigMemtableAllocationTypeOK handles this case with default header values.

Config value
*/
type FindConfigMemtableAllocationTypeOK struct {
	Payload string
}

func (o *FindConfigMemtableAllocationTypeOK) GetPayload() string {
	return o.Payload
}

func (o *FindConfigMemtableAllocationTypeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFindConfigMemtableAllocationTypeDefault creates a FindConfigMemtableAllocationTypeDefault with default headers values
func NewFindConfigMemtableAllocationTypeDefault(code int) *FindConfigMemtableAllocationTypeDefault {
	return &FindConfigMemtableAllocationTypeDefault{
		_statusCode: code,
	}
}

/*
FindConfigMemtableAllocationTypeDefault handles this case with default header values.

unexpected error
*/
type FindConfigMemtableAllocationTypeDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the find config memtable allocation type default response
func (o *FindConfigMemtableAllocationTypeDefault) Code() int {
	return o._statusCode
}

func (o *FindConfigMemtableAllocationTypeDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *FindConfigMemtableAllocationTypeDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *FindConfigMemtableAllocationTypeDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}
