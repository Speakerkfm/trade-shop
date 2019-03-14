package store

import (
	"fmt"
	"github.com/go-redis/cache"
	"time"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

const cacheLiveTime = 10 * time.Minute
const inventoryRedisPrefix = "rate_inventory_"

type Inventory struct {
	UserID uuid.UUID `gorm:"primary_key"`
	ItemID uuid.UUID `gorm:"primary_key"`
	Name   string    `gorm:"-"`
	Count  int64
}

func (st *Store) GetInventoryByUserId(userID uuid.UUID) []*Inventory {
	var Items []*Inventory

	cacheKey := fmt.Sprint(inventoryRedisPrefix, userID)
	if err := st.codec.Get(cacheKey, Items); err == nil {
		return Items
	}

	st.gorm.Raw(`
		select i.user_id, i.item_id, its.name, i.count
		from inventory i join items its 
		on i.item_id = its.id
		where i.user_id = ?`, userID).Scan(&Items)

	if err := st.codec.Set(&cache.Item{
		Key:        cacheKey,
		Object:     Items,
		Expiration: cacheLiveTime,
	}); err != nil {
		panic(err)
	}

	return Items
}

func (st *Store) ClearInventoryCache(userID uuid.UUID) {
	cacheKey := fmt.Sprint(inventoryRedisPrefix, userID)
	st.codec.Redis.Del(cacheKey)
}

func (st *Store) AddItemToUser(db *gorm.DB, userID uuid.UUID, item *ItemSale) error {
	inventory := Inventory{
		UserID: userID,
		ItemID: item.ItemID,
	}

	if err := db.First(&inventory).Error; notFound(err) {
		inventory.Count = item.Count

		return db.Create(&inventory).Error
	}

	inventory.Count += item.Count

	return db.Save(&inventory).Error
}

func (st *Store) RemoveItemFromUser(db *gorm.DB, userID uuid.UUID, itemID uuid.UUID, count int64) error {
	inventory := Inventory{UserID: userID, ItemID: itemID}

	if err := db.First(&inventory).Error; err != nil {
		return fmt.Errorf("not enough items")
	}

	if inventory.Count < count {
		return fmt.Errorf("not enough items")
	}

	if inventory.Count == count {
		return db.Delete(&inventory).Error
	}

	if inventory.Count > count {
		inventory.Count -= count
		return db.Save(&inventory).Error
	}

	return nil
}
