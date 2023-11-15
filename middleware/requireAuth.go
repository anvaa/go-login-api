package middleware

import (
	"initializers"
	"models"
	"net/http"
	"github.com/gin-gonic/gin"
	"os"	
	"github.com/golang-jwt/jwt/v5"
)

func RequireAuth(c *gin.Context) {
	// Get the JWT string from the header
	tokenString, err := c.Cookie("goAuth")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})
		return
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil}, 
	)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})
		return
	}
	
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})
		return
	}

	// Get the user from the database
	var user models.Users
	initializers.DB.Where("id = ?", claims["sub"]).First(&user)

	if user.Id == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})
		return
	}
	
	// Attach the user to the context
	c.Set("user", user)
	c.Next()
}