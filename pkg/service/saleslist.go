package service

import (
	"github.com/go-openapi/strfmt"
	"github.com/satori/go.uuid"
	"trade-shop/pkg/models"
	"trade-shop/pkg/store"
)

type SalesListService struct {
	store store.StoreInterface
}

func NewSalesListService(store *store.Store) *SalesListService {
	sls := &SalesListService{
		store: store,
	}

	return sls
}

func (sls *SalesListService) GetSalesListJSON() []*models.Sale {
	salesList, _ := sls.store.GetSaleItemList()

	var salesBody []*models.Sale
	var salesMap = make(map[uuid.UUID]int, len(salesList))
	var count = 0

	for idx := range salesList {
		item := &models.SaleItemsItems0{
			ID:    strfmt.UUID(salesList[idx].ItemID.String()),
			Name:  salesList[idx].Name,
			Count: salesList[idx].Count,
			Price: salesList[idx].Price,
		}

		_, ok := salesMap[salesList[idx].SaleID]
		if !ok {
			salesBody = append(salesBody, &models.Sale{ID: strfmt.UUID(salesList[idx].ItemID.String())})
			salesMap[salesList[idx].SaleID] = count
			count++
		}

		val, _ := salesMap[salesList[idx].SaleID]
		salesBody[val].Items = append(salesBody[val].Items, item)
		salesBody[val].TotalCount += item.Price * float64(item.Count)
	}

	return salesBody
}
