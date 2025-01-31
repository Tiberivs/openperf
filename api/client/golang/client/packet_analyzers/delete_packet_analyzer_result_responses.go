// Code generated by go-swagger; DO NOT EDIT.

package packet_analyzers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeletePacketAnalyzerResultReader is a Reader for the DeletePacketAnalyzerResult structure.
type DeletePacketAnalyzerResultReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeletePacketAnalyzerResultReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeletePacketAnalyzerResultNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeletePacketAnalyzerResultNoContent creates a DeletePacketAnalyzerResultNoContent with default headers values
func NewDeletePacketAnalyzerResultNoContent() *DeletePacketAnalyzerResultNoContent {
	return &DeletePacketAnalyzerResultNoContent{}
}

/* DeletePacketAnalyzerResultNoContent describes a response with status code 204, with default header values.

No Content
*/
type DeletePacketAnalyzerResultNoContent struct {
}

func (o *DeletePacketAnalyzerResultNoContent) Error() string {
	return fmt.Sprintf("[DELETE /packet/analyzer-results/{id}][%d] deletePacketAnalyzerResultNoContent ", 204)
}

func (o *DeletePacketAnalyzerResultNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
