package amqp

import (
	"encoding/json"
	amqp "github.com/rabbitmq/amqp091-go"
)

func PublishWithDelay[Obj any](exchangeName string, object *Obj, delay int) error {
	header := amqp.Table{
		"x-delay": delay * 1000 * 60,
	}
	body, err := json.Marshal(object)
	if err != nil {
		return err
	}

	ch := GetChannel()
	errPub := ch.Publish(
		exchangeName,
		"",
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Headers:     header,
			Body:        body,
		},
	)

	return errPub
}
