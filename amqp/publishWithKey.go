package amqp

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/e2site/sharks-go-lib/otl"
	amqp "github.com/rabbitmq/amqp091-go"
	"go.opentelemetry.io/otel"
)

func PublishMessageWithKey[Obj any](ctx context.Context, exchangeName string, object *Obj, key string) error {
	body, err := json.Marshal(object)
	if err != nil {
		return err
	}

	tr := otel.Tracer("amqp")
	amqpContext, messageSpan := tr.Start(ctx, fmt.Sprintf("AMQP - publish - %s", exchangeName))
	defer messageSpan.End()

	headers := otl.InjectAMQPHeaders(amqpContext)

	ch := GetChannel()
	errPub := ch.Publish(
		exchangeName,
		key,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
			Headers:     headers,
		},
	)

	return errPub
}
