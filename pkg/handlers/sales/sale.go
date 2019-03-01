package sales

import (
	"github.com/go-openapi/runtime/middleware"
	"trade-shop/pkg/httperrors"
	"trade-shop/pkg/restapi/operations/sales"
)

func (c *Context) SaleItems(params sales.SaleParams) middleware.Responder {
	userID, ok := c.auth.GetUserAuth(params.HTTPRequest)
	if !ok {
		return sales.NewSaleUnauthorized()
	}

	if err := c.sale.CreateLot(*userID, params.Body); err != nil {
		return sales.NewSaleBadRequest().WithPayload(&httperrors.NotEnoughItems)
	}

	return sales.NewSaleOK()
}
