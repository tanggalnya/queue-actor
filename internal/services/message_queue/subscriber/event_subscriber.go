package subscriber

type EventSubscriber interface {
	Process() error
}
