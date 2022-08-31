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

// ColumnFamilyMetricsCompressionMetadataOffHeapMemoryUsedGetReader is a Reader for the ColumnFamilyMetricsCompressionMetadataOffHeapMemoryUsedGet structure.
type ColumnFamilyMetricsCompressionMetadataOffHeapMemoryUsedGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ColumnFamilyMetricsCompressionMetadataOffHeapMemoryUsedGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewColumnFamilyMetricsCompressionMetadataOffHeapMemoryUsedGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewColumnFamilyMetricsCompressionMetadataOffHeapMemoryUsedGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewColumnFamilyMetricsCompressionMetadataOffHeapMemoryUsedGetOK creates a ColumnFamilyMetricsCompressionMetadataOffHeapMemoryUsedGetOK with default headers values
func NewColumnFamilyMetricsCompressionMetadataOffHeapMemoryUsedGetOK() *ColumnFamilyMetricsCompressionMetadataOffHeapMemoryUsedGetOK {
	return &ColumnFamilyMetricsCompressionMetadataOffHeapMemoryUsedGetOK{}
}

/*
ColumnFamilyMetricsCompressionMetadataOffHeapMemoryUsedGetOK handles this case with default header values.

ColumnFamilyMetricsCompressionMetadataOffHeapMemoryUsedGetOK column family metrics compression metadata off heap memory used get o k
*/
type ColumnFamilyMetricsCompressionMetadataOffHeapMemoryUsedGetOK struct {
	Payload interface{}
}

func (o *ColumnFamilyMetricsCompressionMetadataOffHeapMemoryUsedGetOK) GetPayload() interface{} {
	return o.Payload
}

func (o *ColumnFamilyMetricsCompressionMetadataOffHeapMemoryUsedGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewColumnFamilyMetricsCompressionMetadataOffHeapMemoryUsedGetDefault creates a ColumnFamilyMetricsCompressionMetadataOffHeapMemoryUsedGetDefault with default headers values
func NewColumnFamilyMetricsCompressionMetadataOffHeapMemoryUsedGetDefault(code int) *ColumnFamilyMetricsCompressionMetadataOffHeapMemoryUsedGetDefault {
	return &ColumnFamilyMetricsCompressionMetadataOffHeapMemoryUsedGetDefault{
		_statusCode: code,
	}
}

/*
ColumnFamilyMetricsCompressionMetadataOffHeapMemoryUsedGetDefault handles this case with default header values.

internal server error
*/
type ColumnFamilyMetricsCompressionMetadataOffHeapMemoryUsedGetDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the column family metrics compression metadata off heap memory used get default response
func (o *ColumnFamilyMetricsCompressionMetadataOffHeapMemoryUsedGetDefault) Code() int {
	return o._statusCode
}

func (o *ColumnFamilyMetricsCompressionMetadataOffHeapMemoryUsedGetDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *ColumnFamilyMetricsCompressionMetadataOffHeapMemoryUsedGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *ColumnFamilyMetricsCompressionMetadataOffHeapMemoryUsedGetDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}
