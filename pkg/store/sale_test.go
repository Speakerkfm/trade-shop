package store

import (
	"github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStore_CreateNewSale(t *testing.T) {
	s := NewStore(Gorm, RedisClient)

	uID, _ := uuid.NewV4()

	sID := s.CreateNewSale(Gorm, uID)
	sale := Sales{ID: sID}

	Gorm.First(&sale)
	assert.True(t, sale.UserID == uID)

	Gorm.Delete(&sale)
}

func TestStore_DeleteSaleBySaleID(t *testing.T) {
	s := NewStore(Gorm, RedisClient)

	uID, _ := uuid.NewV4()

	sID := s.CreateNewSale(Gorm, uID)

	err := s.DeleteSaleBySaleID(Gorm, sID)
	assert.Nil(t, err)
}

func TestStore_GetSellerBySaleID(t *testing.T) {
	s := NewStore(Gorm, RedisClient)

	uID, _ := uuid.NewV4()

	sID := s.CreateNewSale(Gorm, uID)

	sellerID, err := s.GetSellerBySaleID(sID)
	assert.Nil(t, err)
	assert.True(t, sellerID == uID)

	Gorm.Delete(Sales{ID: sID})
}
