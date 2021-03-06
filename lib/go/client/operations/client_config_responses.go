// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/toggler-io/toggler/lib/go/models"
)

// ClientConfigReader is a Reader for the ClientConfig structure.
type ClientConfigReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ClientConfigReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewClientConfigOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewClientConfigBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewClientConfigInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewClientConfigOK creates a ClientConfigOK with default headers values
func NewClientConfigOK() *ClientConfigOK {
	return &ClientConfigOK{}
}

/*ClientConfigOK handles this case with default header values.

ClientConfigResponse returns information about the requester's rollout feature enrollment statuses.
*/
type ClientConfigOK struct {
	Payload *models.ClientConfigResponseBody
}

func (o *ClientConfigOK) Error() string {
	return fmt.Sprintf("[GET /client/config.json][%d] clientConfigOK  %+v", 200, o.Payload)
}

func (o *ClientConfigOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ClientConfigResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewClientConfigBadRequest creates a ClientConfigBadRequest with default headers values
func NewClientConfigBadRequest() *ClientConfigBadRequest {
	return &ClientConfigBadRequest{}
}

/*ClientConfigBadRequest handles this case with default header values.

ErrorResponse will contains a response about request that had some kind of problem.
The details will be included in the body.
*/
type ClientConfigBadRequest struct {
	Payload *models.ErrorResponseBody
}

func (o *ClientConfigBadRequest) Error() string {
	return fmt.Sprintf("[GET /client/config.json][%d] clientConfigBadRequest  %+v", 400, o.Payload)
}

func (o *ClientConfigBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewClientConfigInternalServerError creates a ClientConfigInternalServerError with default headers values
func NewClientConfigInternalServerError() *ClientConfigInternalServerError {
	return &ClientConfigInternalServerError{}
}

/*ClientConfigInternalServerError handles this case with default header values.

ErrorResponse will contains a response about request that had some kind of problem.
The details will be included in the body.
*/
type ClientConfigInternalServerError struct {
	Payload *models.ErrorResponseBody
}

func (o *ClientConfigInternalServerError) Error() string {
	return fmt.Sprintf("[GET /client/config.json][%d] clientConfigInternalServerError  %+v", 500, o.Payload)
}

func (o *ClientConfigInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*ClientConfigBody client config body
swagger:model ClientConfigBody
*/
type ClientConfigBody struct {

	// PilotExtID is the public uniq id that identify the caller pilot
	// Required: true
	PilotExtID *string `json:"id"`

	// ReleaseFlags are the list of private release flag name that should be matched against the pilot and state the enrollment for each.
	// Required: true
	ReleaseFlags []string `json:"release_flags"`
}

// Validate validates this client config body
func (o *ClientConfigBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validatePilotExtID(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateReleaseFlags(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ClientConfigBody) validatePilotExtID(formats strfmt.Registry) error {

	if err := validate.Required("Body"+"."+"id", "body", o.PilotExtID); err != nil {
		return err
	}

	return nil
}

func (o *ClientConfigBody) validateReleaseFlags(formats strfmt.Registry) error {

	if err := validate.Required("Body"+"."+"release_flags", "body", o.ReleaseFlags); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *ClientConfigBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ClientConfigBody) UnmarshalBinary(b []byte) error {
	var res ClientConfigBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
