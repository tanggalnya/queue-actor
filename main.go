package main

import (
	"log"
	"os"

	"github.com/tanggalnya/queue-actor/internal/handlers"
	"github.com/tanggalnya/queue-actor/internal/services/message_queue/subscriber"
)

var (
	server = "server"
	worker = "worker"
)

func main() {
	arg := os.Args[1]
	switch arg {
	case server:
		r := handlers.SetupRouter()
		err := r.Run(":8085")
		if err != nil {
			log.Panicln("Cannot spawn server")
			return
		}
	case worker:
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
	default:
		log.Panicln("Need 1 args")
		return
	}
}
