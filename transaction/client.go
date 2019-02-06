package transaction

import (
	"encoding/json"
	"github.com/streadway/amqp"
)

type ClientConfig struct {
	Address string
}
type TransactionClient struct {
	Connection *amqp.Connection
	Channel    *amqp.Channel
}

func (c *TransactionClient) Connect(config ClientConfig) error {
	conn, err := amqp.Dial(config.Address)
	if err != nil {
		return err
	}
	c.Connection = conn

	ch, err := conn.Channel()
	if err != nil {
		return err
	}
	c.Channel = ch
	return nil
}

func (c *TransactionClient) NewTransaction(transaction *TransactionPayload) error {
	body, err := json.Marshal(transaction)
	if err != nil {
		return err
	}
	err = c.Channel.Publish(
		"transaction", // exchange
		"",            // routing key
		false,         // mandatory
		false,         // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		})
	return err
}
