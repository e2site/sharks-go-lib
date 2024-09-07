package tests

import (
	"github.com/e2site/sharks-go-lib/server"
	"github.com/gin-gonic/gin"
	"testing"
	"time"
)

func TestServer(t *testing.T) {

	t.Setenv("OTEL_EXPORTER_OTLP_ENDPOINT", "http://localhost:4318")

	server.CreateServer("test", "localhost:6831", func(r *gin.Engine) {
		r.GET("/test", func(context *gin.Context) {
			time.Sleep(1000)
			context.Status(200)
		})
	})
}
