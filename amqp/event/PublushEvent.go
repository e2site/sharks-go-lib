package event

import (
	"github.com/e2site/sharks-go-lib/amqp"
	"github.com/e2site/sharks-go-lib/log"
)

const EventBusExchanger = "event-bus-exchanger"

var statusInit = false

type EventMessage struct {
	EventType string
	Data      any
}

func DeclareEventBus() {
	if !statusInit {
		amqp.DeclareFanout(EventBusExchanger)
		statusInit = true
	}
}

func CreateEvent[Obj any](event string, data *Obj) {
	DeclareEventBus()

	var newEvent EventMessage
	newEvent.EventType = event
	newEvent.Data = data

	err := amqp.PublishMessageWithoutTracer(EventBusExchanger, &newEvent)
	if err != nil {
		log.Log(err)
	}
}
