package store

import (
	"encoding/json"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/satori/go.uuid"
	"time"
)

const cacheLiveTime = 10 * time.Minute

type Inventory struct {
	UserID uuid.UUID `gorm:"primary_key"`
	ItemID uuid.UUID `gorm:"primary_key"`
	Name   string    `gorm:"-"`
	Count  int64
}

func (st *Store) GetInventoryByUserId(userID uuid.UUID) []*Inventory {
	var Items []*Inventory

	cacheKey := fmt.Sprint("rate_inventory_", userID)
	rate := st.codec.Redis.Get(cacheKey)

	if rate.Val() != "" {
		if bytes, err := rate.Bytes(); err == nil {
			if err := json.Unmarshal(bytes, &Items); err == nil {
				return Items
			}
		}
	}

	st.gorm.Raw(`
		select i.user_id, i.item_id, its.name, i.count
		from inventory i join items its 
		on i.item_id = its.id
		where i.user_id = ?`, userID).Scan(&Items)

	if value, err := json.Marshal(Items); err == nil {
		if err := st.codec.Redis.Set(cacheKey, value, cacheLiveTime).Err(); err != nil {
			panic(err)
		}
	}

	return Items
}

func (st *Store) ClearInventoryCache(userID uuid.UUID) {
	cacheKey := fmt.Sprint("rate_inventory_", userID)
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

	inventory.Count = inventory.Count + item.Count

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
		inventory.Count = inventory.Count - count
		return db.Save(&inventory).Error
	}

	return nil
}
