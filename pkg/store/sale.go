package store

import (
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

type Sales struct {
	ID     uuid.UUID `gorm:"primary_key"`
	UserID uuid.UUID `gorm:"column:user_id"`
}

func (s *Sales) BeforeCreate(scope *gorm.Scope) error {
	id := uuid.NewV4()
	return scope.SetColumn("id", id)
}

func (st *Store) CreateNewSale(db *gorm.DB, userID uuid.UUID) uuid.UUID {
	sale := Sales{
		UserID: uuid.FromStringOrNil(userID.String()),
	}

	if err := db.Create(&sale).Error; err != nil {
		panic(err)
	}

	return sale.ID
}

func (st *Store) DeleteSaleBySaleID(db *gorm.DB, saleID uuid.UUID) error {
	return db.Delete(&Sales{ID: saleID}).Error
}

func (st *Store) GetSellerBySaleID(saleID uuid.UUID) (uuid.UUID, error) {
	var sale Sales

	err := st.gorm.First(&sale, "ID = ?", saleID).Error

	return sale.UserID, err
}
