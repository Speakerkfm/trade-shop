package exchange

import (
	"trade-shop/pkg/restapi/operations/exchange"
	"trade-shop/pkg/service/serviceiface"

	"github.com/go-openapi/runtime/middleware"
)

type Context struct {
	e    serviceiface.ExchangeService
	auth serviceiface.AuthService
}

func NewContext(e serviceiface.ExchangeService, auth serviceiface.AuthService) *Context {
	return &Context{e: e, auth: auth}
}

func (c *Context) GetExchangeRates(params exchange.ExchangeRatesParams) middleware.Responder {
	_, ok := c.auth.GetUserAuth(params.HTTPRequest)
	if !ok {
		return exchange.NewExchangeRatesUnauthorized()
	}

	return exchange.NewExchangeRatesOK().WithPayload(c.e.CallRates())
}
