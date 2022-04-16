package operators

import (
	"context"
	"github.com/tanggalnya/queue-actor/pkg/spreadsheet"
)

type GuestBookInsertOperator struct {
	sc spreadsheet.Client
}

func (g GuestBookInsertOperator) BeforeProcess() error {
	//TODO implement me
	return nil
}

func (g GuestBookInsertOperator) Process(ctx context.Context) error {
	// TODO
	// Call Google Drive API for upsert folder (skip now)
	// Call Google sheet pkg
	// Create new sheet of guest book with row template
	// CRUD to that sheet
	return nil
}

func (g GuestBookInsertOperator) AfterProcess() error {
	//TODO implement me
	return nil
}

func NewGuestBookInsertOperator(sc spreadsheet.Client) GuestBookInsertOperator {
	return GuestBookInsertOperator{
		sc: sc,
	}
}
