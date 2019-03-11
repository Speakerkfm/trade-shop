package store

import (
	"fmt"
	"testing"
	"time"

	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
)

const iName1 = "item1"
const iName2 = "item2"

func TestStore_GetInventoryByUserId(t *testing.T) {
	s := NewStore(Gorm, RedisClient)

	iID1 := uuid.NewV4()
	iID2 := uuid.NewV4()
	uID := uuid.NewV4()
	u := User{
		ID: uID,
	}
	item1 := Item{
		ID:   iID1,
		Name: iName1,
	}
	item2 := Item{
		ID:   iID2,
		Name: iName2,
	}
	inv1 := Inventory{
		UserID: uID,
		ItemID: iID1,
		Count:  5,
	}
	inv2 := Inventory{
		UserID: uID,
		ItemID: iID2,
		Count:  7,
	}

	Gorm.Create(&u)
	Gorm.Table("items").Create(&item1)
	Gorm.Table("items").Create(&item2)
	Gorm.Create(&inv1)
	Gorm.Create(&inv2)

	res1 := s.GetInventoryByUserId(uID)
	cacheKey := fmt.Sprint("rate_inventory_", uID)
	rs := s.codec.Redis.Get(cacheKey)
	assert.True(t, res1[1].Name == iName2)
	assert.True(t, rs.Val() != "")

	Gorm.Delete(&u)
	Gorm.Table("items").Delete(&item1)
	Gorm.Table("items").Delete(&item2)
	Gorm.Delete(&inv1)
	Gorm.Delete(&inv2)
}

func TestStore_ClearInventoryCache(t *testing.T) {
	s := NewStore(Gorm, RedisClient)

	uID := uuid.NewV4()
	cacheKey := fmt.Sprint("rate_inventory_", uID)

	s.codec.Redis.Set(cacheKey, "12345", 10*time.Minute)

	s.ClearInventoryCache(uID)

	res := s.codec.Redis.Get(cacheKey)
	assert.True(t, res.Val() == "")
}

func TestStore_AddItemToUser(t *testing.T) {
	s := NewStore(Gorm, RedisClient)

	iID1 := uuid.NewV4()
	iID2 := uuid.NewV4()
	uID := uuid.NewV4()
	u := User{
		ID: uID,
	}
	item1 := Item{
		ID:   iID1,
		Name: iName1,
	}
	item2 := Item{
		ID:   iID2,
		Name: iName2,
	}
	inv1 := Inventory{
		UserID: uID,
		ItemID: iID1,
		Count:  5,
	}
	inv2 := Inventory{
		UserID: uID,
		ItemID: iID2,
	}

	Gorm.Create(&u)
	Gorm.Table("items").Create(&item1)
	Gorm.Table("items").Create(&item2)
	Gorm.Create(&inv1)

	err := s.AddItemToUser(Gorm, uID, &ItemSale{ItemID: iID1, Name: iName1, Count: 5})
	s.gorm.First(&inv1)
	assert.Nil(t, err)
	assert.True(t, inv1.Count == 10)

	err = s.AddItemToUser(Gorm, uID, &ItemSale{ItemID: iID2, Name: iName2, Count: 3})
	s.gorm.First(&inv2)
	assert.Nil(t, err)
	assert.True(t, inv2.Count == 3)

	Gorm.Delete(&u)
	Gorm.Table("items").Delete(&item1)
	Gorm.Table("items").Delete(&item2)
	Gorm.Delete(&inv1)
	Gorm.Delete(&inv2)
}

func TestStore_RemoveItemFromUser(t *testing.T) {
	s := NewStore(Gorm, RedisClient)

	iID1 := uuid.NewV4()
	iID2 := uuid.NewV4()
	uID := uuid.NewV4()
	u := User{
		ID: uID,
	}
	item1 := Item{
		ID:   iID1,
		Name: iName1,
	}
	item2 := Item{
		ID:   iID2,
		Name: iName2,
	}
	inv1 := Inventory{
		UserID: uID,
		ItemID: iID1,
		Count:  5,
	}
	inv2 := Inventory{
		UserID: uID,
		ItemID: iID2,
		Count:  1,
	}

	Gorm.Create(&u)
	Gorm.Table("items").Create(&item1)
	Gorm.Table("items").Create(&item2)
	Gorm.Create(&inv1)

	err := s.RemoveItemFromUser(Gorm, uID, iID1, 4)
	s.gorm.First(&inv1)
	assert.True(t, inv1.Count == 1)
	assert.Nil(t, err)

	err = s.RemoveItemFromUser(Gorm, uID, iID2, 4)
	s.gorm.First(&inv2)
	assert.True(t, inv2.Count == 1)
	assert.True(t, err.Error() == "not enough items")

	Gorm.Delete(&u)
	Gorm.Table("items").Delete(&item1)
	Gorm.Table("items").Delete(&item2)
	Gorm.Delete(&inv1)
	Gorm.Delete(&inv2)
}
