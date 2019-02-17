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

// ErrorResult Ответ API с ошибкой
// swagger:model errorResult
type ErrorResult struct {

	// error
	// Required: true
	Error *ErrorResultError `json:"error"`
}

// Validate validates this error result
func (m *ErrorResult) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateError(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ErrorResult) validateError(formats strfmt.Registry) error {

	if err := validate.Required("error", "body", m.Error); err != nil {
		return err
	}

	if m.Error != nil {
		if err := m.Error.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("error")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ErrorResult) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ErrorResult) UnmarshalBinary(b []byte) error {
	var res ErrorResult
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ErrorResultError error result error
// swagger:model ErrorResultError
type ErrorResultError struct {

	// HTTP-статус ответа
	// Required: true
	Code string `json:"code"`

	// Человекопонятное описание ошибки
	// Required: true
	Description string `json:"description"`
}

// Validate validates this error result error
func (m *ErrorResultError) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ErrorResultError) validateCode(formats strfmt.Registry) error {

	if err := validate.RequiredString("error"+"."+"code", "body", string(m.Code)); err != nil {
		return err
	}

	return nil
}

func (m *ErrorResultError) validateDescription(formats strfmt.Registry) error {

	if err := validate.RequiredString("error"+"."+"description", "body", string(m.Description)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ErrorResultError) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ErrorResultError) UnmarshalBinary(b []byte) error {
	var res ErrorResultError
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}