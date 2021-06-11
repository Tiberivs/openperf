// Code generated by go-swagger; DO NOT EDIT.

package memory_generator

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/spirent/openperf/api/client/golang/models"
)

// CreateMemoryGeneratorReader is a Reader for the CreateMemoryGenerator structure.
type CreateMemoryGeneratorReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateMemoryGeneratorReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateMemoryGeneratorCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateMemoryGeneratorCreated creates a CreateMemoryGeneratorCreated with default headers values
func NewCreateMemoryGeneratorCreated() *CreateMemoryGeneratorCreated {
	return &CreateMemoryGeneratorCreated{}
}

/* CreateMemoryGeneratorCreated describes a response with status code 201, with default header values.

Created
*/
type CreateMemoryGeneratorCreated struct {

	/* URI of created memory generator
	 */
	Location string

	Payload *models.MemoryGenerator
}

func (o *CreateMemoryGeneratorCreated) Error() string {
	return fmt.Sprintf("[POST /memory-generators][%d] createMemoryGeneratorCreated  %+v", 201, o.Payload)
}
func (o *CreateMemoryGeneratorCreated) GetPayload() *models.MemoryGenerator {
	return o.Payload
}

func (o *CreateMemoryGeneratorCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Location
	hdrLocation := response.GetHeader("Location")

	if hdrLocation != "" {
		o.Location = hdrLocation
	}

	o.Payload = new(models.MemoryGenerator)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}