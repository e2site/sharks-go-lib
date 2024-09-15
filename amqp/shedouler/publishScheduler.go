package shedouler

import (
	"encoding/json"
	"github.com/e2site/sharks-go-lib/amqp"
	amqp2 "github.com/rabbitmq/amqp091-go"
	"strconv"
	"time"
)

var cntRnd int64 = 1

func PublishScheduler[Obj any](exchanger string, object *Obj, scTime time.Time) error {
	body, err := json.Marshal(object)
	if err != nil {
		return err
	}

	if cntRnd > 999 {
		cntRnd = 1
	}

	ttl := scTime.Sub(time.Now()).Milliseconds() + cntRnd

	cntRnd++

	if ttl < 0 {
		ttl = 1
	}

	ch := amqp.GetChannel()
	errPub := ch.Publish(
		exchanger,
		KEY_SCHEDULER,
		false,
		false,
		amqp2.Publishing{
			ContentType: "application/json",
			Body:        body,
			Expiration:  strconv.FormatInt(ttl, 10),
		},
	)

	return errPub
}
