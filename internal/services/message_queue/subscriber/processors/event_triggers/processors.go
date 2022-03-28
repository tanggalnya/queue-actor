package event_triggers

type Processor interface {
	GuestBookInsert() error
	GuestBookDelete() error
}
