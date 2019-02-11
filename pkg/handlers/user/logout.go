package user

import (
	"github.com/go-openapi/runtime/middleware"
	"trade-shop/pkg/restapi/operations/user"
)

func (c *Context) LogoutUser(params user.LogoutParams) middleware.Responder {
	session, err := c.rst.Get(params.HTTPRequest, "session-key")
	if err != nil {
		panic(err)
	}

	session.Options.MaxAge = -1

	_, ok := c.auth.GetUserAuth(params.HTTPRequest)
	if !ok {
		return user.NewLogoutUnauthorized()
	}

	return user.NewLogoutOK()
}
