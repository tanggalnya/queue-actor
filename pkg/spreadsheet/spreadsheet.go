package spreadsheet

import (
	"context"
	"gopkg.in/Iwark/spreadsheet.v2"
)

type Client interface {
	CreateSpreadsheet(context context.Context) (spreadsheet.Spreadsheet, error)
}

type Config struct {
	FileSecretLocation string
}

type spreadSheetClient struct {
	service *spreadsheet.Service
}

func (s spreadSheetClient) CreateSpreadsheet(context context.Context) (spreadsheet.Spreadsheet, error) {
	ss, err := s.service.CreateSpreadsheet(spreadsheet.Spreadsheet{
		Properties: spreadsheet.Properties{
			Title: "spreadsheet title",
		},
	})
	return ss, err
}

func New(ss spreadsheet.Service) *spreadSheetClient {
	return &spreadSheetClient{
		service: &ss,
	}
}
