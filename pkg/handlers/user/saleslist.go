package user

import (
	"github.com/go-openapi/runtime/middleware"
	"trade-shop/pkg/restapi/operations/user"
)

func (c *Context) GetUserSalesList(params user.UserSalesListParams) middleware.Responder {
	userID, ok := c.auth.GetUserAuth(params.HTTPRequest)
	if !ok {
		return user.NewUserSalesListUnauthorized()
	}

	return user.NewUserSalesListOK().WithPayload(c.sale.MakeUserSalesList(*userID))
}
