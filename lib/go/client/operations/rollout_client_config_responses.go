// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/toggler-io/toggler/lib/go/models"
)

// RolloutClientConfigReader is a Reader for the RolloutClientConfig structure.
type RolloutClientConfigReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RolloutClientConfigReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewRolloutClientConfigOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewRolloutClientConfigBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewRolloutClientConfigInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewRolloutClientConfigOK creates a RolloutClientConfigOK with default headers values
func NewRolloutClientConfigOK() *RolloutClientConfigOK {
	return &RolloutClientConfigOK{}
}

/*RolloutClientConfigOK handles this case with default header values.

RolloutClientConfigResponse returns information about the requester's rollout feature enrollment statuses.
*/
type RolloutClientConfigOK struct {
	Payload *models.RolloutClientConfigResponseBody
}

func (o *RolloutClientConfigOK) Error() string {
	return fmt.Sprintf("[POST /client/config.json][%d] rolloutClientConfigOK  %+v", 200, o.Payload)
}

func (o *RolloutClientConfigOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RolloutClientConfigResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRolloutClientConfigBadRequest creates a RolloutClientConfigBadRequest with default headers values
func NewRolloutClientConfigBadRequest() *RolloutClientConfigBadRequest {
	return &RolloutClientConfigBadRequest{}
}

/*RolloutClientConfigBadRequest handles this case with default header values.

ErrorResponse will contains a response about request that had some kind of problem.
The details will be included in the body.
*/
type RolloutClientConfigBadRequest struct {
	Payload *models.ErrorResponseBody
}

func (o *RolloutClientConfigBadRequest) Error() string {
	return fmt.Sprintf("[POST /client/config.json][%d] rolloutClientConfigBadRequest  %+v", 400, o.Payload)
}

func (o *RolloutClientConfigBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRolloutClientConfigInternalServerError creates a RolloutClientConfigInternalServerError with default headers values
func NewRolloutClientConfigInternalServerError() *RolloutClientConfigInternalServerError {
	return &RolloutClientConfigInternalServerError{}
}

/*RolloutClientConfigInternalServerError handles this case with default header values.

ErrorResponse will contains a response about request that had some kind of problem.
The details will be included in the body.
*/
type RolloutClientConfigInternalServerError struct {
	Payload *models.ErrorResponseBody
}

func (o *RolloutClientConfigInternalServerError) Error() string {
	return fmt.Sprintf("[POST /client/config.json][%d] rolloutClientConfigInternalServerError  %+v", 500, o.Payload)
}

func (o *RolloutClientConfigInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
