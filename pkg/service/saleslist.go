package service

import (
	"trade-shop/pkg/models"
	"trade-shop/pkg/store"
)

type SalesList struct {
	store store.StoreInterface
}

func NewSalesList(store *store.Store) *SalesList {
	sls := &SalesList{
		store: store,
	}

	return sls
}

func (sls *SalesList) GetSalesList() []*models.Sale {
	salesList, _ := sls.store.GetSaleItemList()

	var salesBody []*models.Sale
	var salesMap = make(map[int64]int, len(salesList))
	var count = 0

	for idx := range salesList {
		item := &models.SaleItemsItems0{
			ID:    salesList[idx].ItemID,
			Name:  salesList[idx].Name,
			Count: salesList[idx].Count,
			Price: salesList[idx].Price,
		}

		_, ok := salesMap[salesList[idx].SaleID]
		if !ok {
			salesBody = append(salesBody, &models.Sale{ID: salesList[idx].SaleID})
			salesMap[salesList[idx].SaleID] = count
			count++
		}

		val, _ := salesMap[salesList[idx].SaleID]
		salesBody[val].Items = append(salesBody[val].Items, item)
		salesBody[val].TotalCount += item.Price * float64(item.Count)
	}

	return salesBody
}
