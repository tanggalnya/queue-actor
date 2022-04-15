package message_queue

type Service interface {
	PublishEvent(body string) error
}
