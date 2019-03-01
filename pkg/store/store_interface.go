// Code generated by ifacemaker. DO NOT EDIT.

package store

import (
	"trade-shop/pkg/models"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

type StoreInterface interface {
	GetInventoryByUserId(userID uuid.UUID) []*Inventory
	AddItemToUser(db *gorm.DB, userID uuid.UUID, item ItemSale) error
	RemoveItemFromUser(db *gorm.DB, userID uuid.UUID, itemID uuid.UUID, count int64) error
	GetItemByID(itemID uuid.UUID) (*Item, bool)
	GetSaleItemList(userID uuid.UUID) ([]ItemSale, error)
	GetItemsInSaleBySaleID(saleID uuid.UUID) ([]ItemSale, error)
	DeleteItemsInSale(db *gorm.DB, saleID uuid.UUID) error
	AddItemToSale(db *gorm.DB, saleID uuid.UUID, item *models.ItemSale) error
	CreateNewSale(db *gorm.DB, userID uuid.UUID) uuid.UUID
	DeleteSaleBySaleID(db *gorm.DB, saleID uuid.UUID) error
	GetSellerBySaleID(saleID uuid.UUID) (uuid.UUID, error)
	CreateTransaction() *gorm.DB
	UserByEmail(email string) (*User, bool)
	UserByUserID(userID uuid.UUID) (*User, bool)
	GetUserBill(userID uuid.UUID) float64
	AddMoneyToUser(db *gorm.DB, userID uuid.UUID, money float64) error
	RemoveMoneyFromUser(db *gorm.DB, userID uuid.UUID, money float64) error
}
