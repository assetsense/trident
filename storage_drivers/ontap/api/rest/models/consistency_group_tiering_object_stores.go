// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ConsistencyGroupTieringObjectStores Object stores to use. Used for placement.
//
//
// swagger:model consistency_group_tiering_object_stores
type ConsistencyGroupTieringObjectStores []*ConsistencyGroupTieringObjectStoresItems0

// Validate validates this consistency group tiering object stores
func (m ConsistencyGroupTieringObjectStores) Validate(formats strfmt.Registry) error {
	var res []error

	iConsistencyGroupTieringObjectStoresSize := int64(len(m))

	if err := validate.MinItems("", "body", iConsistencyGroupTieringObjectStoresSize, 0); err != nil {
		return err
	}

	if err := validate.MaxItems("", "body", iConsistencyGroupTieringObjectStoresSize, 2); err != nil {
		return err
	}

	for i := 0; i < len(m); i++ {
		if swag.IsZero(m[i]) { // not required
			continue
		}

		if m[i] != nil {
			if err := m[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName(strconv.Itoa(i))
				}
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validate this consistency group tiering object stores based on the context it is used
func (m ConsistencyGroupTieringObjectStores) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	for i := 0; i < len(m); i++ {

		if m[i] != nil {
			if err := m[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName(strconv.Itoa(i))
				}
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ConsistencyGroupTieringObjectStoresItems0 consistency group tiering object stores items0
//
// swagger:model ConsistencyGroupTieringObjectStoresItems0
type ConsistencyGroupTieringObjectStoresItems0 struct {

	// The name of the object store to use. Used for placement.
	Name string `json:"name,omitempty"`
}

// Validate validates this consistency group tiering object stores items0
func (m *ConsistencyGroupTieringObjectStoresItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this consistency group tiering object stores items0 based on context it is used
func (m *ConsistencyGroupTieringObjectStoresItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ConsistencyGroupTieringObjectStoresItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConsistencyGroupTieringObjectStoresItems0) UnmarshalBinary(b []byte) error {
	var res ConsistencyGroupTieringObjectStoresItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}