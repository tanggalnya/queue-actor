package init

import (
	"github.com/tanggalnya/queue-actor/appcontext"
	"github.com/tanggalnya/queue-actor/internal/domain"
	"github.com/tanggalnya/queue-actor/internal/services/message_queue/subscriber"
	"github.com/tanggalnya/queue-actor/internal/services/message_queue/subscriber/processors/event_triggers"
	"github.com/tanggalnya/queue-actor/internal/services/message_queue/subscriber/processors/event_triggers/operators"
)

func Initialize() *appcontext.EventSubscribers {
	registerGuestBookOperators()

	cfg := subscriber.Config{
		Uri:          "amqp://guest:guest@localhost",
		ExchangeName: "guest-book.events.topic",
		ExchangeKind: "topic",
		ConsumerName: "queue-actor",
		QueueName:    domain.EventTables.GuestBook,
		Processor:    event_triggers.NewProcessorFactory(domain.EventTables.GuestBook),
	}
	guestBookES := subscriber.NewEventSubscriber(cfg)
	return &appcontext.EventSubscribers{GuestBook: guestBookES}
}

func registerGuestBookOperators() {
	operators.RegisterEventOperator(domain.EventTriggers.Insert, operators.NewGuestBookInsertOperator())
}
