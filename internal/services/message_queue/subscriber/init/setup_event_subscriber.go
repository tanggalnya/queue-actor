package init

import (
	"github.com/tanggalnya/queue-actor/appcontext"
	"github.com/tanggalnya/queue-actor/internal/config"
	"github.com/tanggalnya/queue-actor/internal/domain"
	"github.com/tanggalnya/queue-actor/internal/services/message_queue/subscriber"
	"github.com/tanggalnya/queue-actor/internal/services/message_queue/subscriber/processors/event_triggers"
	"github.com/tanggalnya/queue-actor/internal/services/message_queue/subscriber/processors/event_triggers/operators"
)

func Initialize(svc *appcontext.Services, ec *appcontext.ExternalClients) *appcontext.EventSubscribers {
	registerGuestBookOperators(svc)

	cfg := subscriber.Config{
		Uri:          config.RabbitMQConfig().Uri,
		ExchangeName: config.RabbitMQConfig().EventTriggersExchangeName,
		ExchangeKind: config.RabbitMQConfig().EventTriggersExchangeKind,
		ConsumerName: config.RabbitMQConfig().ConsumerName,
		QueueName:    config.RabbitMQConfig().EventTriggersGuestBookQueueNames,
		Processor:    event_triggers.NewProcessorFactory(),
	}
	gbt := subscriber.NewEventTriggersSubscriber(cfg)
	return &appcontext.EventSubscribers{Triggers: gbt}
}

func registerGuestBookOperators(s *appcontext.Services) {
	operators.RegisterEventOperator(domain.EventTriggers.Insert, operators.NewGuestBookInsertOperator(s.Spreadsheet, s.Drive))
}
