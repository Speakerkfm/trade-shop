package store

import (
	"errors"
	"github.com/jinzhu/gorm"
	"github.com/satori/go.uuid"
)

type Inventory struct {
	UserID uuid.UUID `gorm:"primary_key"`
	ItemID uuid.UUID `gorm:"primary_key"`
	Name   string    `gorm:"-"`
	Count  int64
}

func (st *Store) GetInventoryByUserId(userID uuid.UUID) []*Inventory {
	var Items []*Inventory
	st.gorm.Raw(`
		select i.user_id, i.item_id, its.name, i.count
		from inventory i join items its 
		on i.item_id = its.id
		where i.user_id = ?`, userID).Scan(&Items)

	return Items
}

func (st *Store) AddItemToUser(db *gorm.DB, userID uuid.UUID, item ItemSale) error {
	inventory := Inventory{
		UserID: userID,
		ItemID: item.ItemID,
		Name:   item.Name,
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
		return errors.New("NotEnoughItems")
	}

	if inventory.Count < count {
		return errors.New("NotEnoughItems")
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
