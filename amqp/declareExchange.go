package amqp

import (
	"github.com/e2site/sharks-go-lib/log"
	amqp "github.com/rabbitmq/amqp091-go"
)

func DeclareExchange(exchangeName string) {
	ch := GetChannel()
	err := ch.ExchangeDeclare(
		exchangeName,
		amqp.ExchangeDirect,
		true,
		false,
		false,
		false,
		nil,
	)
	log.CheckEndLogFatal(err)
}
