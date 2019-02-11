package user

import (
	"github.com/go-openapi/runtime/middleware"
	"trade-shop/pkg/restapi/operations/user"
)

func (c *Context) GetInventoryList(params user.InventoryParams) middleware.Responder {
	usr, ok := c.auth.GetUserAuth(params.HTTPRequest)
	if !ok {
		return user.NewInventoryUnauthorized()
	}

	inv, _ := c.st.GetInventoryByUserId(usr)

	return user.NewInventoryOK().WithPayload(inv)
}
