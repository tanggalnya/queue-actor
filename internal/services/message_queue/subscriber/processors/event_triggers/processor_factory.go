package event_triggers

import "github.com/tanggalnya/queue-actor/internal/domain"

type factory struct{}

func (f factory) NewProcessor(tableName domain.EventTriggerType) Processor {
	switch tableName {
	case domain.EventTriggers.GuestBook:
		return newGuestBookProcessor()
	default:
		panic("No default provider")
	}
}
