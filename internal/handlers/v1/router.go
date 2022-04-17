package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/tanggalnya/queue-actor/internal/config"

	"github.com/tanggalnya/queue-actor/internal/services/message_queue/publisher"
)

func GuestBook(gb *gin.RouterGroup) {
	cgb := gb.Group("/guest-book", gin.BasicAuth(gin.Accounts{
		config.GuestBookConfig().Username: config.GuestBookConfig().Password,
	}))

	cfg := publisher.Config{
		Uri:          config.RabbitMQConfig().Uri,
		QueueName:    config.RabbitMQConfig().EventTriggersGuestBookQueueNames,
		ExchangeName: config.RabbitMQConfig().EventTriggersExchangeName,
	}
	pe := publisher.NewPublishEvent(cfg)

	cgb.POST("/", GuestBookHandler(pe))
}
