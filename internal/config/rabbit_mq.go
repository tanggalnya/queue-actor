package config

import (
	"fmt"
)

var (
	rabbitMQPrefix      = "amqp://"
	rabbitMQUriTemplate = "%s%s:%s@%s"
)

type RabbitMQCfg struct {
	host                             string
	username                         string
	password                         string
	Uri                              string
	EventTriggersGuestBookQueueNames string
	EventTriggersExchangeName        string
	EventTriggersExchangeKind        string
	ConsumerName                     string
}

var rabbitMQ RabbitMQCfg

func RabbitMQConfig() RabbitMQCfg {
	return rabbitMQ
}

func initRabbitMQConfig() {
	rabbitMQ = RabbitMQCfg{
		host:                             getStringOrPanic("RABBIT_MQ_HOST"),
		username:                         getStringOrPanic("RABBIT_MQ_USERNAME"),
		password:                         getStringOrPanic("RABBIT_MQ_PASSWORD"),
		EventTriggersGuestBookQueueNames: getStringOrPanic("RABBIT_MQ_EVENT_TRIGGERS_GUEST_BOOK_QUEUE_NAME"),
		EventTriggersExchangeName:        getStringOrPanic("RABBIT_MQ_EVENT_TRIGGERS_EXCHANGE_NAME"),
		EventTriggersExchangeKind:        getStringOrPanic("RABBIT_MQ_EVENT_TRIGGERS_EXCHANGE_KIND"),
		ConsumerName:                     getStringOrPanic("RABBIT_MQ_CONSUMER_NAME"),
	}

	rabbitMQ.Uri = fmt.Sprintf(rabbitMQUriTemplate, rabbitMQPrefix, rabbitMQ.username, rabbitMQ.password, rabbitMQ.host)
}
