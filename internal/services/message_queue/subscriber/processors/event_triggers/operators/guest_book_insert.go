package operators

import (
	"context"
	"fmt"
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
	// TODO: change this
	//ss, err := g.sc.CreateSpreadsheet(ctx)
	ss, err := g.sc.InsertRow(ctx, "1WK7roc53yeTnZxA2neldVflbZVM7OG1Rk8kdD66CzVk", nil)
	fmt.Println(ss, err)
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
