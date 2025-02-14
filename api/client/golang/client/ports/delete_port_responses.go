// Code generated by go-swagger; DO NOT EDIT.

package ports

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeletePortReader is a Reader for the DeletePort structure.
type DeletePortReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeletePortReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeletePortNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeletePortNoContent creates a DeletePortNoContent with default headers values
func NewDeletePortNoContent() *DeletePortNoContent {
	return &DeletePortNoContent{}
}

/* DeletePortNoContent describes a response with status code 204, with default header values.

No Content
*/
type DeletePortNoContent struct {
}

func (o *DeletePortNoContent) Error() string {
	return fmt.Sprintf("[DELETE /ports/{id}][%d] deletePortNoContent ", 204)
}

func (o *DeletePortNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
