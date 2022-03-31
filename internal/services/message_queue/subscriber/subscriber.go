package subscriber

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/tanggalnya/queue-actor/internal/domain"
	"github.com/tanggalnya/queue-actor/internal/services/message_queue"
	"github.com/tanggalnya/queue-actor/internal/services/message_queue/subscriber/processors/event_triggers"
	"github.com/wagslane/go-rabbitmq"
)

type Config struct {
	Uri          string
	ExchangeName string
	ExchangeKind string
	ConsumerName string
	QueueName    string
}

type subscriberClient struct {
	uri          string
	exchangeName string
	exchangeKind string
	consumerName string
	queueName    string
	processor    event_triggers.Factory
}

func (s subscriberClient) Consume() error {
	consumer, err := rabbitmq.NewConsumer(s.uri, amqp.Config{}, rabbitmq.WithConsumerOptionsLogging)
	if err != nil {
		log.Fatal(err)
	}

	// wait for server to acknowledge the cancel
	const noWait = false
	defer consumer.Disconnect()
	defer consumer.StopConsuming(s.consumerName, noWait)

	err = consumer.StartConsuming(func(d rabbitmq.Delivery) rabbitmq.Action {
		log.Printf("consumed: %v", string(d.Body))
		// rabbitmq.Ack, rabbitmq.NackDiscard, rabbitmq.NackRequeue

		pr := s.processor.NewProcessor(domain.EventTables.GuestBook) //TODO: change this
		err := pr.Insert()
		if err != nil {
			return 0
		}

		return rabbitmq.Ack
	}, s.queueName, []string{"routing_key", "routing_key_2"}, rabbitmq.WithConsumeOptionsConcurrency(10), rabbitmq.WithConsumeOptionsQueueDurable, rabbitmq.WithConsumeOptionsQuorum, rabbitmq.WithConsumeOptionsBindingExchangeName(s.exchangeName), rabbitmq.WithConsumeOptionsBindingExchangeKind(s.exchangeKind), rabbitmq.WithConsumeOptionsBindingExchangeDurable, rabbitmq.WithConsumeOptionsConsumerName(s.consumerName))
	if err != nil {
		log.Fatal(err)
	}

	// block main thread - wait for shutdown signal
	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-sigs
		fmt.Println()
		fmt.Println(sig)
		done <- true
	}()

	fmt.Println("awaiting signal")
	<-done
	fmt.Println("stopping consumer")
	return nil
}

func NewSubscribeEvent(cfg Config) message_queue.Subscriber {
	return &subscriberClient{
		uri:          cfg.Uri,
		exchangeName: cfg.ExchangeName,
		exchangeKind: cfg.ExchangeKind,
		consumerName: cfg.ConsumerName,
		queueName:    cfg.QueueName,
	}
}
