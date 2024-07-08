package middleware

import (
	"context"
	"go-kanban/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type contextKey string

const userContextKey contextKey = "user"

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		decodedToken, err := utils.GetToken(c)
		if err != nil || decodedToken == nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"success": false,
				"message": "Unauthorized, please login",
			})
			c.Abort()
			return
		}

		user := utils.LoggedUser(decodedToken)
		if user == nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"success": false,
				"message": "Invalid token data",
			})
			c.Abort()
			return
		}

		ctx := context.WithValue(c.Request.Context(), userContextKey, user)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}

// Helper function to get the user from the Gin context
func GetUser(c *gin.Context) map[string]interface{} {
	user, ok := c.Request.Context().Value(userContextKey).(map[string]interface{})
	if !ok {
		return nil
	}
	return user
}
