package domain

var EventTables = newEventTables()

type EventTable string

type eventTablesRegistry struct {
	GuestBook      EventTable
	GuestAttending EventTable
}

func newEventTables() eventTablesRegistry {
	return eventTablesRegistry{
		GuestBook:      "guest_book",
		GuestAttending: "guest_attending",
	}
}
