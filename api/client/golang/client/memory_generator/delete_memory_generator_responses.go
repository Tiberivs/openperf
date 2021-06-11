// Code generated by go-swagger; DO NOT EDIT.

package memory_generator

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeleteMemoryGeneratorReader is a Reader for the DeleteMemoryGenerator structure.
type DeleteMemoryGeneratorReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteMemoryGeneratorReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteMemoryGeneratorNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteMemoryGeneratorNoContent creates a DeleteMemoryGeneratorNoContent with default headers values
func NewDeleteMemoryGeneratorNoContent() *DeleteMemoryGeneratorNoContent {
	return &DeleteMemoryGeneratorNoContent{}
}

/* DeleteMemoryGeneratorNoContent describes a response with status code 204, with default header values.

No Content
*/
type DeleteMemoryGeneratorNoContent struct {
}

func (o *DeleteMemoryGeneratorNoContent) Error() string {
	return fmt.Sprintf("[DELETE /memory-generators/{id}][%d] deleteMemoryGeneratorNoContent ", 204)
}

func (o *DeleteMemoryGeneratorNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}