package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

const X_TRACE_ID = "X-Trace-Id"

func TraceIDMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {
		// Проверяем, есть ли уже X-Trace-Id в заголовках
		traceID := c.GetHeader(X_TRACE_ID)
		// Если заголовок отсутствует, создаем новый trace ID
		if traceID == "" {
			traceID = uuid.New().String()
			c.Writer.Header().Set(X_TRACE_ID, traceID) // Добавляем его в ответ
		}
		// Добавляем Trace ID в контекст для дальнейшего использования в других обработчиках
		c.Set(X_TRACE_ID, traceID)

		// Продолжаем обработку запроса
		c.Next()
	}
}
