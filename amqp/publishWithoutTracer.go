package amqp

import (
	"encoding/json"
	amqp "github.com/rabbitmq/amqp091-go"
)

func PublishMessageWithoutTracer[Obj any](exchangeName string, object *Obj) error {
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
			Body:        body,
		},
	)

	return errPub
}
