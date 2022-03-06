package v1

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestCreateGuestBookRoute(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	createGuestBook(c)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.JSONEq(t, `{"success": "true"}`, w.Body.String())
}
