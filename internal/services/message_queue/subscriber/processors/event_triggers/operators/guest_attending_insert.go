package operators

type GuestAttendingInsertOperator struct{}

func (g GuestAttendingInsertOperator) Process() error {
	//TODO implement me
	return nil
}

func (g GuestAttendingInsertOperator) AfterProcess() error {
	//TODO implement me
	return nil
}

func NewGuestAttendingInsertOperator() GuestAttendingInsertOperator {
	return GuestAttendingInsertOperator{}
}
