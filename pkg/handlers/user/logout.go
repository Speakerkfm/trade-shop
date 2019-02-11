package user

import (
	"github.com/go-openapi/runtime/middleware"
	"trade-shop/pkg/restapi/operations/user"
)

func (c *Context) LogoutUser(params user.LogoutParams) middleware.Responder {
	_, ok := c.auth.GetUserAuth(params.HTTPRequest)
	if !ok {
		return user.NewLogoutUnauthorized()
	}

	c.rst.Options.MaxAge = -1

	return user.NewLogoutOK()
}
