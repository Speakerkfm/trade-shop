//ifacemaker -f sale.go -s Sale -i Sale -p serviceiface -o serviceiface/sale.go
package service

import (
	"fmt"
	"trade-shop/pkg/models"
	"trade-shop/pkg/service/serviceiface"
	"trade-shop/pkg/store"

	"github.com/go-openapi/strfmt"
	uuid "github.com/satori/go.uuid"
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
			s.st.RollbackTransaction(tx)

			return err
		}

		if err := s.st.AddItemToSale(tx, saleID, val); err != nil {
			s.st.RollbackTransaction(tx)

			return err
		}
	}

	s.st.CommitTransaction(tx)

	return nil
}

func (s *Sale) Purchase(userID uuid.UUID, sellerID uuid.UUID, saleID uuid.UUID) error {
	_, found := s.st.UserByUserID(userID)
	if !found {
		return fmt.Errorf("user not found")
	}

	seller, found := s.st.UserByUserID(sellerID)
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
		money += val.Price * float64(val.Count)

		if err = s.st.AddItemToUser(tx, userID, val); err != nil {
			s.st.RollbackTransaction(tx)

			return err
		}
	}

	if money, err = formatFloat(money); err != nil {
		panic(err)
	}

	if err = s.st.RemoveMoneyFromUser(tx, userID, money); err != nil {
		s.st.RollbackTransaction(tx)

		return err
	}

	if err := s.st.AddMoneyToUser(tx, sellerID, money); err != nil {
		s.st.RollbackTransaction(tx)

		return err
	}

	if err := s.st.DeleteItemsInSale(tx, saleID); err != nil {
		s.st.RollbackTransaction(tx)

		return err
	}

	if err := s.st.DeleteSaleBySaleID(tx, saleID); err != nil {
		s.st.RollbackTransaction(tx)

		return err
	}

	s.st.CommitTransaction(tx)

	return s.mailer.SendNotificationEmail(seller.Email, itemList)
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

		val := salesMap[salesList[idx].SaleID]
		salesBody[val].Items = append(salesBody[val].Items, item)
		salesBody[val].TotalCount += item.Price * float64(item.Count)

		if salesBody[val].TotalCount, err = formatFloat(salesBody[val].TotalCount); err != nil {
			panic(err)
		}
	}

	return salesBody
}

func (s *Sale) MakeUserSalesList(userID uuid.UUID) []*models.Sale {
	var err error
	salesList, err := s.st.GetUserSaleItemList(userID)
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

		val := salesMap[salesList[idx].SaleID]
		salesBody[val].Items = append(salesBody[val].Items, item)
		salesBody[val].TotalCount += item.Price * float64(item.Count)

		if salesBody[val].TotalCount, err = formatFloat(salesBody[val].TotalCount); err != nil {
			panic(err)
		}
	}

	return salesBody
}

func (s *Sale) Cancel(userID uuid.UUID, saleID uuid.UUID) error {
	_, found := s.st.UserByUserID(userID)
	if !found {
		return fmt.Errorf("user not found")
	}

	s.st.ClearInventoryCache(userID)

	itemList, err := s.st.GetItemsInSaleBySaleID(saleID)
	if err != nil {
		return err
	}

	tx := s.st.CreateTransaction()
	money := 0.0

	for _, val := range itemList {
		money += val.Price * float64(val.Count)

		if err = s.st.AddItemToUser(tx, userID, val); err != nil {
			s.st.RollbackTransaction(tx)

			return err
		}
	}

	if money, err = formatFloat(money); err != nil {
		panic(err)
	}

	if err := s.st.AddMoneyToUser(tx, userID, money); err != nil {
		s.st.RollbackTransaction(tx)

		return err
	}

	if err := s.st.DeleteItemsInSale(tx, saleID); err != nil {
		s.st.RollbackTransaction(tx)

		return err
	}

	if err := s.st.DeleteSaleBySaleID(tx, saleID); err != nil {
		s.st.RollbackTransaction(tx)

		return err
	}

	s.st.CommitTransaction(tx)

	return nil
}
