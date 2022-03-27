package subscriber

import "github.com/tanggalnya/queue-actor/internal/services/message_queue"

type Config struct {
	Uri          string
	ExchangeName string
	ExchangeKind string
	ConsumerName string
}

type subscriberClient struct {
	uri          string
	exchangeName string
	exchangeKind string
	consumerName string
}

func (s subscriberClient) Consume() error {
	//TODO implement me
	panic("implement me")
}

func NewSubscribeEvent(cfg Config) message_queue.Subscriber {
	return &subscriberClient{
		uri:          cfg.Uri,
		exchangeName: cfg.ExchangeName,
		exchangeKind: cfg.ExchangeKind,
		consumerName: cfg.ConsumerName,
	}
}
