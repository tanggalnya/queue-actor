package config

import "github.com/tanggalnya/queue-actor/pkg/spreadsheet"

var googleSpreadsheet spreadsheet.Config

func GoogleSpreadsheetConfig() spreadsheet.Config {
	return googleSpreadsheet
}

func initGoogleSpreadsheetConfig() {
	googleSpreadsheet = spreadsheet.Config{
		FileSecretLocation: "/Users/ardinusawan/Coding/tanggalnya/queue-actor/credentials.json",
	}
}
