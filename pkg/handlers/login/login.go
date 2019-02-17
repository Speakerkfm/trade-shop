package login

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/gorilla/sessions"
	"gopkg.in/boj/redistore.v1"
	"net/http"
	"trade-shop/pkg/httperrors"
	"trade-shop/pkg/restapi/operations/login"
	"trade-shop/pkg/service/serviceiface"
	"trade-shop/pkg/store"
)

type Context struct {
	st   store.StoreInterface
	rst  *redistore.RediStore
	auth serviceiface.AuthService
}

type loginSessionWriter struct {
	r *http.Request
}

func NewContext(st store.StoreInterface, rst *redistore.RediStore, auth serviceiface.AuthService) *Context {
	return &Context{st: st, rst: rst, auth: auth}
}

func (ts *loginSessionWriter) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {
	sessions.Save(ts.r, rw)

	rw.Header().Del(runtime.HeaderContentType)
	rw.Header().Set("Location", "/user/inventory")

	rw.WriteHeader(302)
}

func (c *Context) AuthByEmailAndPassword(params login.LoginParams) middleware.Responder {
	user, ok := c.st.UserByEmail(params.Body.Email)
	if !ok {
		return login.NewLoginUnauthorized().WithPayload(&httperrors.WrongUsernameOrPassword)
	}

	if !user.PasswordValid(params.Body.Password) {
		return login.NewLoginUnauthorized().WithPayload(&httperrors.WrongUsernameOrPassword)
	}

	session, err := c.rst.Get(params.HTTPRequest, "session-key")
	if err != nil {
		panic(err)
	}

	session.Values["userID"] = user.ID.String()

	return &loginSessionWriter{r: params.HTTPRequest}
}
