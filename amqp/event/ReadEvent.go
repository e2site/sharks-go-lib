package event

import (
	"github.com/rabbitmq/amqp091-go"
)

const EventHeaderName = "event"

type eventName struct {
	EventType string
}

func GetEvent(data *amqp091.Delivery) string {
	evntName, ok := data.Headers[EventHeaderName].(string)
	if ok {
		return evntName
	}
	return ""
}
