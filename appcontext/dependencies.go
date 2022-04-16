package appcontext

import (
	"google.golang.org/api/sheets/v4"

	"github.com/tanggalnya/queue-actor/internal/services/message_queue/subscriber"
	"github.com/tanggalnya/queue-actor/pkg/spreadsheet"
)

type EventSubscribers struct {
	Triggers subscriber.EventSubscriber
}

type ExternalClients struct {
	GoogleSheetClient sheets.Service
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
