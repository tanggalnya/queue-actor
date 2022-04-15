package v1

import (
	"github.com/gin-gonic/gin"

	"github.com/tanggalnya/queue-actor/internal/domain"
	"github.com/tanggalnya/queue-actor/internal/services/message_queue/publisher"
)

func GuestBook(gb *gin.RouterGroup) {
	cgb := gb.Group("/guest-book", gin.BasicAuth(gin.Accounts{
		"admin": "admin123", //TODO: read from env
	}))

	cfg := publisher.Config{
		Uri:          "amqp://guest:guest@localhost", //TODO: read from env
		QueueName:    domain.EventTriggerTables.GuestBook,
		ExchangeName: "guest-book.events.topic", //TODO: read from env
	}
	pe := publisher.NewPublishEvent(cfg)

	cgb.POST("/", GuestBookHandler(pe))
}
