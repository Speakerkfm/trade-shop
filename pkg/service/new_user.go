package service

import (
	"fmt"
	"trade-shop/pkg/store"
)

type UserService struct {
	st store.StoreInterface
}

func NewUserService(st store.StoreInterface) *UserService {
	return &UserService{
		st: st,
	}
}

func (u *UserService) RegisterNewUser(email string, password string) error {
	if u.st.IsEmailTaken(email) {
		return fmt.Errorf("email is taken")
	}

	_, err := u.st.CreateNewUser(email, password)

	return err
}
