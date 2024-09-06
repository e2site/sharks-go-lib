package amqp

import (
	"github.com/e2site/sharks-go-lib/log"
	amqp "github.com/rabbitmq/amqp091-go"
)

var conAMQP *amqp.Connection
var chAMQP *amqp.Channel

func InitAMQP(dsn string) {
	conn, err := amqp.Dial(dsn)
	log.CheckEndLogFatal(err)
	conAMQP = conn
	ch, errCh := conAMQP.Channel()
	log.CheckEndLogFatal(errCh)
	chAMQP = ch
}

func GetChannel() *amqp.Channel {
	return chAMQP
}

func CloseAMQP() {
	chAMQP.Close()
	conAMQP.Close()
}
