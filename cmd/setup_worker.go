package cmd

import (
	"github.com/tanggalnya/queue-actor/appcontext"
	"github.com/tanggalnya/queue-actor/internal/config"
	subscriber "github.com/tanggalnya/queue-actor/internal/services/message_queue/subscriber/init"
	"github.com/tanggalnya/queue-actor/pkg/spreadsheet"
)

func InitWorkerBase(envPath string) error {
	if err := config.LoadWorker(envPath); err != nil {
		return err
	}
	return nil
}

func InitWorkerDependencies() (*appcontext.WorkerDependencies, error) {
	extClients, err := InitializeExternalClients()
	if err != nil {
		return nil, err
	}
	services := setupWorkerServices(extClients)
	eventsSubscriber := subscriber.Initialize(services, extClients)

	return &appcontext.WorkerDependencies{
		EventsSubscribers: eventsSubscriber,
		Services:          services,
		ExternalClients:   extClients,
	}, nil
}

func setupWorkerServices(extClient *appcontext.ExternalClients) *appcontext.Services {
	spreadsheetService := spreadsheet.New(extClient.GoogleSheetClient)

	return &appcontext.Services{
		Spreadsheet: spreadsheetService,
	}
}
