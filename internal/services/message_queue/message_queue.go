package message_queue

type Publisher interface {
	PublishEvent(body string) error
}

type Subscriber interface {
	Consume() error
}
