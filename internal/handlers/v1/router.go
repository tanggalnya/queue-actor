package v1

import "github.com/gin-gonic/gin"

func GuestBook(gb *gin.RouterGroup) {
	cgb := gb.Group("/guest-book")

	cgb.POST("/", createGuestBook)
}
