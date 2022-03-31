package operators

type GuestAttendingInsertOperator struct{}

func (g GuestAttendingInsertOperator) Process() error {
	//TODO implement me
	panic("implement me")
}

func (g GuestAttendingInsertOperator) AfterProcess() error {
	//TODO implement me
	panic("implement me")
}

func NewGuestAttendingInsertOperator() GuestAttendingInsertOperator {
	return GuestAttendingInsertOperator{}
}
