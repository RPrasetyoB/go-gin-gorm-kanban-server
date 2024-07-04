package utils

import (
	"github.com/gin-gonic/gin"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		err := c.Errors.Last()
		if err != nil {
			// Tangani error di sini
			switch e := err.Err.(type) {
			case *CustomError:
				c.JSON(e.Code, gin.H{
					"error": e.Message,
				})
			default:
				c.JSON(500, gin.H{
					"error": err.Error(),
				})
			}

			c.Abort()
		}
	}
}
