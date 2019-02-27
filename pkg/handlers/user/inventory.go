package user

import (
	"github.com/go-openapi/runtime/middleware"
	"trade-shop/pkg/restapi/operations/user"
)

func (c *Context) GetInventoryList(params user.InventoryParams) middleware.Responder {
	userID, ok := c.auth.GetUserAuth(params.HTTPRequest)
	if !ok {
		return user.NewInventoryUnauthorized()
	}

	items := c.inv.GetInventoryJSON(*userID)

	return user.NewInventoryOK().WithPayload(items)
}
