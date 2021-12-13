// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// OPAIntegrationSettings o p a integration settings
//
// swagger:model OPAIntegrationSettings
type OPAIntegrationSettings struct {

	// Enabled is the flag for enabling OPA integration
	Enabled bool `json:"enabled,omitempty"`

	// Enable mutation
	ExperimentalEnableMutation bool `json:"experimentalEnableMutation,omitempty"`

	// WebhookTimeout is the timeout that is set for the gatekeeper validating webhook admission review calls.
	// By default 10 seconds.
	WebhookTimeoutSeconds int32 `json:"webhookTimeoutSeconds,omitempty"`

	// audit resources
	AuditResources *ResourceRequirements `json:"auditResources,omitempty"`

	// controller resources
	ControllerResources *ResourceRequirements `json:"controllerResources,omitempty"`
}

// Validate validates this o p a integration settings
func (m *OPAIntegrationSettings) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAuditResources(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateControllerResources(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OPAIntegrationSettings) validateAuditResources(formats strfmt.Registry) error {
	if swag.IsZero(m.AuditResources) { // not required
		return nil
	}

	if m.AuditResources != nil {
		if err := m.AuditResources.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("auditResources")
			}
			return err
		}
	}

	return nil
}

func (m *OPAIntegrationSettings) validateControllerResources(formats strfmt.Registry) error {
	if swag.IsZero(m.ControllerResources) { // not required
		return nil
	}

	if m.ControllerResources != nil {
		if err := m.ControllerResources.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("controllerResources")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this o p a integration settings based on the context it is used
func (m *OPAIntegrationSettings) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAuditResources(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateControllerResources(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OPAIntegrationSettings) contextValidateAuditResources(ctx context.Context, formats strfmt.Registry) error {

	if m.AuditResources != nil {
		if err := m.AuditResources.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("auditResources")
			}
			return err
		}
	}

	return nil
}

func (m *OPAIntegrationSettings) contextValidateControllerResources(ctx context.Context, formats strfmt.Registry) error {

	if m.ControllerResources != nil {
		if err := m.ControllerResources.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("controllerResources")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *OPAIntegrationSettings) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OPAIntegrationSettings) UnmarshalBinary(b []byte) error {
	var res OPAIntegrationSettings
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
