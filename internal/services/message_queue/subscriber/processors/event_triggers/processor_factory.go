package event_triggers

import (
	"github.com/tanggalnya/queue-actor/internal/domain"
)

type Factory interface {
	NewProcessor(table domain.EventTriggerTable, trigger domain.EventTriggerType) Processor
}

type factory struct{}

func (f factory) NewProcessor(table domain.EventTriggerTable, trigger domain.EventTriggerType) Processor {
	switch table {
	case domain.EventTriggerTables.GuestBook:
		return newGuestBookProcessor(trigger)
	default:
		panic("No default provider")
	}
}

func NewProcessorFactory() Factory {
	return factory{}
}
