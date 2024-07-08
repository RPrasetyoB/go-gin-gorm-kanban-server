package routers

import (
	"go-kanban/controllers"
	middleware "go-kanban/middlewares"

	"github.com/gin-gonic/gin"
)

type TodoRouterImpl struct {
	router         *gin.Engine
	todoController controllers.TodoController
}

func NewTodoRouter(router *gin.Engine, todoController controllers.TodoController) *TodoRouterImpl {
	return &TodoRouterImpl{
		router:         router,
		todoController: todoController,
	}
}

func (r *TodoRouterImpl) TodoRoutes() {
	todoRouter := r.router.Group("/api/v1/todo", middleware.Auth())
	{
		todoRouter.POST("", r.todoController.CreateTodo)
		todoRouter.GET("", r.todoController.GetUserTodos)
		todoRouter.GET("/:id", r.todoController.GetTodoById)
		todoRouter.PUT("/:id", r.todoController.UpdateTodo)
		todoRouter.DELETE("/:id", r.todoController.DeleteTodo)
	}
}
