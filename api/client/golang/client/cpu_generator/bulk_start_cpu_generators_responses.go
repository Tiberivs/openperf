// Code generated by go-swagger; DO NOT EDIT.

package cpu_generator

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/spirent/openperf/api/client/golang/models"
)

// BulkStartCPUGeneratorsReader is a Reader for the BulkStartCPUGenerators structure.
type BulkStartCPUGeneratorsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *BulkStartCPUGeneratorsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewBulkStartCPUGeneratorsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewBulkStartCPUGeneratorsOK creates a BulkStartCPUGeneratorsOK with default headers values
func NewBulkStartCPUGeneratorsOK() *BulkStartCPUGeneratorsOK {
	return &BulkStartCPUGeneratorsOK{}
}

/* BulkStartCPUGeneratorsOK describes a response with status code 200, with default header values.

Success
*/
type BulkStartCPUGeneratorsOK struct {
	Payload []*models.CPUGeneratorResult
}

func (o *BulkStartCPUGeneratorsOK) Error() string {
	return fmt.Sprintf("[POST /cpu-generators/x/bulk-start][%d] bulkStartCpuGeneratorsOK  %+v", 200, o.Payload)
}
func (o *BulkStartCPUGeneratorsOK) GetPayload() []*models.CPUGeneratorResult {
	return o.Payload
}

func (o *BulkStartCPUGeneratorsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
