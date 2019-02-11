package inventory

import (
	"github.com/go-openapi/runtime/middleware"
	"trade-shop/pkg/restapi/operations/inventory"
	"trade-shop/pkg/service/serviceiface"
	"trade-shop/pkg/store"
)

type Context struct {
	st   store.StoreInterface
	auth serviceiface.AuthService
}

func NewContext(st store.StoreInterface, auth serviceiface.AuthService) *Context {
	return &Context{st: st, auth: auth}
}

func (c *Context) GetInventoryList(params inventory.InventoryParams) middleware.Responder {
	user, ok := c.auth.GetUserAuth(params.HTTPRequest)
	if !ok {
		return inventory.NewInventoryUnauthorized()
	}

	inv, _ := c.st.GetInventoryByUserId(user)

	return inventory.NewInventoryOK().WithPayload(inv)
}
