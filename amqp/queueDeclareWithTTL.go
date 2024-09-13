package amqp

import (
	log2 "github.com/e2site/sharks-go-lib/log"
	amqp "github.com/rabbitmq/amqp091-go"
)

func QueueDeclareWithTTL(queueName string, ttl int) *amqp.Queue {
	ch := GetChannel()
	q, err := ch.QueueDeclare(
		queueName,
		true,
		true,
		false,
		false,
		amqp.Table{
			amqp.QueueMessageTTLArg: ttl,
		},
	)
	log2.CheckEndLogFatal(err)
	return &q
}
