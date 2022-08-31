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

// CacheServiceInvalidateCounterCachePostReader is a Reader for the CacheServiceInvalidateCounterCachePost structure.
type CacheServiceInvalidateCounterCachePostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CacheServiceInvalidateCounterCachePostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCacheServiceInvalidateCounterCachePostOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCacheServiceInvalidateCounterCachePostDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCacheServiceInvalidateCounterCachePostOK creates a CacheServiceInvalidateCounterCachePostOK with default headers values
func NewCacheServiceInvalidateCounterCachePostOK() *CacheServiceInvalidateCounterCachePostOK {
	return &CacheServiceInvalidateCounterCachePostOK{}
}

/*
CacheServiceInvalidateCounterCachePostOK handles this case with default header values.

CacheServiceInvalidateCounterCachePostOK cache service invalidate counter cache post o k
*/
type CacheServiceInvalidateCounterCachePostOK struct {
}

func (o *CacheServiceInvalidateCounterCachePostOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCacheServiceInvalidateCounterCachePostDefault creates a CacheServiceInvalidateCounterCachePostDefault with default headers values
func NewCacheServiceInvalidateCounterCachePostDefault(code int) *CacheServiceInvalidateCounterCachePostDefault {
	return &CacheServiceInvalidateCounterCachePostDefault{
		_statusCode: code,
	}
}

/*
CacheServiceInvalidateCounterCachePostDefault handles this case with default header values.

internal server error
*/
type CacheServiceInvalidateCounterCachePostDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the cache service invalidate counter cache post default response
func (o *CacheServiceInvalidateCounterCachePostDefault) Code() int {
	return o._statusCode
}

func (o *CacheServiceInvalidateCounterCachePostDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *CacheServiceInvalidateCounterCachePostDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *CacheServiceInvalidateCounterCachePostDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}
