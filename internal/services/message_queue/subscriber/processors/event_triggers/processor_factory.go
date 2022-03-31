package event_triggers

import "github.com/tanggalnya/queue-actor/internal/domain"

type Factory interface {
	NewProcessor(triggerType domain.EventTriggerType) Processor
}

type factory struct{}

func (f factory) NewProcessor(eventTrigger domain.EventTriggerType) Processor {
	switch eventTrigger {
	case domain.EventTriggers.GuestBook, domain.EventTriggers.GuestAttending:
		return newDefaultProcessor()
	default:
		panic("No default provider")
	}
}
