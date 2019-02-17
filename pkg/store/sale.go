package store

import (
	"github.com/satori/go.uuid"
)

type Sale struct {
	ID     uuid.UUID `gorm:"primary_key"`
	UserID uuid.UUID `gorm:"column:user_id"`
}

func (st *Store) CreateNewSale(userID uuid.UUID) (uuid.UUID, bool) {
	id, _ := uuid.NewV4()
	sale := Sale{
		ID:     id,
		UserID: uuid.FromStringOrNil(userID.String()),
	}

	err := st.gorm.Create(&sale).Error

	return id, found(err)
}

func (st *Store) DeleteSaleBySaleID(saleID uuid.UUID) bool {
	err := st.gorm.
		Table("sales").
		Where("sale_id = ?", saleID).
		Delete(&Inventory{}).Error

	return found(err)
}

/*
func (st *Store) GetSellerBySaleID(saleID uuid.UUID) (uuid.UUID, error){
	var sale Sale

	err := st.gorm.Table("sales").First(&sale, "ID = ?", saleID).Error

	return sale.UserID, err
}

func (st *Store) Purchase(userID uuid.UUID, saleID uuid.UUID) error {
	sellerID, err := st.GetSellerBySaleID(saleID)
	if err != nil {
		panic(err)
	}

	tx := st.gorm.Begin()

	items, ok := st.GetItemsInSaleBySaleID(saleID)

	if !ok {
		tx.Rollback()
		panic()
	}
}

*/
