// Code generated by ifacemaker. DO NOT EDIT.

package serviceiface

import (
	"trade-shop/pkg/models"

	uuid "github.com/satori/go.uuid"
)

type Inventory interface {
	MakeInventory(userID uuid.UUID) []*models.Item
}
