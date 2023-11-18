package middleware

import (
	"models"
	"net/http"
	"github.com/gin-gonic/gin"
)

func IsAdmin(c *gin.Context) {
	
	user := c.MustGet("user").(models.Users)
	
	if user.Role != "admin" {
		onErr(c)
		return
	}
	c.Next()
}

func onErr(c *gin.Context) {
	c.Redirect(http.StatusPermanentRedirect, "/v/home") 
	c.Abort()
}