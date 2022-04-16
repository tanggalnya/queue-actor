package init

import (
	"github.com/tanggalnya/queue-actor/appcontext"
	"github.com/tanggalnya/queue-actor/internal/domain"
	"github.com/tanggalnya/queue-actor/internal/services/message_queue/subscriber"
	"github.com/tanggalnya/queue-actor/internal/services/message_queue/subscriber/processors/event_triggers"
	"github.com/tanggalnya/queue-actor/internal/services/message_queue/subscriber/processors/event_triggers/operators"
)

func Initialize(svc *appcontext.Services, ec *appcontext.ExternalClients) *appcontext.EventSubscribers {
	registerGuestBookOperators(svc)

	cfg := subscriber.Config{
		Uri:          "amqp://guest:guest@localhost",
		ExchangeName: "guest-book.events.topic",
		ExchangeKind: "topic",
		ConsumerName: "queue-actor",
		QueueName:    domain.EventTriggerTables.GuestBook,
		Processor:    event_triggers.NewProcessorFactory(),
	}
	ets := subscriber.NewEventTriggersSubscriber(cfg)
	return &appcontext.EventSubscribers{Triggers: ets}
}

func registerGuestBookOperators(s *appcontext.Services) {
	operators.RegisterEventOperator(domain.EventTriggers.Insert, operators.NewGuestBookInsertOperator(s.Spreadsheet))
}
