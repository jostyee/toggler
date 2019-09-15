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

// IsFeatureGloballyEnabledRequestBody is feature globally enabled request body
// swagger:model IsFeatureGloballyEnabledRequestBody
type IsFeatureGloballyEnabledRequestBody struct {

	// Feature is the release flag name that is needed to be checked for enrollment
	// Required: true
	Feature *string `json:"feature"`
}

// Validate validates this is feature globally enabled request body
func (m *IsFeatureGloballyEnabledRequestBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFeature(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IsFeatureGloballyEnabledRequestBody) validateFeature(formats strfmt.Registry) error {

	if err := validate.Required("feature", "body", m.Feature); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IsFeatureGloballyEnabledRequestBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IsFeatureGloballyEnabledRequestBody) UnmarshalBinary(b []byte) error {
	var res IsFeatureGloballyEnabledRequestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
