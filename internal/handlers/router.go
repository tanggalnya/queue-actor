package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/tanggalnya/queue-actor/internal/handlers/ping"
	"github.com/tanggalnya/queue-actor/internal/handlers/v1"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", ping.Ping)

	v1Group := r.Group("/v1")
	v1.GuestBook(v1Group)

	return r
}
