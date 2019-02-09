package store

type ItemSale struct {
	SaleID int64  `gorm:"column:sale_id"`
	ItemID int64  `gorm:"column:item_id"`
	Name   string `gorm:"-"`
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
