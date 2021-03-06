// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"
)

// NewCancelParams creates a new CancelParams object
// no default values defined in spec.
func NewCancelParams() CancelParams {

	return CancelParams{}
}

// CancelParams contains all the bound params for the cancel operation
// typically these are obtained from a http.Request
//
// swagger:parameters cancel
type CancelParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*идентификатор продажи
	  Required: true
	  In: path
	*/
	SaleID strfmt.UUID
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewCancelParams() beforehand.
func (o *CancelParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rSaleID, rhkSaleID, _ := route.Params.GetOK("sale_id")
	if err := o.bindSaleID(rSaleID, rhkSaleID, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindSaleID binds and validates parameter SaleID from path.
func (o *CancelParams) bindSaleID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	// Format: uuid
	value, err := formats.Parse("uuid", raw)
	if err != nil {
		return errors.InvalidType("sale_id", "path", "strfmt.UUID", raw)
	}
	o.SaleID = *(value.(*strfmt.UUID))

	if err := o.validateSaleID(formats); err != nil {
		return err
	}

	return nil
}

// validateSaleID carries on validations for parameter SaleID
func (o *CancelParams) validateSaleID(formats strfmt.Registry) error {

	if err := validate.FormatOf("sale_id", "path", "uuid", o.SaleID.String(), formats); err != nil {
		return err
	}
	return nil
}
