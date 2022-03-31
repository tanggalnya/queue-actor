package operators

type GuestBookInsertOperator struct{}

func (g GuestBookInsertOperator) Process() error {
	//TODO implement me
	panic("implement me")
}

func (g GuestBookInsertOperator) AfterProcess() error {
	//TODO implement me
	panic("implement me")
}

func NewGuestBookInsertOperator() GuestBookInsertOperator {
	return GuestBookInsertOperator{}
}
