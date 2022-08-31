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

// ColumnFamilyMetricsMemtableOnHeapSizeByNameGetReader is a Reader for the ColumnFamilyMetricsMemtableOnHeapSizeByNameGet structure.
type ColumnFamilyMetricsMemtableOnHeapSizeByNameGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ColumnFamilyMetricsMemtableOnHeapSizeByNameGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewColumnFamilyMetricsMemtableOnHeapSizeByNameGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewColumnFamilyMetricsMemtableOnHeapSizeByNameGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewColumnFamilyMetricsMemtableOnHeapSizeByNameGetOK creates a ColumnFamilyMetricsMemtableOnHeapSizeByNameGetOK with default headers values
func NewColumnFamilyMetricsMemtableOnHeapSizeByNameGetOK() *ColumnFamilyMetricsMemtableOnHeapSizeByNameGetOK {
	return &ColumnFamilyMetricsMemtableOnHeapSizeByNameGetOK{}
}

/*
ColumnFamilyMetricsMemtableOnHeapSizeByNameGetOK handles this case with default header values.

ColumnFamilyMetricsMemtableOnHeapSizeByNameGetOK column family metrics memtable on heap size by name get o k
*/
type ColumnFamilyMetricsMemtableOnHeapSizeByNameGetOK struct {
	Payload interface{}
}

func (o *ColumnFamilyMetricsMemtableOnHeapSizeByNameGetOK) GetPayload() interface{} {
	return o.Payload
}

func (o *ColumnFamilyMetricsMemtableOnHeapSizeByNameGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewColumnFamilyMetricsMemtableOnHeapSizeByNameGetDefault creates a ColumnFamilyMetricsMemtableOnHeapSizeByNameGetDefault with default headers values
func NewColumnFamilyMetricsMemtableOnHeapSizeByNameGetDefault(code int) *ColumnFamilyMetricsMemtableOnHeapSizeByNameGetDefault {
	return &ColumnFamilyMetricsMemtableOnHeapSizeByNameGetDefault{
		_statusCode: code,
	}
}

/*
ColumnFamilyMetricsMemtableOnHeapSizeByNameGetDefault handles this case with default header values.

internal server error
*/
type ColumnFamilyMetricsMemtableOnHeapSizeByNameGetDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the column family metrics memtable on heap size by name get default response
func (o *ColumnFamilyMetricsMemtableOnHeapSizeByNameGetDefault) Code() int {
	return o._statusCode
}

func (o *ColumnFamilyMetricsMemtableOnHeapSizeByNameGetDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *ColumnFamilyMetricsMemtableOnHeapSizeByNameGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *ColumnFamilyMetricsMemtableOnHeapSizeByNameGetDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}
