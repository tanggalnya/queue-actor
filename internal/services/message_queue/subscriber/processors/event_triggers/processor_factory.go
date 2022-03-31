package event_triggers

import "github.com/tanggalnya/queue-actor/internal/domain"

type Factory interface {
	NewProcessor(table domain.EventTable) Processor
}

type factory struct{}

func (f factory) NewProcessor(table domain.EventTable) Processor {
	switch table {
	case domain.EventTables.GuestBook, domain.EventTables.GuestAttending:
		return newDefaultProcessor(table)
	default:
		panic("No default provider")
	}
}
