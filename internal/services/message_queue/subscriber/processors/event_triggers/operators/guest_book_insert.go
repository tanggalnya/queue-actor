package operators

import (
	"context"
	"log"

	"github.com/tanggalnya/queue-actor/pkg/gdrive"
	"github.com/tanggalnya/queue-actor/pkg/spreadsheet"
)

type GuestBookInsertOperator struct {
	sc spreadsheet.Client
	dc gdrive.Client
}

func (g GuestBookInsertOperator) Process(ctx context.Context) error {
	folderID, err := g.dc.CreateFolder(ctx, "Data Wedding Komang & Kesha", "1HiDJYDb1Nx5Pe7TLWZHVMkx_NLa88c30")
	if err != nil {
		return err
	}
	file, err := g.dc.CreateFile(ctx, "Guest Book Komang & Kesha", folderID, gdrive.MimeTypes.CSV.GoogleMimeType())
	if err != nil {
		return err
	}
	_, err = g.sc.InsertRow(ctx, file.Id, nil)
	if err != nil {
		return err
	}
	log.Printf("Successfuly insert data to file: %s\n", file.Name)
	return nil
}

func (g GuestBookInsertOperator) AfterProcess() error {
	//TODO implement me
	return nil
}

func NewGuestBookInsertOperator(sc spreadsheet.Client, dc gdrive.Client) GuestBookInsertOperator {
	return GuestBookInsertOperator{
		sc: sc,
		dc: dc,
	}
}
