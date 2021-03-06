// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import gorm "github.com/jinzhu/gorm"
import mock "github.com/stretchr/testify/mock"
import models "trade-shop/pkg/models"
import store "trade-shop/pkg/store"
import uuid "github.com/satori/go.uuid"

// StoreInterface is an autogenerated mock type for the StoreInterface type
type StoreInterface struct {
	mock.Mock
}

// AddItemToSale provides a mock function with given fields: db, saleID, item
func (_m *StoreInterface) AddItemToSale(db *gorm.DB, saleID uuid.UUID, item *models.ItemSale) error {
	ret := _m.Called(db, saleID, item)

	var r0 error
	if rf, ok := ret.Get(0).(func(*gorm.DB, uuid.UUID, *models.ItemSale) error); ok {
		r0 = rf(db, saleID, item)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// AddItemToUser provides a mock function with given fields: db, userID, item
func (_m *StoreInterface) AddItemToUser(db *gorm.DB, userID uuid.UUID, item *store.ItemSale) error {
	ret := _m.Called(db, userID, item)

	var r0 error
	if rf, ok := ret.Get(0).(func(*gorm.DB, uuid.UUID, *store.ItemSale) error); ok {
		r0 = rf(db, userID, item)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// AddMoneyToUser provides a mock function with given fields: db, userID, money
func (_m *StoreInterface) AddMoneyToUser(db *gorm.DB, userID uuid.UUID, money float64) error {
	ret := _m.Called(db, userID, money)

	var r0 error
	if rf, ok := ret.Get(0).(func(*gorm.DB, uuid.UUID, float64) error); ok {
		r0 = rf(db, userID, money)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ClearInventoryCache provides a mock function with given fields: userID
func (_m *StoreInterface) ClearInventoryCache(userID uuid.UUID) {
	_m.Called(userID)
}

// CommitTransaction provides a mock function with given fields: tx
func (_m *StoreInterface) CommitTransaction(tx *gorm.DB) {
	_m.Called(tx)
}

// CreateNewSale provides a mock function with given fields: db, userID
func (_m *StoreInterface) CreateNewSale(db *gorm.DB, userID uuid.UUID) uuid.UUID {
	ret := _m.Called(db, userID)

	var r0 uuid.UUID
	if rf, ok := ret.Get(0).(func(*gorm.DB, uuid.UUID) uuid.UUID); ok {
		r0 = rf(db, userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(uuid.UUID)
		}
	}

	return r0
}

// CreateNewUser provides a mock function with given fields: email, password
func (_m *StoreInterface) CreateNewUser(email string, password string) (*store.User, error) {
	ret := _m.Called(email, password)

	var r0 *store.User
	if rf, ok := ret.Get(0).(func(string, string) *store.User); ok {
		r0 = rf(email, password)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*store.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(email, password)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateTransaction provides a mock function with given fields:
func (_m *StoreInterface) CreateTransaction() *gorm.DB {
	ret := _m.Called()

	var r0 *gorm.DB
	if rf, ok := ret.Get(0).(func() *gorm.DB); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gorm.DB)
		}
	}

	return r0
}

// DeleteItemsInSale provides a mock function with given fields: db, saleID
func (_m *StoreInterface) DeleteItemsInSale(db *gorm.DB, saleID uuid.UUID) error {
	ret := _m.Called(db, saleID)

	var r0 error
	if rf, ok := ret.Get(0).(func(*gorm.DB, uuid.UUID) error); ok {
		r0 = rf(db, saleID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteSaleBySaleID provides a mock function with given fields: db, saleID
func (_m *StoreInterface) DeleteSaleBySaleID(db *gorm.DB, saleID uuid.UUID) error {
	ret := _m.Called(db, saleID)

	var r0 error
	if rf, ok := ret.Get(0).(func(*gorm.DB, uuid.UUID) error); ok {
		r0 = rf(db, saleID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetInventoryByUserId provides a mock function with given fields: userID
func (_m *StoreInterface) GetInventoryByUserId(userID uuid.UUID) []*store.Inventory {
	ret := _m.Called(userID)

	var r0 []*store.Inventory
	if rf, ok := ret.Get(0).(func(uuid.UUID) []*store.Inventory); ok {
		r0 = rf(userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*store.Inventory)
		}
	}

	return r0
}

// GetItemsInSaleBySaleID provides a mock function with given fields: saleID
func (_m *StoreInterface) GetItemsInSaleBySaleID(saleID uuid.UUID) ([]*store.ItemSale, error) {
	ret := _m.Called(saleID)

	var r0 []*store.ItemSale
	if rf, ok := ret.Get(0).(func(uuid.UUID) []*store.ItemSale); ok {
		r0 = rf(saleID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*store.ItemSale)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uuid.UUID) error); ok {
		r1 = rf(saleID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetSaleItemList provides a mock function with given fields: userID
func (_m *StoreInterface) GetSaleItemList(userID uuid.UUID) ([]*store.ItemSale, error) {
	ret := _m.Called(userID)

	var r0 []*store.ItemSale
	if rf, ok := ret.Get(0).(func(uuid.UUID) []*store.ItemSale); ok {
		r0 = rf(userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*store.ItemSale)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uuid.UUID) error); ok {
		r1 = rf(userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetSellerBySaleID provides a mock function with given fields: saleID
func (_m *StoreInterface) GetSellerBySaleID(saleID uuid.UUID) (uuid.UUID, error) {
	ret := _m.Called(saleID)

	var r0 uuid.UUID
	if rf, ok := ret.Get(0).(func(uuid.UUID) uuid.UUID); ok {
		r0 = rf(saleID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(uuid.UUID)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uuid.UUID) error); ok {
		r1 = rf(saleID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUserBill provides a mock function with given fields: userID
func (_m *StoreInterface) GetUserBill(userID uuid.UUID) float64 {
	ret := _m.Called(userID)

	var r0 float64
	if rf, ok := ret.Get(0).(func(uuid.UUID) float64); ok {
		r0 = rf(userID)
	} else {
		r0 = ret.Get(0).(float64)
	}

	return r0
}

// GetUserSaleItemList provides a mock function with given fields: userID
func (_m *StoreInterface) GetUserSaleItemList(userID uuid.UUID) ([]*store.ItemSale, error) {
	ret := _m.Called(userID)

	var r0 []*store.ItemSale
	if rf, ok := ret.Get(0).(func(uuid.UUID) []*store.ItemSale); ok {
		r0 = rf(userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*store.ItemSale)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uuid.UUID) error); ok {
		r1 = rf(userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IsEmailTaken provides a mock function with given fields: email
func (_m *StoreInterface) IsEmailTaken(email string) bool {
	ret := _m.Called(email)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(email)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// RemoveItemFromUser provides a mock function with given fields: db, userID, itemID, count
func (_m *StoreInterface) RemoveItemFromUser(db *gorm.DB, userID uuid.UUID, itemID uuid.UUID, count int64) error {
	ret := _m.Called(db, userID, itemID, count)

	var r0 error
	if rf, ok := ret.Get(0).(func(*gorm.DB, uuid.UUID, uuid.UUID, int64) error); ok {
		r0 = rf(db, userID, itemID, count)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RemoveMoneyFromUser provides a mock function with given fields: db, userID, money
func (_m *StoreInterface) RemoveMoneyFromUser(db *gorm.DB, userID uuid.UUID, money float64) error {
	ret := _m.Called(db, userID, money)

	var r0 error
	if rf, ok := ret.Get(0).(func(*gorm.DB, uuid.UUID, float64) error); ok {
		r0 = rf(db, userID, money)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RollbackTransaction provides a mock function with given fields: tx
func (_m *StoreInterface) RollbackTransaction(tx *gorm.DB) {
	_m.Called(tx)
}

// UserByEmail provides a mock function with given fields: email
func (_m *StoreInterface) UserByEmail(email string) (*store.User, bool) {
	ret := _m.Called(email)

	var r0 *store.User
	if rf, ok := ret.Get(0).(func(string) *store.User); ok {
		r0 = rf(email)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*store.User)
		}
	}

	var r1 bool
	if rf, ok := ret.Get(1).(func(string) bool); ok {
		r1 = rf(email)
	} else {
		r1 = ret.Get(1).(bool)
	}

	return r0, r1
}

// UserByUserID provides a mock function with given fields: userID
func (_m *StoreInterface) UserByUserID(userID uuid.UUID) (*store.User, bool) {
	ret := _m.Called(userID)

	var r0 *store.User
	if rf, ok := ret.Get(0).(func(uuid.UUID) *store.User); ok {
		r0 = rf(userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*store.User)
		}
	}

	var r1 bool
	if rf, ok := ret.Get(1).(func(uuid.UUID) bool); ok {
		r1 = rf(userID)
	} else {
		r1 = ret.Get(1).(bool)
	}

	return r0, r1
}
