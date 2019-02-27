package sales

import (
	"trade-shop/pkg/service/serviceiface"
	"trade-shop/pkg/store"
)

type Context struct {
	st   store.StoreInterface
	sale serviceiface.Sale
	auth serviceiface.AuthService
}

func NewContext(st store.StoreInterface, sale serviceiface.Sale, auth serviceiface.AuthService) *Context {
	return &Context{st: st, sale: sale, auth: auth}
}
