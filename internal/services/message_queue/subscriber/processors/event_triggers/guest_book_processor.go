package event_triggers

type guestBookProcessor struct{}

func (g guestBookProcessor) GuestBookInsert() error {
	//TODO implement me
	panic("implement me")
}

func (g guestBookProcessor) GuestBookDelete() error {
	//TODO implement me
	panic("implement me")
}

func newGuestBookProcessor() guestBookProcessor {
	return guestBookProcessor{}
}
