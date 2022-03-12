package main

import (
	"log"

	"github.com/tanggalnya/queue-actor/internal/handlers"
)

func main() {
	r := handlers.SetupRouter()
	err := r.Run(":8085")
	if err != nil {
		log.Panicln("Cannot spawn server")

		return
	}
}