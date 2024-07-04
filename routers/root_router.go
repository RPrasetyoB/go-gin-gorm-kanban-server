package routers

import (
	"go-kanban/http/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitRoutes(router *gin.Engine) {
	router.GET("/", func(ctx *gin.Context) {
		response := response.SuccessResponse{
			Success: true,
			Code:    http.StatusOK,
			Message: "Welcome to RPB Api",
			Version: "1.0.0",
		}
		ctx.JSON(http.StatusOK, response)
	})
}
