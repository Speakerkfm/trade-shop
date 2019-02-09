package store

import "github.com/satori/go.uuid"

type Sale struct {
	ID          uuid.UUID `gorm:"primary_key"`
	UserID      int32     `gorm:"column:user_id"`
	ItemsInSale *Item     `gorm:"-"`
}
