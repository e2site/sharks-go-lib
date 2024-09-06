package middleware

import (
	log2 "github.com/e2site/sharks-go-lib/log"
	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"github.com/opentracing/opentracing-go/log"
)

var span opentracing.Span

func GetCurrentSpan() opentracing.Span {
	return span
}

func TracingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tracer := opentracing.GlobalTracer() // Получаем глобальный трейсер

		// Извлечение контекста из заголовков запроса
		wireContext, err := tracer.Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(c.Request.Header))
		if err != nil && err != opentracing.ErrSpanContextNotFound {
			log2.Log(err)
		}

		// Создаем новый спан, продолжая цепочку если был передан контекст
		span = tracer.StartSpan(c.Request.Method+" "+c.FullPath(), ext.RPCServerOption(wireContext))
		defer span.Finish()

		// Добавляем контекст со спаном в запрос для последующего использования
		ctx := opentracing.ContextWithSpan(c.Request.Context(), span)
		c.Request = c.Request.WithContext(ctx)

		// Логируем начало обработки запроса
		span.LogFields(
			log.String("event", "request_started"),
			log.String("url", c.Request.URL.String()),
		)

		// Выполняем следующий обработчик
		c.Next()

		// Логируем конец обработки запроса
		span.LogFields(
			log.String("event", "request_finished"),
			log.Int("status_code", c.Writer.Status()),
		)

		// Добавляем теги для статуса ответа
		ext.HTTPStatusCode.Set(span, uint16(c.Writer.Status()))

		if len(c.Errors) > 0 {
			// Логируем ошибки, если они есть
			span.LogFields(log.String("error", c.Errors.String()))
			ext.Error.Set(span, true)
		}
	}
}
