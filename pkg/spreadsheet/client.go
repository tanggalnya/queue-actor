package spreadsheet

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/api/sheets/v4"
)

type Client interface {
	CreateSpreadsheet(ctx context.Context) (*sheets.Spreadsheet, error)
	InsertRow(ctx context.Context, sheetID string, valueRange *sheets.ValueRange) (*sheets.AppendValuesResponse, error)
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
	// How to work with google drive API
	// https://developers.google.com/sheets/api/guides/create#work_with_google_drive_folders
	// similar problem
	// https://stackoverflow.com/questions/63935583/create-update-google-drive-permission-in-go-created-spreadsheet-using-api-but
	return resp, err
}

func (s spreadSheetClient) InsertRow(ctx context.Context, sheetID string, valueRange *sheets.ValueRange) (*sheets.AppendValuesResponse, error) {
	//readRange := "Class Data!A1:A"
	//values := make(map[])
	//sv := &sheets.ValueRange{
	//	MajorDimension:  "",
	//	Range:           readRange,
	//	Values:          ,
	//	ServerResponse:  googleapi.ServerResponse{},
	//	ForceSendFields: nil,
	//	NullFields:      nil,
	//}
	//s.service.Spreadsheets.Values.Append(sheetID, readRange)

	// https://developers.google.com/sheets/api/reference/rest/v4/spreadsheets.values/append
	// The A1 notation of a range to search for a logical table of data.
	// Values will be appended after the last row of the table.
	range2 := "A1" // TODO: Update placeholder value.

	// How the input data should be interpreted.
	valueInputOption := "RAW" // TODO: Update placeholder value.

	// How the input data should be inserted.
	insertDataOption := "INSERT_ROWS" // TODO: Update placeholder value.

	// https://stackoverflow.com/a/34293641/4154982
	ss := "first"
	rest := []interface{}{"second", "two"}

	all := append([]interface{}{ss}, rest...)

	var matrix [][]interface{}
	matrix = append(matrix, all)
	fmt.Println(matrix)
	rb := &sheets.ValueRange{
		Values: matrix,
	}

	resp, err := s.service.Spreadsheets.Values.Append(sheetID, range2, rb).ValueInputOption(valueInputOption).InsertDataOption(insertDataOption).Context(ctx).Do()
	if err != nil {
		log.Fatal(err)
	}

	// TODO: Change code below to process the `resp` object:
	fmt.Printf("%#v\n", resp)
	return resp, nil
}

func New(ss sheets.Service) *spreadSheetClient {
	return &spreadSheetClient{
		service: &ss,
	}
}
