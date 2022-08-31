// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/scylladb/scylla-operator/pkg/mermaidclient/internal/models"
)

// GetClusterClusterIDReader is a Reader for the GetClusterClusterID structure.
type GetClusterClusterIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetClusterClusterIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetClusterClusterIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetClusterClusterIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetClusterClusterIDOK creates a GetClusterClusterIDOK with default headers values
func NewGetClusterClusterIDOK() *GetClusterClusterIDOK {
	return &GetClusterClusterIDOK{}
}

/*
GetClusterClusterIDOK handles this case with default header values.

Cluster info
*/
type GetClusterClusterIDOK struct {
	Payload *models.Cluster
}

func (o *GetClusterClusterIDOK) Error() string {
	return fmt.Sprintf("[GET /cluster/{cluster_id}][%d] getClusterClusterIdOK  %+v", 200, o.Payload)
}

func (o *GetClusterClusterIDOK) GetPayload() *models.Cluster {
	return o.Payload
}

func (o *GetClusterClusterIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Cluster)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetClusterClusterIDDefault creates a GetClusterClusterIDDefault with default headers values
func NewGetClusterClusterIDDefault(code int) *GetClusterClusterIDDefault {
	return &GetClusterClusterIDDefault{
		_statusCode: code,
	}
}

/*
GetClusterClusterIDDefault handles this case with default header values.

Unexpected error
*/
type GetClusterClusterIDDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the get cluster cluster ID default response
func (o *GetClusterClusterIDDefault) Code() int {
	return o._statusCode
}

func (o *GetClusterClusterIDDefault) Error() string {
	return fmt.Sprintf("[GET /cluster/{cluster_id}][%d] GetClusterClusterID default  %+v", o._statusCode, o.Payload)
}

func (o *GetClusterClusterIDDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetClusterClusterIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
