package status

import (
	"github.com/e2site/sharks-go-lib/amqp"
	"github.com/e2site/sharks-go-lib/log"
)

const StatusExchanger = "status-ws-exchanger"

var statusInit = false

type StatusMessage struct {
	TelegramId int64
	Message    string
}

func PublishStatus(telegramId int64, message string) {
	if !statusInit {
		amqp.DeclareFanout(StatusExchanger)
		statusInit = true
	}
	var newMsg StatusMessage
	newMsg.TelegramId = telegramId
	newMsg.Message = message
	err := amqp.PublishMessageWithoutTracer(StatusExchanger, &newMsg)
	if err != nil {
		log.Log(err)
	}

}
