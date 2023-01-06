package controllers

import (
	"net/http"
	"strings"

	utils "personal-inventory/backend/utils"

	"github.com/gin-gonic/gin"
)

func Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		tokenArr := strings.Split(token, " ")

		if len(tokenArr) != 2 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid auth token"})
			c.Abort()
			return
		}

		claims, err := utils.ValidateToken(tokenArr[1])
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		c.Set("username", claims.Username)

		c.Next()
	}
}
