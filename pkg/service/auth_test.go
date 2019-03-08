package service

import (
	"net/http"
	"testing"
	"trade-shop/pkg/mocks"

	"github.com/gorilla/sessions"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
)

func TestAuthService_GetUserAuth(t *testing.T) {
	rst := mocks.RediStore{}
	au := NewAuthService(&rst)
	uID, _ := uuid.NewV4()

	cases := map[string]struct {
		request *http.Request
		userID  *uuid.UUID
		session sessions.Session
		authed  bool
	}{
		"auth_ok": {
			request: &http.Request{},
			userID:  &uID,
			session: sessions.Session{Values: map[interface{}]interface{}{"userID": uID.String()}},
			authed:  true,
		},
		"auth_bad": {
			request: &http.Request{},
			userID:  nil,
			session: sessions.Session{Values: map[interface{}]interface{}{}},
			authed:  false,
		},
	}

	for _, test := range cases {
		rst.Mock.On("Get", test.request, "session-key").Return(&test.session, nil)
		userID, ok := au.GetUserAuth(test.request)

		assert.True(t, ok == test.authed)
		if test.authed {
			assert.True(t, *userID == *test.userID)
		}
	}
}
