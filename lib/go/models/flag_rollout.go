// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// FlagRollout flag rollout
// swagger:model FlagRollout
type FlagRollout struct {

	// RandSeed allows you to configure the randomness for the percentage based pilot enrollment selection.
	// This value could have been neglected by using the flag name as random seed,
	// but that would reduce the flexibility for edge cases where you want
	// to use a similar pilot group as a successful flag rollout before.
	RandSeed int64 `json:"rand_seed_salt,omitempty"`

	// strategy
	Strategy *FlagRolloutStrategy `json:"strategy,omitempty"`
}

// Validate validates this flag rollout
func (m *FlagRollout) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateStrategy(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FlagRollout) validateStrategy(formats strfmt.Registry) error {

	if swag.IsZero(m.Strategy) { // not required
		return nil
	}

	if m.Strategy != nil {
		if err := m.Strategy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("strategy")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *FlagRollout) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FlagRollout) UnmarshalBinary(b []byte) error {
	var res FlagRollout
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
