package event_triggers

import (
	"context"
	"log"

	"github.com/tanggalnya/queue-actor/internal/domain"
	"github.com/tanggalnya/queue-actor/internal/services/message_queue/subscriber/processors/event_triggers/operators"
)

type guestBookProcessor struct {
	operator operators.Operator
}

func (g guestBookProcessor) Insert() error {
	log.Println("guestBookProcessor Insert data")

	oi := operators.Get(domain.EventTriggers.Insert)
	err := g.process(oi)
	if err != nil {
		log.Fatalf("Error in guestBookProcessor Insert. Error %s", err.Error())
	}
	return nil
}

func (g guestBookProcessor) Delete() error {
	log.Println("guestBookProcessor Delete data")

	od := operators.Get(domain.EventTriggers.Delete)
	err := g.process(od)
	if err != nil {
		log.Fatalf("Error in guestBookProcessor Insert. Error %s", err.Error())
	}
	return nil
}

func (g guestBookProcessor) Update() error {
	log := log.Logger{}
	log.Println("guestBookProcessor Update data")

	ou := operators.Get(domain.EventTriggers.Update)
	err := g.process(ou)
	if err != nil {
		log.Fatalf("Error in guestBookProcessor Update. Error %s", err.Error())
	}
	return nil
}

func (g guestBookProcessor) process(op operators.Operator) error {
	var err error

	err = op.Process(context.Background())
	if err != nil {
		return err
	}

	err = op.AfterProcess()
	if err != nil {
		return err
	}
	return nil
}

func newGuestBookProcessor(trigger domain.EventTriggerType) guestBookProcessor {
	var op operators.Operator
	switch trigger {
	case domain.EventTriggers.Insert:
		op = operators.GuestBookInsertOperator{}
	}

	return guestBookProcessor{
		operator: op,
	}
}
