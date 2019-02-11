package user

import (
	"gopkg.in/boj/redistore.v1"
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
