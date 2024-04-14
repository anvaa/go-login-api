package middleware

import (
	"controllers"
	"models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func IsAuth(c *gin.Context) {
	
	user := c.MustGet("user").(models.Users)
	
	if !user.IsAuth {
		onErrAuth(c)
		return
	}
	c.Next()
}

func onErrAuth(c *gin.Context) {
	controllers.Logout(c)
	c.Redirect(http.StatusPermanentRedirect, "/") 
	c.Abort()
}