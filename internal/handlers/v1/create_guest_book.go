package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/tanggalnya/queue-actor/internal/services/message_queue"
)

func createGuestBook(c *gin.Context) {
	cfg := message_queue.Config{Uri: "amqp://guest:guest@localhost"} //TODO: change this
	publisher := message_queue.NewPublishEvent(cfg)
	err := publisher.PublishEvent("events payload") //TODO: change this
	if err != nil {
		return
	}

	c.JSON(200, gin.H{
		"success": "true",
	})
}
