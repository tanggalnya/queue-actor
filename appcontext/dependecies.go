package appcontext

import "github.com/tanggalnya/queue-actor/internal/services/message_queue/subscriber"

type EventSubscribers struct {
	GuestBook subscriber.EventSubscriber
}
