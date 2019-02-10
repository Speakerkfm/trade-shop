package store

import (
	"github.com/satori/go.uuid"
	"trade-shop/pkg/models"
)

type ItemSale struct {
	SaleID uuid.UUID `gorm:"column:sale_id"`
	ItemID uuid.UUID `gorm:"column:item_id"`
	Name   string    `gorm:"-"`
	Count  int64
	Price  float64
}

func (st *Store) GetSaleItemList() ([]ItemSale, bool) {
	var ItemSales []ItemSale
	err := st.gorm.Raw(`
		select *
		from item_sale i join items its 
		on i.item_id = its.id`).Scan(&ItemSales).Error

	return ItemSales, found(err)
}

func (st *Store) GetItemsInSaleBySaleID(saleID uuid.UUID) ([]ItemSale, bool) {
	var ItemSales []ItemSale
	err := st.gorm.Raw(`
		select *
		from item_sale i join items its 
		on i.item_id = its.id
		where i.sale_id = ?`, saleID).Scan(&ItemSales).Error

	return ItemSales, found(err)
}

func (st *Store) DeleteItemsInSaleBySaleID(saleID uuid.UUID) bool {
	err := st.gorm.
		Table("item_sale").
		Where("sale_id = ?", saleID).
		Delete(&ItemSale{}).Error

	return found(err)
}

func (st *Store) AddItemToSale(saleID uuid.UUID, item *models.SaleItemsItems0) bool {
	itemSale := ItemSale{
		SaleID: saleID,
		ItemID: uuid.FromStringOrNil(item.ID.String()),
		Name:   item.Name,
		Count:  item.Count,
		Price:  item.Price,
	}

	err := st.gorm.Create(&itemSale).Error

	return found(err)
}
