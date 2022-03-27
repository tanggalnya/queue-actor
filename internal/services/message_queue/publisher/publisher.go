package publisher

import (
	"github.com/tanggalnya/queue-actor/internal/services/message_queue"
	"log"

	"github.com/rabbitmq/amqp091-go"
	"github.com/wagslane/go-rabbitmq"
)

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

func (p publisherClient) PublishEvent(message string) error {
	publisher, err := rabbitmq.NewPublisher(p.uri, amqp091.Config{}, rabbitmq.WithPublisherOptionsLogging)
	if err != nil {
		log.Fatal(err)
	}
	err = publisher.Publish([]byte(message), []string{"routing_key"}, rabbitmq.WithPublishOptionsContentType("application/json"), rabbitmq.WithPublishOptionsMandatory, rabbitmq.WithPublishOptionsPersistentDelivery, rabbitmq.WithPublishOptionsExchange("events"))
	if err != nil {
		log.Fatal(err)
	}

	returns := publisher.NotifyReturn()
	go func() {
		for r := range returns {
			log.Printf("message returned from server: %s", string(r.Body))
		}
	}()
	return nil
}

func NewPublishEvent(cfg Config) message_queue.Publisher {
	return &publisherClient{
		uri:       cfg.Uri,
		queueName: cfg.QueueName,
		reliable:  cfg.Reliable,
	}
}
