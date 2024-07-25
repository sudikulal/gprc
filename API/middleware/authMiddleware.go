package middleware

import (
	"journal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Auth(c *gin.Context) {
	accessToken := c.GetHeader("access_token")
	if accessToken == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "access token cannot be empty"})
		c.Abort()
		return
	}

	userObj, err := utils.DecodeToken(accessToken)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		c.Abort()
		return
	}

	c.Set("userObj", userObj)

	c.Next()
}
