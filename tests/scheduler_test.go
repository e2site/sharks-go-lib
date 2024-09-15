package tests

import (
	"github.com/e2site/sharks-go-lib/amqp"
	"github.com/e2site/sharks-go-lib/amqp/shedouler"
	"testing"
	"time"
)

func TestScheduler(t *testing.T) {
	amqp.InitAMQP("amqp://guest:guest@host.docker.internal:5672/")
	defer amqp.CloseAMQP()
	shedouler.CreateScheduler("test_scheduler", "test_sc_queue")
	shedouler.PublishScheduler("test_scheduler", &map[string]string{
		"Test": "teeet",
	}, time.Now().Add(34*time.Second))
}
