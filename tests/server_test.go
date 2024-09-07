package tests

import (
	client2 "github.com/e2site/sharks-go-lib/client"
	"github.com/e2site/sharks-go-lib/server"
	"github.com/gin-gonic/gin"
	"testing"
)

func TestServer(t *testing.T) {

	t.Setenv("OTEL_EXPORTER_OTLP_ENDPOINT", "http://localhost:4318")

	server.CreateServer("test", func(r *gin.Engine) {
		r.GET("/test", func(context *gin.Context) {
			client := client2.CreateClient("balance_req")
			client.SetHeaders(map[string]string{
				"TelegramAuth": " query_id=AAH1Nt4fAAAAAPU23h-f5c3U&user=%7B%22id%22%3A534656757%2C%22first_name%22%3A%22Filipp%22%2C%22last_name%22%3A%22Chelyshkov%22%2C%22language_code%22%3A%22ru%22%2C%22allows_write_to_pm%22%3Atrue%7D&auth_date=1722433483&hash=95df25a03aca85f1d8cfb280dc7e75de94c552ea2f9f4f7565988b73379917bc",
			})
			client.R().Get("http://localhost/api/balance")
			context.Status(200)
		})
	})
}
