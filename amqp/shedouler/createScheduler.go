package shedouler

import (
	"github.com/e2site/sharks-go-lib/amqp"
	"github.com/e2site/sharks-go-lib/log"
	amqp2 "github.com/rabbitmq/amqp091-go"
)

const KEY_SCHEDULER = "scheduler"

func CreateScheduler(runExchanger string, runQueue string) {
	schedoulerQueue := getSchedulerName(runQueue)

	ch := amqp.GetChannel()
	amqp.DeclareExchange(runExchanger)

	schQ, errQ := ch.QueueDeclare(
		schedoulerQueue,
		true,
		false,
		false,
		false,
		amqp2.Table{
			"x-dead-letter-exchange":    runExchanger,
			"x-dead-letter-routing-key": "run",
		},
	)
	log.CheckEndLogFatal(errQ)
	rQ := amqp.QueueDeclare(runQueue)
	runQ := ch.QueueBind(
		rQ.Name,
		"run",
		runExchanger,
		false,
		nil,
	)
	log.CheckEndLogFatal(runQ)

	errBindQSch := ch.QueueBind(
		schQ.Name,
		KEY_SCHEDULER,
		runExchanger,
		false,
		nil,
	)

	log.CheckEndLogFatal(errBindQSch)

}

func getSchedulerName(exchangerName string) string {
	return "schedule_" + exchangerName
}
