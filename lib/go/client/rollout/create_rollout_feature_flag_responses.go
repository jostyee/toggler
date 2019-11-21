// Code generated by go-swagger; DO NOT EDIT.

package rollout

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/toggler-io/toggler/lib/go/models"
)

// CreateRolloutFeatureFlagReader is a Reader for the CreateRolloutFeatureFlag structure.
type CreateRolloutFeatureFlagReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateRolloutFeatureFlagReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateRolloutFeatureFlagOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateRolloutFeatureFlagBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateRolloutFeatureFlagInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateRolloutFeatureFlagOK creates a CreateRolloutFeatureFlagOK with default headers values
func NewCreateRolloutFeatureFlagOK() *CreateRolloutFeatureFlagOK {
	return &CreateRolloutFeatureFlagOK{}
}

/*CreateRolloutFeatureFlagOK handles this case with default header values.

CreateRolloutFeatureFlagResponse returns information about the requester's rollout feature enrollment status.
*/
type CreateRolloutFeatureFlagOK struct {
	Payload interface{}
}

func (o *CreateRolloutFeatureFlagOK) Error() string {
	return fmt.Sprintf("[POST /release/flag/create.json][%d] createRolloutFeatureFlagOK  %+v", 200, o.Payload)
}

func (o *CreateRolloutFeatureFlagOK) GetPayload() interface{} {
	return o.Payload
}

func (o *CreateRolloutFeatureFlagOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateRolloutFeatureFlagBadRequest creates a CreateRolloutFeatureFlagBadRequest with default headers values
func NewCreateRolloutFeatureFlagBadRequest() *CreateRolloutFeatureFlagBadRequest {
	return &CreateRolloutFeatureFlagBadRequest{}
}

/*CreateRolloutFeatureFlagBadRequest handles this case with default header values.

ErrorResponse will contains a response about request that had some kind of problem.
The details will be included in the body.
*/
type CreateRolloutFeatureFlagBadRequest struct {
	Payload *models.ErrorResponseBody
}

func (o *CreateRolloutFeatureFlagBadRequest) Error() string {
	return fmt.Sprintf("[POST /release/flag/create.json][%d] createRolloutFeatureFlagBadRequest  %+v", 400, o.Payload)
}

func (o *CreateRolloutFeatureFlagBadRequest) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *CreateRolloutFeatureFlagBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateRolloutFeatureFlagInternalServerError creates a CreateRolloutFeatureFlagInternalServerError with default headers values
func NewCreateRolloutFeatureFlagInternalServerError() *CreateRolloutFeatureFlagInternalServerError {
	return &CreateRolloutFeatureFlagInternalServerError{}
}

/*CreateRolloutFeatureFlagInternalServerError handles this case with default header values.

ErrorResponse will contains a response about request that had some kind of problem.
The details will be included in the body.
*/
type CreateRolloutFeatureFlagInternalServerError struct {
	Payload *models.ErrorResponseBody
}

func (o *CreateRolloutFeatureFlagInternalServerError) Error() string {
	return fmt.Sprintf("[POST /release/flag/create.json][%d] createRolloutFeatureFlagInternalServerError  %+v", 500, o.Payload)
}

func (o *CreateRolloutFeatureFlagInternalServerError) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *CreateRolloutFeatureFlagInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
