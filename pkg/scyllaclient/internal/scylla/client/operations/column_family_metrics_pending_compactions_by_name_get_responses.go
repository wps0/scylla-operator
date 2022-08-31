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

// ColumnFamilyMetricsPendingCompactionsByNameGetReader is a Reader for the ColumnFamilyMetricsPendingCompactionsByNameGet structure.
type ColumnFamilyMetricsPendingCompactionsByNameGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ColumnFamilyMetricsPendingCompactionsByNameGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewColumnFamilyMetricsPendingCompactionsByNameGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewColumnFamilyMetricsPendingCompactionsByNameGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewColumnFamilyMetricsPendingCompactionsByNameGetOK creates a ColumnFamilyMetricsPendingCompactionsByNameGetOK with default headers values
func NewColumnFamilyMetricsPendingCompactionsByNameGetOK() *ColumnFamilyMetricsPendingCompactionsByNameGetOK {
	return &ColumnFamilyMetricsPendingCompactionsByNameGetOK{}
}

/*
ColumnFamilyMetricsPendingCompactionsByNameGetOK handles this case with default header values.

ColumnFamilyMetricsPendingCompactionsByNameGetOK column family metrics pending compactions by name get o k
*/
type ColumnFamilyMetricsPendingCompactionsByNameGetOK struct {
	Payload int32
}

func (o *ColumnFamilyMetricsPendingCompactionsByNameGetOK) GetPayload() int32 {
	return o.Payload
}

func (o *ColumnFamilyMetricsPendingCompactionsByNameGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewColumnFamilyMetricsPendingCompactionsByNameGetDefault creates a ColumnFamilyMetricsPendingCompactionsByNameGetDefault with default headers values
func NewColumnFamilyMetricsPendingCompactionsByNameGetDefault(code int) *ColumnFamilyMetricsPendingCompactionsByNameGetDefault {
	return &ColumnFamilyMetricsPendingCompactionsByNameGetDefault{
		_statusCode: code,
	}
}

/*
ColumnFamilyMetricsPendingCompactionsByNameGetDefault handles this case with default header values.

internal server error
*/
type ColumnFamilyMetricsPendingCompactionsByNameGetDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the column family metrics pending compactions by name get default response
func (o *ColumnFamilyMetricsPendingCompactionsByNameGetDefault) Code() int {
	return o._statusCode
}

func (o *ColumnFamilyMetricsPendingCompactionsByNameGetDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *ColumnFamilyMetricsPendingCompactionsByNameGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *ColumnFamilyMetricsPendingCompactionsByNameGetDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}
