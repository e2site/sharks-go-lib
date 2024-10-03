package amqp

import (
	"errors"
	"fmt"
	"github.com/e2site/sharks-go-lib/log"
	amqp "github.com/rabbitmq/amqp091-go"
	"time"
)

var conAMQP *amqp.Connection
var chAMQP *amqp.Channel

var statusConnect bool

var connectListChanel map[int]*chan bool
var idxConnection = 0

func SubscribeConnection(handlerChangeConnect func(status bool)) {
	c := make(chan bool)
	connectListChanel[idxConnection] = &c
	idxConnection++

	go func() {
		for {
			conStatus := <-c
			handlerChangeConnect(conStatus)
		}
	}()
}

func GetStatusConnected() bool {
	return statusConnect
}

func SendSubscribeStatus(status bool) {
	for _, s := range connectListChanel {
		select {
		case *s <- status:
			break
		default:
			break
		}
	}
}

func InitAMQP(dsn string, reconnectHandler func()) {
	conn, err := amqp.Dial(dsn)
	log.CheckEndLogFatal(err)
	conAMQP = conn
	ch, errCh := conAMQP.Channel()
	log.CheckEndLogFatal(errCh)
	chAMQP = ch

	statusConnect = true

	connectListChanel = make(map[int]*chan bool)

	statusConnection(reconnectHandler, func() bool {
		nCon, nErr := amqp.Dial(dsn)
		if nErr == nil {
			conAMQP = nCon
			nCh, nErrCh := conAMQP.Channel()
			if nErrCh == nil {
				chAMQP = nCh
				statusConnect = true
				SendSubscribeStatus(true)
				return true
			}
		}
		return false
	})
}

func GetChannel() *amqp.Channel {
	return chAMQP
}

func CloseAMQP() {
	chAMQP.Close()
	conAMQP.Close()
}

func statusConnection(reconnectHandler func(), newConnHandler func() bool) {
	go func() {
		for {
			reason, ok := <-conAMQP.NotifyClose(make(chan *amqp.Error))
			if !ok {
				log.Log(errors.New("rabbitmq connection closed"))
				break
			}
			log.Log(errors.New(fmt.Sprintf("rabbitmq connection closed unexpectedly, reason: %v", reason)))
			statusConnect = false
			SendSubscribeStatus(false)

			for {

				time.Sleep(time.Duration(1) * time.Second)

				if newConnHandler() {
					reconnectHandler()
					log.Log(errors.New("rabbitmq reconnect success"))
					break
				}
				log.Log(errors.New("rabbitmq reconnect failed"))
			}
		}
	}()
}
