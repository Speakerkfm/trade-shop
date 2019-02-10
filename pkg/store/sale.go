package store

import "github.com/satori/go.uuid"

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
