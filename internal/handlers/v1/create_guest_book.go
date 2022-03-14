package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/tanggalnya/queue-actor/internal/services/message_queue"
)

func createGuestBook(c *gin.Context) {
	cfg := message_queue.Config{Uri: "http://fillme"} //TODO: change this
	_ = message_queue.NewPublishEvent(cfg)            //TODO: change this

	c.JSON(200, gin.H{
		"success": "true",
	})
}
