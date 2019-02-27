package sales

import (
	"github.com/go-openapi/runtime/middleware"
	"trade-shop/pkg/httperrors"
	"trade-shop/pkg/restapi/operations/sale"
)

func (c *Context) SaleItems(params sale.SaleParams) middleware.Responder {
	userID, ok := c.auth.GetUserAuth(params.HTTPRequest)
	if !ok {
		return sale.NewSaleUnauthorized()
	}

	if err := c.sale.CreateLot(*userID, params.Body); err != nil {
		return sale.NewSaleBadRequest().WithPayload(&httperrors.NotEnoughItems)
	}

	return sale.NewSaleOK()
}
