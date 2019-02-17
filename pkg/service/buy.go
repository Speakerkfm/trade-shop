package service

import (
	"trade-shop/pkg/store"
)

type BuyService struct {
	st *store.Store
}

func NewBuyService(st *store.Store) *BuyService {
	bs := &BuyService{
		st: st,
	}

	return bs
}

//func (bs *BuyService) Purchase(userID uuid.UUID, saleID uuid.UUID) error {
//}
