package status

import (
	"encoding/json"
	"github.com/e2site/sharks-go-lib/amqp"
	"github.com/e2site/sharks-go-lib/log"
	"github.com/google/uuid"
)

func ReadStatus(handler func(message *StatusMessage)) {
	id := uuid.New().String()

	var declareFunc = func() {
		queue := amqp.QueueDeclareWithTTL(id, 60000)
		amqp.QueueBind(StatusExchanger, queue)
		msg := amqp.AmqpConsume(queue.Name)
		go func() {
			for d := range msg {
				var status StatusMessage
				errJson := json.Unmarshal(d.Body, &status)
				if errJson != nil {
					log.Log(errJson)
				} else {
					handler(&status)
				}
				d.Ack(false)
			}
		}()
	}

	amqp.SubscribeConnection(func(status bool) {
		if status {
			declareFunc()
		}
	})

	declareFunc()

}
