package sales

import (
	"github.com/go-openapi/runtime/middleware"
	"trade-shop/pkg/restapi/operations/sales"
)

func (c *Context) GetSalesList(params sales.SalesListParams) middleware.Responder {
	if _, ok := c.auth.GetUserAuth(params.HTTPRequest); !ok {
		return sales.NewSalesListUnauthorized()
	}
	return sales.NewSalesListOK().WithPayload(c.sale.GetSalesListJSON())
}
