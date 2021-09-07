// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Invite invite
//
// swagger:model Invite
type Invite struct {

	// created at
	// Required: true
	CreatedAt *int64 `json:"created_at"`

	// email
	// Required: true
	Email *string `json:"email"`

	// expires at
	// Required: true
	ExpiresAt *int64 `json:"expires_at"`

	// Map from group id to name.
	Groups map[string]string `json:"groups,omitempty"`

	// id
	// Required: true
	ID *string `json:"id"`

	// resource type
	ResourceType string `json:"resource_type,omitempty"`

	// status
	// Required: true
	Status *string `json:"status"`

	// updated at
	UpdatedAt int64 `json:"updated_at,omitempty"`
}

// Validate validates this invite
func (m *Invite) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEmail(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExpiresAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Invite) validateCreatedAt(formats strfmt.Registry) error {

	if err := validate.Required("created_at", "body", m.CreatedAt); err != nil {
		return err
	}

	return nil
}

func (m *Invite) validateEmail(formats strfmt.Registry) error {

	if err := validate.Required("email", "body", m.Email); err != nil {
		return err
	}

	return nil
}

func (m *Invite) validateExpiresAt(formats strfmt.Registry) error {

	if err := validate.Required("expires_at", "body", m.ExpiresAt); err != nil {
		return err
	}

	return nil
}

func (m *Invite) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *Invite) validateStatus(formats strfmt.Registry) error {

	if err := validate.Required("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Invite) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Invite) UnmarshalBinary(b []byte) error {
	var res Invite
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}