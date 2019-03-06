//ifacemaker -f mailer.go -s Mailer -i Mailer -p serviceiface -o serviceiface/mailer.go
package service

import (
	"encoding/json"
	"trade-shop/pkg/store"
)

const (
	typeEmailNotification = "email_notification"
)

type Message struct {
	EmailTo   string                 `json:"email"`
	EmailType string                 `json:"email_type"`
	Data      map[string]interface{} `json:"data"`
}

type Mailer struct {
	amqpService IAmqpService
}

func NewMailer(amqp *Queue) *Mailer {
	return &Mailer{amqpService: amqp}
}

func (mailer *Mailer) SendNotificationEmail(to string, itemList []*store.ItemSale) error {
	msg := Message{
		EmailTo:   to,
		EmailType: typeEmailNotification,
		Data:      map[string]interface{}{"items": itemList},
	}

	data, err := json.Marshal(&msg)
	if err != nil {
		return err
	}

	return mailer.amqpService.Publish(data)
}
