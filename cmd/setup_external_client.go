package cmd

import (
	"context"
	"log"

	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"

	"github.com/tanggalnya/queue-actor/appcontext"
	"github.com/tanggalnya/queue-actor/internal/config"
)

func InitializeExternalClients() (*appcontext.ExternalClients, error) {
	return &appcontext.ExternalClients{
		GoogleSheetClient: *googleSheetClient(),
	}, nil
}

func checkError(err error) {
	if err != nil {
		log.Panicf("Cannot init external client, err: %v", err.Error())
	}
}

func googleSheetClient() *sheets.Service {
	sheetCfg := config.GoogleSpreadsheetConfig()
	sheetsService, err := sheets.NewService(context.Background(), option.WithCredentialsFile(sheetCfg.FileSecretLocation))
	checkError(err)
	return sheetsService
}
