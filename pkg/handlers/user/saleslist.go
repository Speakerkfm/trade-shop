package user

import (
	"trade-shop/pkg/restapi/operations/user"

	"github.com/go-openapi/runtime/middleware"
)

func (c *Context) GetUserSalesList(params user.UserSalesListParams) middleware.Responder {
	userID, ok := c.auth.GetUserAuth(params.HTTPRequest)
	if !ok {
		return user.NewUserSalesListUnauthorized()
	}

	return user.NewUserSalesListOK().WithPayload(c.sale.MakeUserSalesList(*userID))
}
