package sales

import (
	"github.com/go-openapi/runtime/middleware"
	"trade-shop/pkg/restapi/operations/sales"
)

func (c *Context) GetSalesList(params sales.SalesListParams) middleware.Responder {
	userID, ok := c.auth.GetUserAuth(params.HTTPRequest)
	if !ok {
		return sales.NewBuyUnauthorized()
	}

	return sales.NewSalesListOK().WithPayload(c.sale.MakeSalesList(*userID))
}
