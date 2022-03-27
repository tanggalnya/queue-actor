package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/tanggalnya/queue-actor/internal/services/message_queue/publisher"
)

func createGuestBook(c *gin.Context) {
	cfg := publisher.Config{Uri: "amqp://guest:guest@localhost"} //TODO: change this
	publisher := publisher.NewPublishEvent(cfg)
	err := publisher.PublishEvent("events payload") //TODO: change this
	if err != nil {
		return
	}

	c.JSON(200, gin.H{
		"success": "true",
	})
}
