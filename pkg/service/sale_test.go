package service

import (
	"fmt"
	"github.com/go-openapi/strfmt"
	"github.com/jinzhu/gorm"
	"github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
	"testing"
	"trade-shop/pkg/mocks"
	"trade-shop/pkg/models"
	"trade-shop/pkg/store"
)

func TestSale_CreateLot(t *testing.T) {
	stI := mocks.StoreInterface{}
	mailer := mocks.Mailer{}

	sl := &Sale{&stI, &mailer}

	sID1, _ := uuid.NewV4()
	uID1, _ := uuid.NewV4()
	sID2, _ := uuid.NewV4()
	uID2, _ := uuid.NewV4()
	iID1, _ := uuid.NewV4()
	iID2, _ := uuid.NewV4()
	errNotEnough := fmt.Errorf("not enough items")

	cases := map[string]struct {
		userID   uuid.UUID
		saleID   uuid.UUID
		tx       *gorm.DB
		itemList []*models.ItemSale
		err      []error
	}{
		"sale with 2 items": {
			userID: uID1,
			saleID: sID1,
			itemList: []*models.ItemSale{
				{ID: strfmt.UUID(iID1.String()), Name: "item1", Count: 3, Price: 12.20},
				{ID: strfmt.UUID(iID2.String()), Name: "item2", Count: 5, Price: 5.20},
			},
			tx:  &gorm.DB{},
			err: []error{nil, nil},
		},
		"sale with not enough items": {
			userID: uID2,
			saleID: sID2,
			itemList: []*models.ItemSale{
				{ID: strfmt.UUID(iID1.String()), Name: "item1", Count: 300, Price: 12.20},
				{ID: strfmt.UUID(iID2.String()), Name: "item2", Count: 5, Price: 5.20},
			},
			tx:  &gorm.DB{},
			err: []error{errNotEnough, nil},
		},
	}

	var wasError error

	for _, test := range cases {
		stI.On("ClearInventoryCache", test.userID)
		stI.On("CreateTransaction").Return(test.tx)
		stI.On("CreateNewSale", test.tx, test.userID).Return(test.saleID)
		wasError = nil
		for idx, item := range test.itemList {
			if wasError == nil {
				stI.On("RemoveItemFromUser", test.tx, test.userID, uuid.FromStringOrNil(item.ID.String()), item.Count).Return(test.err[idx])
				wasError = test.err[idx]
			}

			if wasError == nil {
				stI.On("AddItemToSale", test.tx, test.saleID, item).Return(nil)
				wasError = test.err[idx]
			}
		}

		if wasError == nil {
			stI.On("CommitTransaction", test.tx)
		} else {
			stI.On("RollbackTransaction", test.tx)
		}

		err := sl.CreateLot(test.userID, test.itemList)
		assert.True(t, err == wasError)
	}
}

func TestSale_Purchase(t *testing.T) {
	stI := mocks.StoreInterface{}
	mailer := mocks.Mailer{}

	sl := NewSale(&stI, &mailer)

	sID, _ := uuid.NewV4()
	slID, _ := uuid.NewV4()
	uID, _ := uuid.NewV4()
	iID1, _ := uuid.NewV4()
	iID2, _ := uuid.NewV4()

	cases := map[string]struct {
		userID   uuid.UUID
		user     *store.User
		sellerID uuid.UUID
		saleID   uuid.UUID
		tx       *gorm.DB
		itemList []*store.ItemSale
		money    float64
		err      error
	}{
		"ok": {
			userID:   uID,
			sellerID: slID,
			saleID:   sID,
			user: &store.User{
				Email: "asdf@mail.com",
			},
			itemList: []*store.ItemSale{
				{SellerID: sID, SaleID: slID, ItemID: iID1, Name: "item1", Count: 3, Price: 12.20},
				{SellerID: sID, SaleID: slID, ItemID: iID2, Name: "item2", Count: 5, Price: 5.20},
			},
			tx:    &gorm.DB{},
			money: 62.6,
			err:   nil,
		},
	}

	for _, test := range cases {
		stI.On("UserByUserID", test.userID).Return(test.user, true)
		stI.On("ClearInventoryCache", test.userID)
		stI.On("ClearInventoryCache", test.sellerID)
		stI.On("GetItemsInSaleBySaleID", test.saleID).Return(test.itemList, nil)
		stI.On("CreateTransaction").Return(test.tx)
		for _, item := range test.itemList {
			stI.On("AddItemToUser", test.tx, test.userID, item).Return(nil)
		}

		stI.On("RemoveMoneyFromUser", test.tx, test.userID, test.money).Return(nil)
		stI.On("AddMoneyToUser", test.tx, test.sellerID, test.money).Return(nil)
		stI.On("DeleteItemsInSale", test.tx, test.saleID).Return(nil)
		stI.On("DeleteSaleBySaleID", test.tx, test.saleID).Return(nil)
		mailer.On("SendNotificationEmail", test.user.Email, test.itemList).Return(nil)

		if test.err == nil {
			stI.On("CommitTransaction", test.tx)
		} else {
			stI.On("RollbackTransaction", test.tx)
		}

		err := sl.Purchase(test.userID, test.sellerID, test.saleID)
		assert.True(t, err == test.err)
	}
}

func TestSale_MakeSalesList(t *testing.T) {
	stI := mocks.StoreInterface{}
	mailer := mocks.Mailer{}

	sl := NewSale(&stI, &mailer)

	sID, _ := uuid.NewV4()
	slID, _ := uuid.NewV4()
	uID, _ := uuid.NewV4()
	iID1, _ := uuid.NewV4()
	iID2, _ := uuid.NewV4()

	cases := map[string]struct {
		userID     uuid.UUID
		salesList  []*store.ItemSale
		resultList []*models.Sale
	}{
		"ok": {
			userID: uID,
			salesList: []*store.ItemSale{
				{SellerID: sID, SaleID: slID, ItemID: iID1, Name: "item1", Count: 3, Price: 12.20},
				{SellerID: sID, SaleID: slID, ItemID: iID2, Name: "item2", Count: 5, Price: 5.20},
			},
			resultList: []*models.Sale{
				{
					ID: strfmt.UUID(slID.String()),
					Items: []*models.SaleItemsItems0{
						{ID: strfmt.UUID(iID1.String()), Name: "item1", Count: 3, Price: 12.20},
						{ID: strfmt.UUID(iID2.String()), Name: "item2", Count: 5, Price: 5.20},
					},
					TotalCount: 62.6,
				},
			},
		},
	}

	for _, test := range cases {
		stI.On("GetSaleItemList", test.userID).Return(test.salesList, nil)

		result := sl.MakeSalesList(test.userID)
		assert.Equal(t, test.resultList, result)
	}
}
