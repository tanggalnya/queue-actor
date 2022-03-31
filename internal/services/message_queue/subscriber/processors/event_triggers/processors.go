package event_triggers

type Processor interface {
	Insert() error
	Delete() error
	Update() error
}
