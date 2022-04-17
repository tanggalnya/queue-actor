package config

import (
	"github.com/tanggalnya/queue-actor/pkg/spreadsheet"
)

var googleSpreadsheet spreadsheet.Config

func GoogleSpreadsheetConfig() spreadsheet.Config {
	return googleSpreadsheet
}

func initGoogleSpreadsheetConfig() {
	googleSpreadsheet = spreadsheet.Config{
		FileSecretLocation: getStringOrPanic("SPREADSHEET_FILE_SECRET_LOCATION"),
	}
}
