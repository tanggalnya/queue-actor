package message_queue

type MessageQueue interface {
	PublishEvent(body string) error
}
