package config

import (
	"go-kanban/controllers"
	"go-kanban/routers"

	"github.com/gin-gonic/gin"
)

type InitConfig struct {
	Router *gin.Engine
}

func Init() *InitConfig {
	db := DatabaseConnection()
	authController := controllers.NewAuthController(db)

	router := gin.Default()

	routers.InitRoutes(router)
	authRouter := routers.NewAuthRouter(router, authController)
	authRouter.RegisterRoutes()

	return &InitConfig{
		Router: router,
	}
}
