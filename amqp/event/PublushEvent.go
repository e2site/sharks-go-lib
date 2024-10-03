package event

import (
	"github.com/e2site/sharks-go-lib/amqp"
	"github.com/e2site/sharks-go-lib/log"
)

const EventBusExchanger = "event-bus-exchanger"

var statusInit = false
var reconnectHandlerInit = false

func DeclareEventBus() {
	if !statusInit {
		amqp.DeclareFanout(EventBusExchanger)
		statusInit = true
	}
	if !reconnectHandlerInit {
		reconnectHandlerInit = true
		amqp.SubscribeConnection(func(status bool) {
			if !status {
				statusInit = false
			}
		})
	}
}

func CreateEvent[Obj any](event string, data *Obj) {
	if !amqp.GetStatusConnected() {
		return
	}

	DeclareEventBus()

	err := amqp.PublishEventWithoutTracer(EventBusExchanger, event, data)
	if err != nil {
		log.Log(err)
	}
}
