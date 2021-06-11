// Code generated by go-swagger; DO NOT EDIT.

package network_generator

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeleteNetworkGeneratorReader is a Reader for the DeleteNetworkGenerator structure.
type DeleteNetworkGeneratorReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteNetworkGeneratorReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteNetworkGeneratorNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteNetworkGeneratorNoContent creates a DeleteNetworkGeneratorNoContent with default headers values
func NewDeleteNetworkGeneratorNoContent() *DeleteNetworkGeneratorNoContent {
	return &DeleteNetworkGeneratorNoContent{}
}

/* DeleteNetworkGeneratorNoContent describes a response with status code 204, with default header values.

No Content
*/
type DeleteNetworkGeneratorNoContent struct {
}

func (o *DeleteNetworkGeneratorNoContent) Error() string {
	return fmt.Sprintf("[DELETE /network/generators/{id}][%d] deleteNetworkGeneratorNoContent ", 204)
}

func (o *DeleteNetworkGeneratorNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}