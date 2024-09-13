package tests

import (
	"fmt"
	"github.com/e2site/sharks-go-lib/amqp"
	"github.com/e2site/sharks-go-lib/amqp/status"
	"testing"
	"time"
)

func TestFanout(t *testing.T) {
	amqp.InitAMQP("amqp://guest:guest@host.docker.internal:5672/")
	defer amqp.CloseAMQP()

	status.ReadStatus(func(message *status.StatusMessage) {
		fmt.Println(message)
	})

	status.PublishStatus(45676, "test")
	status.PublishStatus(45676, "test 2")

	time.Sleep(10 * time.Second)
}
