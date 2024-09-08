package amqp

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/e2site/sharks-go-lib/otl"
	amqp "github.com/rabbitmq/amqp091-go"
	"go.opentelemetry.io/otel"
)

func PublishWithDelay[Obj any](ctx context.Context, exchangeName string, object *Obj, delay int) error {

	tr := otel.Tracer("amqp")
	amqpContext, messageSpan := tr.Start(ctx, fmt.Sprintf("AMQP - publish - %s", exchangeName))
	defer messageSpan.End()

	headers := otl.InjectAMQPHeaders(amqpContext)

	headers["x-delay"] = delay * 1000 * 60

	body, err := json.Marshal(object)
	if err != nil {
		return err
	}

	ch := GetChannel()
	errPub := ch.Publish(
		exchangeName,
		"",
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Headers:     headers,
			Body:        body,
		},
	)

	return errPub
}
