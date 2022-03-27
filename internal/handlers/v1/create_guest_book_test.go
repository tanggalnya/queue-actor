package v1

import (
	"github.com/tanggalnya/queue-actor/internal/services/message_queue/publisher/mocks"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateGuestBookRoute(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	mq := new(mocks.MessageQueue)
	mq.On("PublishEvent", mock.Anything).Return(nil)

	createGuestBook(c)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.JSONEq(t, `{"success": "true"}`, w.Body.String())
}
