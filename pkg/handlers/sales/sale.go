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

	switch err := c.sale.CreateLot(*userID, params.Body); err.Error() {
	case "not enough items":
		return sales.NewSaleBadRequest().WithPayload(&httperrors.NotEnoughItems)
	default:
		return sales.NewSaleOK()
	}
}
