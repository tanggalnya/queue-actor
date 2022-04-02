package events

import "github.com/tanggalnya/queue-actor/internal/domain"

type TriggersEvent struct {
	BaseEvent
	Type  domain.EventTriggerType  `json:"type"`
	Table domain.EventTriggerTable `json:"table"`
}
