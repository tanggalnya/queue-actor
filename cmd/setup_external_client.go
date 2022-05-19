package cmd

import (
	"context"
	"log"

	"google.golang.org/api/drive/v3"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"

	"github.com/tanggalnya/queue-actor/appcontext"
	"github.com/tanggalnya/queue-actor/internal/config"
)

func InitializeExternalClients() (*appcontext.ExternalClients, error) {
	return &appcontext.ExternalClients{
		GoogleSheetClient: *googleSheetClient(),
		GoogleDriveClient: *googleDriveClient(),
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

func googleDriveClient() *drive.Service {
	driveCfg := config.GoogleDriveConfig()
	driveService, err := drive.NewService(context.Background(), option.WithCredentialsFile(driveCfg.FileSecretLocation))
	checkError(err)
	return driveService
}
