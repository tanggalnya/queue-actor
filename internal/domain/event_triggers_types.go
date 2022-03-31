package domain

var EventTriggers = newEventTriggersTypes()

type EventTriggerType string

type eventTriggerRegistry struct {
	Insert EventTriggerType
	Delete EventTriggerType
	Update EventTriggerType
}

func newEventTriggersTypes() eventTriggerRegistry {
	return eventTriggerRegistry{
		Insert: "insert",
		Delete: "delete",
		Update: "update",
	}
}
