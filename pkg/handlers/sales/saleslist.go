package sales

import (
	"github.com/go-openapi/runtime/middleware"
	"trade-shop/pkg/restapi/operations/sales"
	"trade-shop/pkg/service/serviceiface"
	"trade-shop/pkg/store"
)

type Context struct {
	st    store.StoreInterface
	sales serviceiface.SalesList
}

func NewContext(st store.StoreInterface, sales serviceiface.SalesList) *Context {
	return &Context{st: st, sales: sales}
}

func (c *Context) GetSalesList(params sales.SalesListParams) middleware.Responder {
	return sales.NewSalesListOK().WithPayload(c.sales.GetSalesList())
}
