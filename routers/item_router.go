package routers

import (
	"go-kanban/controllers"
	middleware "go-kanban/middlewares"

	"github.com/gin-gonic/gin"
)

type ItemRouterImpl struct {
	router         *gin.Engine
	itemController controllers.ItemController
}

func NewItemRouter(router *gin.Engine, itemController controllers.ItemController) *ItemRouterImpl {
	return &ItemRouterImpl{
		router:         router,
		itemController: itemController,
	}
}

func (r *ItemRouterImpl) ItemRoutes() {
	itemRouter := r.router.Group("/api/v1/item", middleware.Auth())
	{
		itemRouter.POST("", r.itemController.CreateItem)
		itemRouter.GET("/todo/:todoId", r.itemController.GetItemList)
		itemRouter.GET("/:id", r.itemController.GetItemById)
		itemRouter.PUT("/:id", r.itemController.UpdateItem)
		itemRouter.DELETE("/:id", r.itemController.DeleteItem)
	}
}
