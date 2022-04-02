package domain

var EventTriggerTables = newEventTriggerTables()

type EventTriggerTable string

type eventTriggerTablesRegistry struct {
	GuestBook      EventTriggerTable
	GuestAttending EventTriggerTable
}

func newEventTriggerTables() eventTriggerTablesRegistry {
	return eventTriggerTablesRegistry{
		GuestBook:      "guest_book",
		GuestAttending: "guest_attending",
	}
}
