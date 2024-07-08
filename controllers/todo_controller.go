package controllers

import (
	"go-kanban/helper"
	"go-kanban/http/request"
	"go-kanban/http/response"
	middleware "go-kanban/middlewares"
	"go-kanban/models"
	repositories "go-kanban/repositories/todo"
	"go-kanban/services"
	"go-kanban/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type TodoController interface {
	CreateTodo(ctx *gin.Context)
	GetUserTodos(ctx *gin.Context)
	GetTodoById(ctx *gin.Context)
	UpdateTodo(ctx *gin.Context)
	DeleteTodo(ctx *gin.Context)
}

type TodoControllerImpl struct {
	todoService services.TodoService
	validator   *validator.Validate
}

func NewTodoController(db *gorm.DB) TodoController {
	todoRepo := repositories.NewTodoImpl(db)
	todoService := services.NewTodoService(todoRepo)
	validator := validator.New()

	return &TodoControllerImpl{
		todoService: todoService,
		validator:   validator,
	}
}

// CreateTodo implements TodoController.
func (t *TodoControllerImpl) CreateTodo(ctx *gin.Context) {
	createRequest := request.TodoRequest{}
	errType := ctx.ShouldBindJSON(&createRequest)
	if errType != nil {
		response := response.ErrorResponse{
			Success: false,
			Code:    http.StatusBadRequest,
			Message: "Payload type invalid",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	err := t.validator.Struct(createRequest)
	if err != nil {
		errorMsg := utils.GetErrorMessage(err)
		response := response.ErrorResponse{
			Success: false,
			Code:    http.StatusBadRequest,
			Message: errorMsg,
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	user := middleware.GetUser(ctx)
	userId := user["ID"].(int)

	todo := &models.Todos{
		Title:       createRequest.Title,
		Description: createRequest.Description,
	}

	createdTodo, err := t.todoService.CreateNewTodoService(userId, todo)
	if err != nil {
		if CustomError, ok := err.(*helper.CustomError); ok {
			response := response.ErrorResponse{
				Success: false,
				Code:    CustomError.Code,
				Message: CustomError.Message,
			}
			ctx.JSON(CustomError.Code, response)
			return
		}

		response := response.ErrorResponse{
			Success: false,
			Code:    http.StatusInternalServerError,
			Message: "Failed to register user",
		}
		ctx.JSON(http.StatusInternalServerError, response)
		return
	}
	webResponse := response.SuccessResponse{
		Success: true,
		Code:    http.StatusOK,
		Message: "New todo added successfully",
		Data:    createdTodo,
	}
	ctx.JSON(http.StatusOK, webResponse)
}

// GetUserTodos implements TodoController.
func (t *TodoControllerImpl) GetUserTodos(ctx *gin.Context) {
	user := middleware.GetUser(ctx)
	userId := user["ID"].(int)

	todolist, err := t.todoService.FindUserTodoService(userId)
	if err != nil {
		if CustomError, ok := err.(*helper.CustomError); ok {
			response := response.ErrorResponse{
				Success: false,
				Code:    CustomError.Code,
				Message: CustomError.Message,
			}
			ctx.JSON(CustomError.Code, response)
			return
		}

		response := response.ErrorResponse{
			Success: false,
			Code:    http.StatusInternalServerError,
			Message: "Failed to create new todo",
		}
		ctx.JSON(http.StatusInternalServerError, response)
		return
	}
	if len(todolist) == 0 {
		response := response.ErrorResponse{
			Success: false,
			Code:    http.StatusNotFound,
			Message: "User hasn't created any todos yet.",
		}
		ctx.JSON(http.StatusNotFound, response)
		return
	}
	webResponse := response.SuccessResponse{
		Success: true,
		Code:    http.StatusOK,
		Message: "Todo list retrieved successfully",
		Data:    todolist,
	}
	ctx.JSON(http.StatusOK, webResponse)
}

// GetTodoById implements TodoController.
func (t *TodoControllerImpl) GetTodoById(ctx *gin.Context) {
	idParam := ctx.Param("id")
	todoId, _ := strconv.Atoi(idParam)
	user := middleware.GetUser(ctx)
	userId := user["ID"].(int)

	todo, err := t.todoService.GetTodoByIdService(todoId, userId)
	if err != nil {
		if CustomError, ok := err.(*helper.CustomError); ok {
			response := response.ErrorResponse{
				Success: false,
				Code:    CustomError.Code,
				Message: CustomError.Message,
			}
			ctx.JSON(CustomError.Code, response)
			return
		}

		response := response.ErrorResponse{
			Success: false,
			Code:    http.StatusInternalServerError,
			Message: "Failed to get todo",
		}
		ctx.JSON(http.StatusInternalServerError, response)
		return
	}
	webResponse := response.SuccessResponse{
		Success: true,
		Code:    http.StatusOK,
		Message: "Todo retrieved successfully",
		Data:    todo,
	}
	ctx.JSON(http.StatusOK, webResponse)
}

// UpdateTodo implements TodoController.
func (t *TodoControllerImpl) UpdateTodo(ctx *gin.Context) {
	createRequest := request.TodoRequest{}
	errType := ctx.ShouldBindJSON(&createRequest)
	if errType != nil {
		response := response.ErrorResponse{
			Success: false,
			Code:    http.StatusBadRequest,
			Message: "Payload type invalid",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	err := t.validator.Struct(createRequest)
	if err != nil {
		errorMsg := utils.GetErrorMessage(err)
		response := response.ErrorResponse{
			Success: false,
			Code:    http.StatusBadRequest,
			Message: errorMsg,
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	idParam := ctx.Param("id")
	todoId, _ := strconv.Atoi(idParam)
	user := middleware.GetUser(ctx)
	userId := user["ID"].(int)

	todo := &models.Todos{
		Title:       createRequest.Title,
		Description: createRequest.Description,
	}

	updatedTodo, err := t.todoService.UpdateTodoService(todoId, userId, todo)
	if err != nil {
		if CustomError, ok := err.(*helper.CustomError); ok {
			response := response.ErrorResponse{
				Success: false,
				Code:    CustomError.Code,
				Message: CustomError.Message,
			}
			ctx.JSON(CustomError.Code, response)
			return
		}

		response := response.ErrorResponse{
			Success: false,
			Code:    http.StatusInternalServerError,
			Message: "Failed to update todo",
		}
		ctx.JSON(http.StatusInternalServerError, response)
		return
	}
	webResponse := response.SuccessResponse{
		Success: true,
		Code:    http.StatusOK,
		Message: "Todo updated successfully",
		Data:    updatedTodo,
	}
	ctx.JSON(http.StatusOK, webResponse)
}

// DeleteTodo implements TodoController.
func (t *TodoControllerImpl) DeleteTodo(ctx *gin.Context) {
	idParam := ctx.Param("id")
	todoId, _ := strconv.Atoi(idParam)
	user := middleware.GetUser(ctx)
	userId := user["ID"].(int)

	_, err := t.todoService.DeleteTodoService(todoId, userId)
	if err != nil {
		if CustomError, ok := err.(*helper.CustomError); ok {
			response := response.ErrorResponse{
				Success: false,
				Code:    CustomError.Code,
				Message: CustomError.Message,
			}
			ctx.JSON(CustomError.Code, response)
			return
		}

		response := response.ErrorResponse{
			Success: false,
			Code:    http.StatusInternalServerError,
			Message: "Failed to delete todo",
		}
		ctx.JSON(http.StatusInternalServerError, response)
		return
	}
	webResponse := response.SuccessResponse{
		Success: true,
		Code:    http.StatusOK,
		Message: "Todo deleted successfully",
	}
	ctx.JSON(http.StatusOK, webResponse)
}
