package amqp

import (
	log2 "github.com/e2site/sharks-go-lib/log"
	amqp "github.com/rabbitmq/amqp091-go"
)

func QueueDeclare(queueName string) *amqp.Queue {
	ch := GetChannel()
	q, err := ch.QueueDeclare(
		queueName,
		true,
		false,
		false,
		false,
		nil,
	)
	log2.CheckEndLogFatal(err)
	return &q
}
