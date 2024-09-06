package amqp

import (
	"github.com/e2site/sharks-go-lib/log"
	amqp "github.com/rabbitmq/amqp091-go"
)

func QueueBind(exchangeName string, queue *amqp.Queue) {
	ch := GetChannel()
	err := ch.QueueBind(
		queue.Name,
		"",
		exchangeName,
		false,
		nil,
	)
	log.CheckEndLogFatal(err)
}
