package user

import (
	"net/http"
	"trade-shop/pkg/restapi/operations/user"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/gorilla/sessions"
)

type logoutSessionWriter struct {
	r *http.Request
}

func (ts *logoutSessionWriter) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {
	err := sessions.Save(ts.r, rw)
	if err != nil {
		panic(err)
	}

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

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

	return &logoutSessionWriter{r: params.HTTPRequest}
}
