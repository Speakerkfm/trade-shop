package user

import (
	"fmt"
	"trade-shop/pkg/restapi/operations/user"

	"github.com/go-openapi/runtime/middleware"
)

func (c *Context) GetInventoryList(params user.InventoryParams) middleware.Responder {
	userID, ok := c.auth.GetUserAuth(params.HTTPRequest)
	if !ok {
		return user.NewInventoryUnauthorized()
	}

	items := c.inv.MakeInventory(*userID)
	bill := fmt.Sprintf("%.2f", c.st.GetUserBill(*userID))

	return user.NewInventoryOK().WithPayload(&user.InventoryOKBody{Bill: bill, Items: items})
}
