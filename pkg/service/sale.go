//ifacemaker -f sale.go -s Sale -i Sale -p serviceiface -o serviceiface/sale.go
package service

import (
	"fmt"
	"github.com/go-openapi/strfmt"
	"github.com/satori/go.uuid"
	"strconv"
	"trade-shop/pkg/models"
	"trade-shop/pkg/service/serviceiface"
	"trade-shop/pkg/store"
)

type Sale struct {
	st     store.StoreInterface
	mailer serviceiface.Mailer
}

func NewSale(st store.StoreInterface, mailer serviceiface.Mailer) *Sale {
	return &Sale{
		st:     st,
		mailer: mailer,
	}
}

func (s *Sale) CreateLot(userID uuid.UUID, itemList []*models.ItemSale) error {
	s.st.ClearInventoryCache(userID)

	tx := s.st.CreateTransaction()

	saleID := s.st.CreateNewSale(tx, userID)

	for _, val := range itemList {
		if err := s.st.RemoveItemFromUser(tx, userID, uuid.FromStringOrNil(val.ID.String()), val.Count); err != nil {
			tx.Rollback()

			return err
		}

		if err := s.st.AddItemToSale(tx, saleID, val); err != nil {
			tx.Rollback()

			return err
		}
	}

	tx.Commit()

	return nil
}

func (s *Sale) Purchase(userID uuid.UUID, sellerID uuid.UUID, saleID uuid.UUID) error {
	user, found := s.st.UserByUserID(userID)
	if !found {
		return fmt.Errorf("user not found")
	}

	s.st.ClearInventoryCache(userID)
	s.st.ClearInventoryCache(sellerID)

	itemList, err := s.st.GetItemsInSaleBySaleID(saleID)
	if err != nil {
		return err
	}

	tx := s.st.CreateTransaction()
	money := 0.0

	for _, val := range itemList {
		money = money + (val.Price * float64(val.Count))

		if err := s.st.AddItemToUser(tx, userID, val); err != nil {
			tx.Rollback()

			return err
		}
	}

	if money, err = strconv.ParseFloat(fmt.Sprintf("%.2f", money), 64); err != nil {
		panic(err)
	}

	if err := s.st.RemoveMoneyFromUser(tx, userID, money); err != nil {
		tx.Rollback()

		return err
	}

	if err := s.st.AddMoneyToUser(tx, sellerID, money); err != nil {
		tx.Rollback()

		return err
	}

	if err := s.st.DeleteItemsInSale(tx, saleID); err != nil {
		tx.Rollback()

		return err
	}

	if err := s.st.DeleteSaleBySaleID(tx, saleID); err != nil {
		tx.Rollback()

		return err
	}

	tx.Commit()

	return s.mailer.SendNotificationEmail(*user.Email, itemList)
}

func (s *Sale) MakeSalesList(userID uuid.UUID) []*models.Sale {
	var err error
	salesList, err := s.st.GetSaleItemList(userID)
	if err != nil {
		panic(err)
	}

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
			salesBody = append(salesBody, &models.Sale{ID: strfmt.UUID(salesList[idx].SaleID.String())})
			salesMap[salesList[idx].SaleID] = count
			count++
		}

		val, _ := salesMap[salesList[idx].SaleID]
		salesBody[val].Items = append(salesBody[val].Items, item)
		salesBody[val].TotalCount += item.Price * float64(item.Count)

		if salesBody[val].TotalCount, err = strconv.ParseFloat(fmt.Sprintf("%.2f", salesBody[val].TotalCount), 64); err != nil {
			panic(err)
		}
	}

	return salesBody
}
