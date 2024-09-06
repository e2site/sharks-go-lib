package amqp

import (
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"os"
)

var conAMQP *amqp.Connection
var chAMQP *amqp.Channel

func InitAMQP(dsn string) {
	conn, err := amqp.Dial(dsn)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	conAMQP = conn
	ch, errCh := conAMQP.Channel()
	if errCh != nil {
		fmt.Println(errCh)
		os.Exit(1)
	}
	chAMQP = ch
}

func GetChannel() *amqp.Channel {
	return chAMQP
}

func CloseAMQP() {
	chAMQP.Close()
	conAMQP.Close()
}
