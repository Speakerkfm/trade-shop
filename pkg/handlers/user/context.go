package user

import (
	"trade-shop/pkg/service/serviceiface"
	"trade-shop/pkg/store"

	redistore "gopkg.in/boj/redistore.v1"
)

type Context struct {
	st   store.StoreInterface
	rst  *redistore.RediStore
	auth serviceiface.AuthService
	sale serviceiface.Sale
	inv  serviceiface.Inventory
}

func NewContext(st store.StoreInterface, rst *redistore.RediStore, auth serviceiface.AuthService, sale serviceiface.Sale, inv serviceiface.Inventory) *Context {
	return &Context{st: st, rst: rst, auth: auth, sale: sale, inv: inv}
}
