package main

import (
	"go-kanban/helper"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func main() {

	log.Info().Msg("server started!")
	routes := gin.Default()

	routes.GET("", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "welcome")
	})

	server := &http.Server{
		Addr:    ":3000",
		Handler: routes,
	}

	err := server.ListenAndServe()
	helper.ErrorHandler(err)
}
