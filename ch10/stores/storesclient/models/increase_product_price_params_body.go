// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// IncreaseProductPriceParamsBody increase product price params body
//
// swagger:model increaseProductPriceParamsBody
type IncreaseProductPriceParamsBody struct {

	// price
	Price float64 `json:"price,omitempty"`
}

// Validate validates this increase product price params body
func (m *IncreaseProductPriceParamsBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this increase product price params body based on context it is used
func (m *IncreaseProductPriceParamsBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *IncreaseProductPriceParamsBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IncreaseProductPriceParamsBody) UnmarshalBinary(b []byte) error {
	var res IncreaseProductPriceParamsBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
