package middleware

import (
	"models"
	"net/http"
	"github.com/gin-gonic/gin"
)

func IsAdmin(c *gin.Context) {
	// Get the user from the context
	user := c.MustGet("user").(models.Users)
	
	if user.Role != "admin" {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})
		return
	}
	
	c.Next()
}