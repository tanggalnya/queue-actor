package main

import (
	"log"

	"github.com/tanggalnya/queue-actor/internal/services/message_queue/subscriber"
)

func main() {
	cfg := subscriber.Config{
		Uri:          "amqp://guest:guest@localhost",
		ExchangeName: "events",
		ExchangeKind: "topic",
		ConsumerName: "gdrive",
		QueueName:    "insert_google_drive",
	}
	consumer := subscriber.NewSubscribeEvent(cfg)
	err := consumer.Consume()
	if err != nil {
		log.Panicln("Cannot spawn worker")
		return
	}

}
