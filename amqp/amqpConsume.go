package amqp

import (
	"github.com/e2site/sharks-go-lib/log"
	"github.com/rabbitmq/amqp091-go"
)

func AmqpConsume(queue string) <-chan amqp091.Delivery {
	ch := GetChannel()

	err := ch.Qos(
		1,     // prefetch count
		0,     // prefetch size
		false, // global
	)
	log.CheckEndLogFatal(err)

	msgs, errConsume := ch.Consume(
		queue, // queue
		"",    // consumer
		false, // auto-ack
		false, // exclusive
		false, // no-local
		false, // no-wait
		nil,   // args
	)
	log.CheckEndLogFatal(errConsume)
	return msgs
}
