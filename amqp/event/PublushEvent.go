package event

import (
	"github.com/e2site/sharks-go-lib/amqp"
	"github.com/e2site/sharks-go-lib/log"
)

const EventBusExchanger = "event-bus-exchanger"

var statusInit = false

func DeclareEventBus() {
	if !statusInit {
		amqp.DeclareFanout(EventBusExchanger)
		statusInit = true
	}
}

func CreateEvent[Obj any](event string, data *Obj) {
	DeclareEventBus()

	err := amqp.PublishEventWithoutTracer(EventBusExchanger, event, data)
	if err != nil {
		log.Log(err)
	}
}
