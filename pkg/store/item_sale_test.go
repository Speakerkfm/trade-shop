package store

import (
	"testing"
	"trade-shop/pkg/models"

	"github.com/go-openapi/strfmt"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
)

func TestStore_GetSaleItemList(t *testing.T) {
	s := NewStore(Gorm, RedisClient)

	sID1, _ := uuid.NewV4()
	sID2, _ := uuid.NewV4()
	iID1, _ := uuid.NewV4()
	iID2, _ := uuid.NewV4()

	item1 := Item{
		ID:   iID1,
		Name: "item1",
	}
	item2 := Item{
		ID:   iID2,
		Name: "item2",
	}
	sale1 := Sales{
		UserID: sID1,
	}
	sale2 := Sales{
		UserID: sID2,
	}

	Gorm.Table("items").Create(&item1)
	Gorm.Table("items").Create(&item2)
	Gorm.Create(&sale1)
	Gorm.Create(&sale2)
	itemSale1 := ItemSale{SaleID: sale1.ID, ItemID: iID1, Count: 8, Price: 12.20}
	itemSale2 := ItemSale{SaleID: sale2.ID, ItemID: iID2, Count: 12, Price: 5.20}
	Gorm.Create(&itemSale1)
	Gorm.Create(&itemSale2)

	res, err := s.GetSaleItemList(sID1)
	assert.Nil(t, err)
	assert.True(t, res[0].Name == "item2")

	Gorm.Table("items").Delete(&item1)
	Gorm.Table("items").Delete(&item2)
	Gorm.Delete(&sale1)
	Gorm.Delete(&sale2)
	Gorm.Delete(&itemSale1)
	Gorm.Delete(&itemSale2)
}

func TestStore_GetItemsInSaleBySaleID(t *testing.T) {
	s := NewStore(Gorm, RedisClient)

	sID, _ := uuid.NewV4()
	iID1, _ := uuid.NewV4()
	iID2, _ := uuid.NewV4()

	item1 := Item{
		ID:   iID1,
		Name: "item1",
	}
	item2 := Item{
		ID:   iID2,
		Name: "item2",
	}
	sale := Sales{
		UserID: sID,
	}

	Gorm.Table("items").Create(&item1)
	Gorm.Table("items").Create(&item2)
	Gorm.Create(&sale)
	itemSale1 := ItemSale{SaleID: sale.ID, ItemID: iID1, Count: 8, Price: 12.20}
	itemSale2 := ItemSale{SaleID: sale.ID, ItemID: iID2, Count: 12, Price: 5.20}
	Gorm.Create(&itemSale1)
	Gorm.Create(&itemSale2)

	items, err := s.GetItemsInSaleBySaleID(sale.ID)
	assert.Nil(t, err)
	assert.True(t, len(items) == 2 && items[1].Name == "item2")

	Gorm.Table("items").Delete(&item1)
	Gorm.Table("items").Delete(&item2)
	Gorm.Delete(&sale)
	Gorm.Delete(&itemSale1)
	Gorm.Delete(&itemSale2)
}

func TestStore_DeleteItemsInSale(t *testing.T) {
	s := NewStore(Gorm, RedisClient)

	sID, _ := uuid.NewV4()
	iID1, _ := uuid.NewV4()
	iID2, _ := uuid.NewV4()

	item1 := Item{
		ID:   iID1,
		Name: "item1",
	}
	item2 := Item{
		ID:   iID2,
		Name: "item2",
	}
	sale := Sales{
		UserID: sID,
	}

	Gorm.Table("items").Create(&item1)
	Gorm.Table("items").Create(&item2)
	Gorm.Create(&sale)
	itemSale1 := ItemSale{SaleID: sale.ID, ItemID: iID1, Count: 8, Price: 12.20}
	itemSale2 := ItemSale{SaleID: sale.ID, ItemID: iID2, Count: 12, Price: 5.20}
	Gorm.Create(&itemSale1)
	Gorm.Create(&itemSale2)

	err := s.DeleteItemsInSale(Gorm, sale.ID)
	assert.Nil(t, err)

	Gorm.Table("items").Delete(&item1)
	Gorm.Table("items").Delete(&item2)
	Gorm.Delete(&sale)
}

func TestStore_AddItemToSale(t *testing.T) {
	s := NewStore(Gorm, RedisClient)

	sID, _ := uuid.NewV4()
	iID1, _ := uuid.NewV4()

	item := models.ItemSale{ID: strfmt.UUID(iID1.String()), Name: "item1", Count: 3, Price: 12.20}

	err := s.AddItemToSale(Gorm, sID, &item)
	assert.Nil(t, err)

	itemSale := ItemSale{SaleID: sID, ItemID: iID1}
	Gorm.First(&itemSale)

	assert.True(t, itemSale.Price == 12.20)

	Gorm.Delete(&itemSale)
}
