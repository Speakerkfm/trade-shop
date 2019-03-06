package service

import (
	"github.com/go-openapi/strfmt"
	"github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
	"testing"
	"trade-shop/pkg/mocks"
	"trade-shop/pkg/models"
	"trade-shop/pkg/store"
)

func TestInventory_MakeInventory(t *testing.T) {
	stI := mocks.StoreInterface{}
	inv := NewInventory(&stI)
	uID, _ := uuid.NewV4()
	uID2, _ := uuid.NewV4()
	iID1, _ := uuid.NewV4()
	iID2, _ := uuid.NewV4()

	cases := map[string]struct {
		userID    uuid.UUID
		inventory []*store.Inventory
		items     []*models.Item
	}{
		"not_empty_inventory": {
			userID: uID,
			inventory: []*store.Inventory{
				{UserID: uID, ItemID: iID1, Name: "item1", Count: 12},
				{UserID: uID, ItemID: iID2, Name: "item2", Count: 5},
			},
			items: []*models.Item{
				{ID: strfmt.UUID(iID1.String()), Name: "item1", Count: 12},
				{ID: strfmt.UUID(iID2.String()), Name: "item2", Count: 5},
			},
		},
		"empty_inventory": {
			userID:    uID2,
			inventory: []*store.Inventory{},
			items:     nil,
		},
	}

	for _, test := range cases {
		stI.Mock.On("GetInventoryByUserId", test.userID).Return(test.inventory)

		items := inv.MakeInventory(test.userID)
		assert.Equal(t, test.items, items)
	}
}
