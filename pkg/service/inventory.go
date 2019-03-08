//ifacemaker -f inventory.go -s Inventory -i Inventory -p serviceiface -o serviceiface/inventory.go
package service

import (
	"trade-shop/pkg/models"
	"trade-shop/pkg/store"

	"github.com/go-openapi/strfmt"
	uuid "github.com/satori/go.uuid"
)

type Inventory struct {
	st store.StoreInterface
}

func NewInventory(st store.StoreInterface) *Inventory {
	return &Inventory{st: st}
}

func (i *Inventory) MakeInventory(userID uuid.UUID) []*models.Item {
	inventory := i.st.GetInventoryByUserId(userID)

	var items []*models.Item

	for _, val := range inventory {
		items = append(items, &models.Item{ID: strfmt.UUID(val.ItemID.String()), Name: val.Name, Count: val.Count})
	}

	return items
}
