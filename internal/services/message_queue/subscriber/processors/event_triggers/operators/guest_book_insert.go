package operators

type GuestBookInsertOperator struct{}

func (g GuestBookInsertOperator) Process() error {
	//TODO implement me
	return nil
}

func (g GuestBookInsertOperator) AfterProcess() error {
	//TODO implement me
	return nil
}

func NewGuestBookInsertOperator() GuestBookInsertOperator {
	return GuestBookInsertOperator{}
}
