package main

import (
	"github.com/tanggalnya/queue-actor/cmd"
	"log"
	"os"

	"github.com/tanggalnya/queue-actor/internal/handlers"
)

var (
	server = "server"
	worker = "worker"
)

func main() {
	currentPath, err := os.Getwd()
	if err != nil {
		log.Panic(err)
	}

	arg := os.Args[1]
	switch arg {
	case server:
		r := handlers.SetupRouter()
		err := r.Run(":8085")
		if err != nil {
			log.Panicf("Cannot spawn server, err: %v", err.Error())
			return
		}
	case worker:
		err := cmd.InitWorkerBase(currentPath)
		if err != nil {
			log.Panicf("Cannot spawn worker, err: %v", err.Error())
			return
		}
		wd, err := cmd.InitWorkerDependencies()
		err = wd.EventsSubscribers.Triggers.Process()
		if err != nil {
			log.Panicf("Cannot spawn worker, err: %v", err.Error())
			return
		}

	default:
		log.Panicln("Need 1 args")
		return
	}
}
