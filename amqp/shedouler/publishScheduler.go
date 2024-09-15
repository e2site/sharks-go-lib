package shedouler

import (
	"encoding/json"
	"github.com/e2site/sharks-go-lib/amqp"
	amqp2 "github.com/rabbitmq/amqp091-go"
	"strconv"
	"time"
)

func PublishScheduler[Obj any](exchanger string, object *Obj, scTime time.Time) error {
	body, err := json.Marshal(object)
	if err != nil {
		return err
	}

	ttl := scTime.Sub(time.Now()).Milliseconds()

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
