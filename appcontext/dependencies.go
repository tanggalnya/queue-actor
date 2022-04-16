package appcontext

import (
	"github.com/tanggalnya/queue-actor/internal/services/message_queue/subscriber"
	"github.com/tanggalnya/queue-actor/pkg/spreadsheet"
	googleSheet "gopkg.in/Iwark/spreadsheet.v2"
)

type EventSubscribers struct {
	Triggers subscriber.EventSubscriber
}

type ExternalClients struct {
	GoogleSheetClient googleSheet.Service
}

type Services struct {
	Spreadsheet spreadsheet.Client
}

type ServerDependency struct {
	Services *Services
}

type WorkerDependencies struct {
	EventsSubscribers *EventSubscribers
	Services          *Services
	ExternalClients   *ExternalClients
}
