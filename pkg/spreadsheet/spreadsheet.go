package spreadsheet

import (
	"context"
	"log"

	"google.golang.org/api/sheets/v4"
)

type Client interface {
	CreateSpreadsheet(context context.Context) (*sheets.Spreadsheet, error)
}

type Config struct {
	FileSecretLocation string
}

type spreadSheetClient struct {
	service *sheets.Service
}

func (s spreadSheetClient) CreateSpreadsheet(ctx context.Context) (*sheets.Spreadsheet, error) {
	rb := &sheets.Spreadsheet{
		Properties: &sheets.SpreadsheetProperties{
			Title: "sheet title",
		},
		// TODO: Add desired fields of the request body.
	}

	resp, err := s.service.Spreadsheets.Create(rb).Context(ctx).Do()
	if err != nil {
		log.Fatal(err)
	}

	// TODO: update permission using google drive API
	// https://github.com/Iwark/spreadsheet/issues/57#issuecomment-1096701536
	return resp, err
}

func New(ss sheets.Service) *spreadSheetClient {
	return &spreadSheetClient{
		service: &ss,
	}
}
