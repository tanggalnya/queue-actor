package message_queue

type Config struct {
	Uri       string
	QueueName string
	Reliable  bool
}

type publisherClient struct {
	uri       string
	queueName string
	reliable  bool
}

func (p publisherClient) PublishEvent(_ string) error {
	return nil
}

func NewPublishEvent(cfg Config) MessageQueue {
	return &publisherClient{
		uri:       cfg.Uri,
		queueName: cfg.QueueName,
		reliable:  cfg.Reliable,
	}
}
