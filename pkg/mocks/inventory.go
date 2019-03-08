// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"
import models "trade-shop/pkg/models"

import uuid "github.com/satori/go.uuid"

// Inventory is an autogenerated mock type for the Inventory type
type Inventory struct {
	mock.Mock
}

// MakeInventory provides a mock function with given fields: userID
func (_m *Inventory) MakeInventory(userID uuid.UUID) []*models.Item {
	ret := _m.Called(userID)

	var r0 []*models.Item
	if rf, ok := ret.Get(0).(func(uuid.UUID) []*models.Item); ok {
		r0 = rf(userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.Item)
		}
	}

	return r0
}