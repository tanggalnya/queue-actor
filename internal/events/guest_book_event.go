package events

import "github.com/tanggalnya/queue-actor/internal/domain"

type GuestBookEvent struct {
	BaseEvent
	Type domain.EventTriggerType `json:"type"`
}
