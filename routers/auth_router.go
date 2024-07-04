package routers

import (
	"go-kanban/controllers"

	"github.com/gin-gonic/gin"
)

type AuthRouterImpl struct {
	router         *gin.Engine
	authController controllers.AuthController
}

func NewAuthRouter(router *gin.Engine, authController controllers.AuthController) *AuthRouterImpl {
	return &AuthRouterImpl{
		router:         router,
		authController: authController,
	}
}

func (r *AuthRouterImpl) RegisterRoutes() {
	authRouter := r.router.Group("/api/v1/auth")
	{
		authRouter.POST("/register", r.authController.CreateUser)
		authRouter.POST("/login", r.authController.LoginUser)
	}
}
