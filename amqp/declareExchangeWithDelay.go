package amqp

import (
	"github.com/e2site/sharks-go-lib/log"
	amqp "github.com/rabbitmq/amqp091-go"
)

func DeclareExchangeWithDelay(exchangeName string) {
	ch := GetChannel()
	err := ch.ExchangeDeclare(
		exchangeName,
		"x-delayed-message",
		true,
		false,
		false,
		false,
		amqp.Table{
			"x-delayed-type": "direct",
		},
	)
	log.CheckEndLogFatal(err)
}
