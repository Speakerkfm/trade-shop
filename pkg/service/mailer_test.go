package service

import (
	"encoding/json"
	"testing"
	"trade-shop/pkg/mocks"
	"trade-shop/pkg/store"

	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
)

func TestMailer_SendNotificationEmail(t *testing.T) {
	amqp := mocks.IAmqpService{}
	mailer := &Mailer{amqpService: &amqp}
	sID, _ := uuid.NewV4()
	saleID, _ := uuid.NewV4()
	iID, _ := uuid.NewV4()

	cases := map[string]struct {
		to       string
		itemList []*store.ItemSale
		ok       bool
		err      error
	}{
		"ok": {
			to:       "asdf@mail.com",
			itemList: []*store.ItemSale{{SellerID: sID, SaleID: saleID, ItemID: iID, Name: "item1", Count: 3, Price: 12.20}},
			ok:       true,
			err:      nil,
		},
	}

	for _, test := range cases {
		msg := Message{
			EmailTo:   test.to,
			EmailType: "email_notification",
			Data:      map[string]interface{}{"items": test.itemList},
		}

		data, _ := json.Marshal(&msg)

		amqp.Mock.On("Publish", data).Return(test.err)
		err := mailer.SendNotificationEmail(test.to, test.itemList)
		assert.True(t, test.ok == (err == nil))
	}
}
