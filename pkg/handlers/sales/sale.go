package sales

import (
	"trade-shop/pkg/httperrors"
	"trade-shop/pkg/restapi/operations/sales"

	"github.com/go-openapi/runtime/middleware"
)

func (c *Context) SaleItems(params sales.SaleParams) middleware.Responder {
	userID, ok := c.auth.GetUserAuth(params.HTTPRequest)
	if !ok {
		return sales.NewSaleUnauthorized()
	}

	err := c.sale.CreateLot(*userID, params.Body)
	if err != nil && err.Error() == "not enough items" {
		return sales.NewSaleBadRequest().WithPayload(&httperrors.NotEnoughItems)
	}

	return sales.NewSaleOK()
}
