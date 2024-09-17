package amqp

import (
	"github.com/e2site/sharks-go-lib/log"
	amqp "github.com/rabbitmq/amqp091-go"
)

func QueueBindWithKey(exchangeName string, queue *amqp.Queue, key string) {
	ch := GetChannel()
	err := ch.QueueBind(
		queue.Name,
		key,
		exchangeName,
		false,
		nil,
	)
	log.CheckEndLogFatal(err)
}
