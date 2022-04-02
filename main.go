package main

import (
	"log"
	"os"

	"github.com/tanggalnya/queue-actor/internal/handlers"
	enqueues "github.com/tanggalnya/queue-actor/internal/services/message_queue/subscriber/init"
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
		eventSubscribers := enqueues.Initialize()
		err := eventSubscribers.Triggers.Process()
		if err != nil {
			log.Panicln("Cannot spawn worker")
			return
		}
	default:
		log.Panicln("Need 1 args")
		return
	}
}
