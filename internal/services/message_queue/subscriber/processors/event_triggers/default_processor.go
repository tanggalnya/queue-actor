package event_triggers

import (
	"log"

	"github.com/tanggalnya/queue-actor/internal/domain"
	"github.com/tanggalnya/queue-actor/internal/services/message_queue/subscriber/processors/event_triggers/operators"
)

type defaultProcessor struct {
	operator operators.Operator
}

func (g defaultProcessor) Insert() error {
	logger := log.Logger{}
	logger.Println("defaultProcessor Insert data")

	oi := operators.Get(domain.EventTriggers.Insert)
	err := g.process(oi)
	if err != nil {
		logger.Fatal("Error in defaultProcessor Insert. Error %s", err.Error())
	}
	return nil
}

func (g defaultProcessor) Delete() error {
	logger := log.Logger{}
	logger.Println("defaultProcessor Delete data")

	od := operators.Get(domain.EventTriggers.Delete)
	err := g.process(od)
	if err != nil {
		logger.Fatal("Error in defaultProcessor Insert. Error %s", err.Error())
	}
	return nil
}

func (g defaultProcessor) Update() error {
	logger := log.Logger{}
	logger.Println("defaultProcessor Update data")

	ou := operators.Get(domain.EventTriggers.Update)
	err := g.process(ou)
	if err != nil {
		logger.Fatal("Error in defaultProcessor Update. Error %s", err.Error())
	}
	return nil
}

func (g defaultProcessor) process(op operators.Operator) error {
	err := op.Process()
	if err != nil {
		return err
	}
	return nil
}

func newDefaultProcessor(table domain.EventTable) defaultProcessor {
	var op operators.Operator
	switch table {
	case domain.EventTables.GuestBook:
		op = operators.GuestBookInsertOperator{}
	case domain.EventTables.GuestAttending:
		op = operators.GuestAttendingInsertOperator{}
	}

	return defaultProcessor{
		operator: op,
	}
}
