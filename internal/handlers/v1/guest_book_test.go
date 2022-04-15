package v1

//import (
//	"net/http"
//	"net/http/httptest"
//	"testing"
//
//	"github.com/gin-gonic/gin"
//	"github.com/stretchr/testify/assert"
//	"github.com/stretchr/testify/mock"
//	"github.com/tanggalnya/queue-actor/internal/services/message_queue/mocks"
//)
//
//func TestGuestBookRoute(t *testing.T) {
//	w := httptest.NewRecorder()
//	c, _ := gin.CreateTestContext(w)
//	mq := new(mocks.MessageQueue)
//	mq.On("PublishEvent", mock.Anything).Return(nil)
//
//	guestBook(c)
//	assert.Equal(t, http.StatusOK, w.Code)
//	assert.JSONEq(t, `{"success": "true"}`, w.Body.String())
//}
