package store

import (
	"github.com/jinzhu/gorm"
	"github.com/satori/go.uuid"
	"trade-shop/pkg/models"
)

type ItemSale struct {
	SellerID uuid.UUID `gorm:"-"`
	SaleID   uuid.UUID `gorm:"column:sale_id"`
	ItemID   uuid.UUID `gorm:"column:item_id"`
	Name     string    `gorm:"-"`
	Count    int64
	Price    float64
}

func (st *Store) GetSaleItemList(userID uuid.UUID) ([]*ItemSale, error) {
	var ItemSales []*ItemSale
	err := st.gorm.Raw(`
		select s.user_id, i.sale_id, i.item_id, its.name, i.count, i.price
		from item_sale i join items its 
		on i.item_id = its.id join sales s on i.sale_id = s.id where s.user_id <> ?`, userID).Scan(&ItemSales).Error

	return ItemSales, err
}

func (st *Store) GetItemsInSaleBySaleID(saleID uuid.UUID) ([]*ItemSale, error) {
	var ItemSales []*ItemSale
	err := st.gorm.Raw(`
		select *
		from item_sale i join items its 
		on i.item_id = its.id
		where i.sale_id = ?`, saleID).Scan(&ItemSales).Error

	return ItemSales, err
}

func (st *Store) DeleteItemsInSale(db *gorm.DB, saleID uuid.UUID) error {
	return db.Where("sale_id = ?", saleID).Delete(&ItemSale{}).Error
}

func (st *Store) AddItemToSale(db *gorm.DB, saleID uuid.UUID, item *models.ItemSale) error {
	itemSale := ItemSale{
		SaleID: saleID,
		ItemID: uuid.FromStringOrNil(item.ID.String()),
		Name:   item.Name,
		Count:  item.Count,
		Price:  item.Price,
	}

	err := db.Create(&itemSale).Error

	return err
}
