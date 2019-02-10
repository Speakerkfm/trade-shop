package store

import (
	"github.com/satori/go.uuid"
)

type Item struct {
	ID   uuid.UUID `gorm:"primary_key"`
	Name *string
}

func (st *Store) GetItemByID(itemID uuid.UUID) (*Item, bool) {
	var item Item

	err := st.gorm.First(&item, "id = ?", itemID).Error

	return &item, found(err)
}
