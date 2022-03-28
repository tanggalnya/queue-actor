package domain

var EventTriggers = newEventTriggersTypes()

type EventTriggerType string

type eventTriggerRegistry struct {
	GuestBook      EventTriggerType
	GuestAttending EventTriggerType
}

func newEventTriggersTypes() eventTriggerRegistry {
	return eventTriggerRegistry{
		GuestBook:      "guest_book",
		GuestAttending: "guest_attending",
	}
}
