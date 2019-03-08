package service

import (
	"testing"
	"trade-shop/pkg/mocks"

	"github.com/stretchr/testify/assert"
)

func TestUserService_RegisterNewUser(t *testing.T) {
	stI := mocks.StoreInterface{}
	us := NewUserService(&stI)

	cases := map[string]struct {
		email    string
		password string
		isTaken  bool
		errText  string
	}{
		"auth_ok": {
			email:    "asdf@mail.com",
			password: "123456",
			isTaken:  false,
		},
		"auth_bad": {
			email:    "asdf12@mail.com",
			password: "123456",
			isTaken:  true,
			errText:  "email is taken",
		},
	}

	for _, test := range cases {
		stI.Mock.On("IsEmailTaken", test.email).Return(test.isTaken)

		if !test.isTaken {
			stI.Mock.On("CreateNewUser", test.email, test.password).Return(nil, nil)
		}

		err := us.RegisterNewUser(test.email, test.password)

		if test.isTaken {
			assert.True(t, err.Error() == test.errText)
		} else {
			assert.Nil(t, err)
		}
	}
}
