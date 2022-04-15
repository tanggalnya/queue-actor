package v1

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/tanggalnya/queue-actor/internal/domain"
	"github.com/tanggalnya/queue-actor/internal/events"
	"github.com/tanggalnya/queue-actor/internal/handlers/apiModels"
	"github.com/tanggalnya/queue-actor/internal/services/message_queue"
)

type getBookingHandler struct {
	ps message_queue.Service
}

func GuestBookHandler(p message_queue.Service) gin.HandlerFunc {
	handler := getBookingHandler{ps: p}
	return handler.guestBook
}

func (g getBookingHandler) guestBook(c *gin.Context) {
	var eventReq apiModels.HasuraEvent
	if err := c.Bind(&eventReq); err != nil {
		c.JSON(http.StatusBadRequest, apiModels.Response{Errors: []apiModels.Err{
			{Code: apiModels.ErrCodes.PayloadError, Message: fmt.Sprintf("Decode Payload got err: %v", err.Error())},
		}})
	}

	attr := make(map[string]interface{})
	attr["event"] = eventReq
	str := events.TriggersEvent{
		BaseEvent: events.BaseEvent{
			Attributes: attr,
		},
		Type:  domain.EventTriggers.Insert,
		Table: domain.EventTriggerTables.GuestBook,
	}
	v, _ := json.Marshal(str)
	err := g.ps.PublishEvent(string(v))
	if err != nil {
		c.JSON(http.StatusInternalServerError, apiModels.Response{Errors: []apiModels.Err{
			{Code: apiModels.ErrCodes.InternalError, Message: fmt.Sprintf("Publish Event got err: %v", err.Error())},
		}})
		return
	}

	c.JSON(http.StatusOK, apiModels.Response{Errors: []apiModels.Err{}})
}
