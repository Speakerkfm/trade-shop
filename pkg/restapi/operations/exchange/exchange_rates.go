// Code generated by go-swagger; DO NOT EDIT.

package exchange

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// ExchangeRatesHandlerFunc turns a function with the right signature into a exchange rates handler
type ExchangeRatesHandlerFunc func(ExchangeRatesParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ExchangeRatesHandlerFunc) Handle(params ExchangeRatesParams) middleware.Responder {
	return fn(params)
}

// ExchangeRatesHandler interface for that can handle valid exchange rates params
type ExchangeRatesHandler interface {
	Handle(ExchangeRatesParams) middleware.Responder
}

// NewExchangeRates creates a new http.Handler for the exchange rates operation
func NewExchangeRates(ctx *middleware.Context, handler ExchangeRatesHandler) *ExchangeRates {
	return &ExchangeRates{Context: ctx, Handler: handler}
}

/*ExchangeRates swagger:route GET /exchange_rates exchange exchangeRates

Курсы обмена валют

*/
type ExchangeRates struct {
	Context *middleware.Context
	Handler ExchangeRatesHandler
}

func (o *ExchangeRates) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewExchangeRatesParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
