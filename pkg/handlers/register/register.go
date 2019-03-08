package register

import (
	"trade-shop/pkg/httperrors"
	"trade-shop/pkg/restapi/operations/register"
	"trade-shop/pkg/service/serviceiface"

	"github.com/go-openapi/runtime/middleware"
)

type Context struct {
	us serviceiface.UserService
}

func NewContext(us serviceiface.UserService) *Context {
	return &Context{us: us}
}

func (c *Context) Register(params register.RegisterParams) middleware.Responder {
	err := c.us.RegisterNewUser(params.Body.Email, params.Body.Password)

	if err != nil && err.Error() == "email is taken" {
		return register.NewRegisterBadRequest().WithPayload(&httperrors.EmailIsTaken)
	}

	return register.NewRegisterOK()
}
