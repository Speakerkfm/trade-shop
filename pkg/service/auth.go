package service

import (
	"github.com/satori/go.uuid"
	"net/http"
	"trade-shop/pkg/service/serviceiface"
)

type AuthService struct {
	rst serviceiface.RediStore
}

func NewAuthService(rst serviceiface.RediStore) *AuthService {
	as := &AuthService{
		rst: rst,
	}

	return as
}

func (a *AuthService) GetUserAuth(r *http.Request) (*uuid.UUID, bool) {
	session, err := a.rst.Get(r, "session-key")
	if err != nil {
		panic(err)
	}

	if session.Values["userID"] == nil {
		return nil, false
	} else {
		id, err := uuid.FromString(session.Values["userID"].(string))
		if err != nil {
			panic(err)
		}

		return &id, true
	}
}
