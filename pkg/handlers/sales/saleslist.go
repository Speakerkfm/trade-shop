package sales

import (
	"github.com/go-openapi/runtime/middleware"
	"trade-shop/pkg/restapi/operations/sales"
	"trade-shop/pkg/service/serviceiface"
	"trade-shop/pkg/store"
)

type Context struct {
	st    store.StoreInterface
	sales serviceiface.SalesListService
	auth  serviceiface.AuthService
}

func NewContext(st store.StoreInterface, sales serviceiface.SalesListService, auth serviceiface.AuthService) *Context {
	return &Context{st: st, sales: sales, auth: auth}
}

func (c *Context) GetSalesList(params sales.SalesListParams) middleware.Responder {
	if _, ok := c.auth.GetUserAuth(params.HTTPRequest); !ok {
		return sales.NewSalesListUnauthorized()
	}
	return sales.NewSalesListOK().WithPayload(c.sales.GetSalesListJSON())
}
