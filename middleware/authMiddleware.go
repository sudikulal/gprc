package middleware

import (
	"journal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Auth(c *gin.Context) {
		accessToken := c.GetHeader("access_token")
		if accessToken == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "access token cannot be empty"})
			c.Abort()
			return
		}

		userId, err := utils.DecodeToken(accessToken)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		c.Set("userId", userId)

		c.Next()
	}

