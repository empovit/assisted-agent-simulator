// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/empovit/assisted-agent-simulator/service/models"
)

// GetInstructionsReader is a Reader for the GetInstructions structure.
type GetInstructionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetInstructionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetInstructionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewGetInstructionsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetInstructionsOK creates a GetInstructionsOK with default headers values
func NewGetInstructionsOK() *GetInstructionsOK {
	return &GetInstructionsOK{}
}

/*GetInstructionsOK handles this case with default header values.

Success.
*/
type GetInstructionsOK struct {
	Payload *models.Step
}

func (o *GetInstructionsOK) Error() string {
	return fmt.Sprintf("[GET /instructions][%d] getInstructionsOK  %+v", 200, o.Payload)
}

func (o *GetInstructionsOK) GetPayload() *models.Step {
	return o.Payload
}

func (o *GetInstructionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Step)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInstructionsInternalServerError creates a GetInstructionsInternalServerError with default headers values
func NewGetInstructionsInternalServerError() *GetInstructionsInternalServerError {
	return &GetInstructionsInternalServerError{}
}

/*GetInstructionsInternalServerError handles this case with default header values.

Unexpected error
*/
type GetInstructionsInternalServerError struct {
}

func (o *GetInstructionsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /instructions][%d] getInstructionsInternalServerError ", 500)
}

func (o *GetInstructionsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
