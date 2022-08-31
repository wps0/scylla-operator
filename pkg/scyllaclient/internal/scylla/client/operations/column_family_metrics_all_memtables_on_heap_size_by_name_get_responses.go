// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"strings"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/scylladb/scylla-operator/pkg/scyllaclient/internal/scylla/models"
)

// ColumnFamilyMetricsAllMemtablesOnHeapSizeByNameGetReader is a Reader for the ColumnFamilyMetricsAllMemtablesOnHeapSizeByNameGet structure.
type ColumnFamilyMetricsAllMemtablesOnHeapSizeByNameGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ColumnFamilyMetricsAllMemtablesOnHeapSizeByNameGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewColumnFamilyMetricsAllMemtablesOnHeapSizeByNameGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewColumnFamilyMetricsAllMemtablesOnHeapSizeByNameGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewColumnFamilyMetricsAllMemtablesOnHeapSizeByNameGetOK creates a ColumnFamilyMetricsAllMemtablesOnHeapSizeByNameGetOK with default headers values
func NewColumnFamilyMetricsAllMemtablesOnHeapSizeByNameGetOK() *ColumnFamilyMetricsAllMemtablesOnHeapSizeByNameGetOK {
	return &ColumnFamilyMetricsAllMemtablesOnHeapSizeByNameGetOK{}
}

/*
ColumnFamilyMetricsAllMemtablesOnHeapSizeByNameGetOK handles this case with default header values.

ColumnFamilyMetricsAllMemtablesOnHeapSizeByNameGetOK column family metrics all memtables on heap size by name get o k
*/
type ColumnFamilyMetricsAllMemtablesOnHeapSizeByNameGetOK struct {
	Payload interface{}
}

func (o *ColumnFamilyMetricsAllMemtablesOnHeapSizeByNameGetOK) GetPayload() interface{} {
	return o.Payload
}

func (o *ColumnFamilyMetricsAllMemtablesOnHeapSizeByNameGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewColumnFamilyMetricsAllMemtablesOnHeapSizeByNameGetDefault creates a ColumnFamilyMetricsAllMemtablesOnHeapSizeByNameGetDefault with default headers values
func NewColumnFamilyMetricsAllMemtablesOnHeapSizeByNameGetDefault(code int) *ColumnFamilyMetricsAllMemtablesOnHeapSizeByNameGetDefault {
	return &ColumnFamilyMetricsAllMemtablesOnHeapSizeByNameGetDefault{
		_statusCode: code,
	}
}

/*
ColumnFamilyMetricsAllMemtablesOnHeapSizeByNameGetDefault handles this case with default header values.

internal server error
*/
type ColumnFamilyMetricsAllMemtablesOnHeapSizeByNameGetDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the column family metrics all memtables on heap size by name get default response
func (o *ColumnFamilyMetricsAllMemtablesOnHeapSizeByNameGetDefault) Code() int {
	return o._statusCode
}

func (o *ColumnFamilyMetricsAllMemtablesOnHeapSizeByNameGetDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *ColumnFamilyMetricsAllMemtablesOnHeapSizeByNameGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *ColumnFamilyMetricsAllMemtablesOnHeapSizeByNameGetDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}
