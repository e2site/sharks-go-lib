package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel/attribute"
	semconv "go.opentelemetry.io/otel/semconv/v1.10.0"

	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
)

func TracingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Создаем новый спан
		ctx := c.Request.Context()
		span := trace.SpanFromContext(ctx)

		// Добавляем заголовок с Trace ID в ответ
		traceId, existTrace := c.Get(X_TRACE_ID)
		if existTrace {
			span.SetAttributes(attribute.String("metadata.tracer", fmt.Sprintf("%s", traceId)))
		}

		telegrammId, existTelegrammId := c.Get(CONTEXT_TG_NAME)
		if existTelegrammId {

			span.SetAttributes(attribute.String("metadata.tguser", fmt.Sprintf("%s", telegrammId)))
		}

		// Логируем начало обработки запроса
		span.AddEvent("request_started", trace.WithAttributes(
			semconv.HTTPMethodKey.String(c.Request.Method),
			semconv.HTTPURLKey.String(c.Request.URL.String()),
		))

		// Выполняем следующий обработчик
		c.Next()

		// Логируем конец обработки запроса
		span.AddEvent("request_finished", trace.WithAttributes(
			semconv.HTTPStatusCodeKey.Int(c.Writer.Status()),
		))

		// Добавляем теги для статуса ответа
		span.SetAttributes(
			semconv.HTTPStatusCodeKey.Int(c.Writer.Status()),
		)

		if len(c.Errors) > 0 {
			// Логируем ошибки, если они есть
			span.AddEvent("error", trace.WithAttributes(
				semconv.HTTPStatusCodeKey.Int(c.Writer.Status()),
			))
			span.SetStatus(codes.Error, "Error")
		}

	}
}
