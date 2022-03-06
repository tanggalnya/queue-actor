package v1

import "github.com/gin-gonic/gin"

func createGuestBook(c *gin.Context) {
	// TODO: put logic here
	c.JSON(200, gin.H{
		"success": "true",
	})
}
