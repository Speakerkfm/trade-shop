package service

import (
	"fmt"
	"trade-shop/pkg/flags"

	"github.com/streadway/amqp"
)

type Queue struct {
	channel    *amqp.Channel
	Connection *amqp.Connection
}

type IAmqpService interface {
	Publish(data []byte) error
}

func NewQueue(conf *flags.Config) (*Queue, error) {
	var err error
	var connection *amqp.Connection
	var channel *amqp.Channel

	URI := fmt.Sprintf("amqp://%s:%s@%s:%d/", conf.AmqpUser, conf.AmqpPassword, conf.AmqpHost, conf.AmqpPort)
	if connection, err = amqp.Dial(URI); err != nil {
		return nil, fmt.Errorf("amqp connection failed err: %s", err)
	}

	if channel, err = connection.Channel(); err != nil {
		return nil, fmt.Errorf("amqp connection failed err: %s", err)
	}

	if _, err = channel.QueueDeclare(
		"mailer",
		true,
		false,
		false,
		false,
		amqp.Table{"x-dead-letter-exchange": "mailer_fail"},
	); err != nil {
		return nil, fmt.Errorf("amqp connection failed err: %s", err)
	}

	return &Queue{channel: channel, Connection: connection}, nil
}

func (q *Queue) Publish(data []byte) error {
	err := q.channel.Publish(
		"",
		"mailer",
		false,
		false,
		amqp.Publishing{
			Headers:      amqp.Table{},
			DeliveryMode: amqp.Persistent,
			ContentType:  "text/plain",
			Body:         data,
		})

	return err
}
