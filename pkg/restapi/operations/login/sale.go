// Code generated by go-swagger; DO NOT EDIT.

package login

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
	strfmt "github.com/go-openapi/strfmt"
	swag "github.com/go-openapi/swag"
)

// SaleHandlerFunc turns a function with the right signature into a sale handler
type SaleHandlerFunc func(SaleParams) middleware.Responder

// Handle executing the request and returning a response
func (fn SaleHandlerFunc) Handle(params SaleParams) middleware.Responder {
	return fn(params)
}

// SaleHandler interface for that can handle valid sale params
type SaleHandler interface {
	Handle(SaleParams) middleware.Responder
}

// NewSale creates a new http.Handler for the sale operation
func NewSale(ctx *middleware.Context, handler SaleHandler) *Sale {
	return &Sale{Context: ctx, Handler: handler}
}

/*Sale swagger:route POST /sale login sale

Продажа предмета пользователем

*/
type Sale struct {
	Context *middleware.Context
	Handler SaleHandler
}

func (o *Sale) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewSaleParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

// SaleBody Информация пользователя для входа в систему
// swagger:model SaleBody
type SaleBody struct {

	// email
	Email string `json:"email,omitempty"`

	// password
	Password string `json:"password,omitempty"`
}

// Validate validates this sale body
func (o *SaleBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *SaleBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *SaleBody) UnmarshalBinary(b []byte) error {
	var res SaleBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
