package amqp

import (
	"context"
	"github.com/e2site/sharks-go-lib/otl"
	"github.com/rabbitmq/amqp091-go"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
)

func AMQPDeliveryTracer(ctx context.Context, delivery *amqp091.Delivery, spanName string) trace.Span {
	ct := otl.ExtractAMQPHeaders(ctx, delivery.Headers)
	// Create a new span
	tr := otel.Tracer("amqp")
	_, messageSpan := tr.Start(ct, spanName)
	return messageSpan
}
