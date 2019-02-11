package login

import (
	"github.com/go-openapi/runtime/middleware"
	"gopkg.in/boj/redistore.v1"
	"trade-shop/pkg/restapi/operations/login"
	"trade-shop/pkg/service/serviceiface"
	"trade-shop/pkg/store"
)

type Context struct {
	st   store.StoreInterface
	rst  *redistore.RediStore
	auth serviceiface.AuthService
}

func NewContext(st store.StoreInterface, rst *redistore.RediStore, auth serviceiface.AuthService) *Context {
	return &Context{st: st, rst: rst, auth: auth}
}

func (c *Context) AuthByEmailAndPassword(params login.LoginParams) middleware.Responder {
	user, ok := c.st.UserByEmail(params.Body.Email)
	if !ok {
		return login.NewLoginUnauthorized()
	}

	session, err := c.rst.Get(params.HTTPRequest, "session-key")
	if err != nil {
		panic(err)
	}

	session.Values["userID"] = user.ID.String()

	return login.NewLoginFound().WithLocation("/user/inventory")
}
