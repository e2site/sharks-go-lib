package amqp

import (
	"encoding/json"
	amqp "github.com/rabbitmq/amqp091-go"
)

func PublishEventWithoutTracer[Obj any](exchangeName string, eventName string, object *Obj) error {
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
			Headers: amqp.Table{
				"event": eventName,
			},
		},
	)

	return errPub
}
