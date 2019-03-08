// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"
import models "trade-shop/pkg/models"

import uuid "github.com/satori/go.uuid"

// Sale is an autogenerated mock type for the Sale type
type Sale struct {
	mock.Mock
}

// CreateLot provides a mock function with given fields: userID, itemList
func (_m *Sale) CreateLot(userID uuid.UUID, itemList []*models.ItemSale) error {
	ret := _m.Called(userID, itemList)

	var r0 error
	if rf, ok := ret.Get(0).(func(uuid.UUID, []*models.ItemSale) error); ok {
		r0 = rf(userID, itemList)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MakeSalesList provides a mock function with given fields: userID
func (_m *Sale) MakeSalesList(userID uuid.UUID) []*models.Sale {
	ret := _m.Called(userID)

	var r0 []*models.Sale
	if rf, ok := ret.Get(0).(func(uuid.UUID) []*models.Sale); ok {
		r0 = rf(userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.Sale)
		}
	}

	return r0
}

// Purchase provides a mock function with given fields: userID, sellerID, saleID
func (_m *Sale) Purchase(userID uuid.UUID, sellerID uuid.UUID, saleID uuid.UUID) error {
	ret := _m.Called(userID, sellerID, saleID)

	var r0 error
	if rf, ok := ret.Get(0).(func(uuid.UUID, uuid.UUID, uuid.UUID) error); ok {
		r0 = rf(userID, sellerID, saleID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}