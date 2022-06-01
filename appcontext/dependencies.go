package appcontext

import (
	"google.golang.org/api/drive/v3"
	"google.golang.org/api/sheets/v4"

	"github.com/tanggalnya/queue-actor/internal/services/message_queue/subscriber"
	"github.com/tanggalnya/queue-actor/pkg/gdrive"
	"github.com/tanggalnya/queue-actor/pkg/spreadsheet"
)

type EventSubscribers struct {
	Triggers subscriber.EventSubscriber
}

type ExternalClients struct {
	GoogleSheetClient sheets.Service
	GoogleDriveClient drive.Service
}

type Services struct {
	Spreadsheet spreadsheet.Client
	Drive       gdrive.Client
}

type ServerDependency struct {
	Services *Services
}

type WorkerDependencies struct {
	EventsSubscribers *EventSubscribers
	Services          *Services
	ExternalClients   *ExternalClients
}
