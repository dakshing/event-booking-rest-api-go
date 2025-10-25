package middleware

import (
	"net/http"

	"github.com/dakshing/event-booking-rest-api-go/utils"
	"github.com/gin-gonic/gin"
)

func Authenticate(c *gin.Context) {
	token := c.GetHeader("Authorization")
	if token == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization token required"})
		return
	}

	loggedInUser, err := utils.ValidateTokenStringAndGetUserID(token)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
	}

	c.Set("userID", loggedInUser)
	c.Next()
}
