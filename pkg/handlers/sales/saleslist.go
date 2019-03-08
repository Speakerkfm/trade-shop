package sales

import (
	"trade-shop/pkg/restapi/operations/sales"

	"github.com/go-openapi/runtime/middleware"
)

func (c *Context) GetSalesList(params sales.SalesListParams) middleware.Responder {
	userID, ok := c.auth.GetUserAuth(params.HTTPRequest)
	if !ok {
		return sales.NewBuyUnauthorized()
	}

	return sales.NewSalesListOK().WithPayload(c.sale.MakeSalesList(*userID))
}
