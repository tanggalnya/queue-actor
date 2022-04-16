package cmd

import (
	"context"
	"github.com/tanggalnya/queue-actor/appcontext"
	"github.com/tanggalnya/queue-actor/internal/config"
	"golang.org/x/oauth2/google"
	"gopkg.in/Iwark/spreadsheet.v2"
	"io/ioutil"
	"log"
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

func googleSheetClient() *spreadsheet.Service {
	sheetCfg := config.GoogleSpreadsheetConfig()
	data, err := ioutil.ReadFile(sheetCfg.FileSecretLocation)
	checkError(err)
	conf, err := google.JWTConfigFromJSON(data, spreadsheet.Scope)
	checkError(err)
	client := conf.Client(context.Background())

	sheetService := spreadsheet.NewServiceWithClient(client)
	return sheetService
}
