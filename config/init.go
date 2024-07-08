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
	db, _ := DatabaseConnection()
	authController := controllers.NewAuthController(db)
	todoController := controllers.NewTodoController(db)

	router := gin.Default()
	routers.InitRoutes(router)
	// auth router
	authRouter := routers.NewAuthRouter(router, authController)
	authRouter.AuthRoutes()
	// todo router
	todoRouter := routers.NewTodoRouter(router, todoController)
	todoRouter.TodoRoutes()

	return &InitConfig{
		Router: router,
	}
}
