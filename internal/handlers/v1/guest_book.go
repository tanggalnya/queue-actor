package v1

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/tanggalnya/queue-actor/internal/domain"
	"github.com/tanggalnya/queue-actor/internal/events"
	"github.com/tanggalnya/queue-actor/internal/services/message_queue/publisher"
)

func guestBook(c *gin.Context) {
	cfg := publisher.Config{Uri: "amqp://guest:guest@localhost", QueueName: domain.EventTriggerTables.GuestBook, ExchangeName: "guest-book.events.topic"} //TODO: change this
	publisher := publisher.NewPublishEvent(cfg)
	attr := make(map[string]interface{})
	attr["hello"] = "world"
	str := events.TriggersEvent{
		BaseEvent: events.BaseEvent{
			Attributes: attr,
		},
		Type:  domain.EventTriggers.Insert,
		Table: domain.EventTriggerTables.GuestBook,
	}
	v, _ := json.Marshal(str)
	err := publisher.PublishEvent(string(v)) //TODO: change this
	if err != nil {
		return
	}

	c.JSON(200, gin.H{
		"success": "true",
	})
}
