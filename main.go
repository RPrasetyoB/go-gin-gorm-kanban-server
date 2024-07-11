package main

import (
	"fmt"
	"go-kanban/config"
	"go-kanban/helper"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var port = 3000

func main() {
	gin.SetMode(gin.ReleaseMode)
	// Load .env file
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Printf(".env load failed")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = ":8888"
	}
	// Initialize application configuration
	initConfig := config.Init()

	// Use custom error handler middleware
	initConfig.Router.Use(helper.ErrorHandler())

	go func() {
		if err := initConfig.Router.Run(port); err != nil {
			fmt.Printf("Failed to start server: %v\n", err)
		}
	}()

	fmt.Printf("Server running at http://localhost%s\n", port)
	select {}
}
