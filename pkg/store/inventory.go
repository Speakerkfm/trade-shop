package store

import (
	"github.com/satori/go.uuid"
	"trade-shop/pkg/models"
)

type Inventory struct {
	UserID uuid.UUID `gorm:"column:user_id"`
	ItemID uuid.UUID `gorm:"column:item_id"`
	Name   string    `gorm:"-"`
	Count  int64
}

func (st *Store) GetInventoryByUserId(userID uuid.UUID) ([]*models.Item, bool) {
	var Items []*models.Item
	err := st.gorm.Raw(`
		select its.ID, its.name, i.count
		from inventory i join items its 
		on i.item_id = its.id`).Scan(&Items).Error

	return Items, found(err)
}

func (st *Store) CreateItemToUser(userID uuid.UUID, item *models.Item) bool {
	inventory := Inventory{
		UserID: userID,
		ItemID: uuid.FromStringOrNil(item.ID.String()),
		Name:   item.Name,
		Count:  item.Count,
	}

	err := st.gorm.Create(&inventory).Error

	return found(err)
}

func (st *Store) UpdateItemToUser(userID uuid.UUID, itemID uuid.UUID, newCount int) bool {
	err := st.gorm.
		Table("inventory").
		Where("user_id = ? and item_id = ?", userID.String(), itemID.String()).
		Updates(map[string]interface{}{"count": newCount}).Error

	return found(err)
}

func (st *Store) DeleteItemFromUser(userID uuid.UUID, itemID uuid.UUID) bool {
	err := st.gorm.
		Table("inventory").
		Where("user_id = ? and item_id = ?", userID, itemID).
		Delete(&Inventory{}).Error

	return found(err)
}
