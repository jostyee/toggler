// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ClientConfigRequest client config request
// swagger:model ClientConfigRequest
type ClientConfigRequest struct {

	// Features are the list of flag name that should be matched against the pilot and state the enrollment for each.
	// Required: true
	Features []string `json:"features"`

	// PilotID is the public uniq id that identify the caller pilot
	// Required: true
	PilotID *string `json:"id"`
}

// Validate validates this client config request
func (m *ClientConfigRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFeatures(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePilotID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClientConfigRequest) validateFeatures(formats strfmt.Registry) error {

	if err := validate.Required("features", "body", m.Features); err != nil {
		return err
	}

	return nil
}

func (m *ClientConfigRequest) validatePilotID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.PilotID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ClientConfigRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClientConfigRequest) UnmarshalBinary(b []byte) error {
	var res ClientConfigRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
