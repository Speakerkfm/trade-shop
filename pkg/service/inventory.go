//ifacemaker -f inventory.go -s Inventory -i Inventory -p serviceiface -o serviceiface/inventory.go
package service

import (
	"github.com/go-openapi/strfmt"
	"github.com/satori/go.uuid"
	"trade-shop/pkg/models"
	"trade-shop/pkg/store"
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
