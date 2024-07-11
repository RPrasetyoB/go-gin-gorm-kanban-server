package controllers

import (
	"go-kanban/helper"
	"go-kanban/http/request"
	"go-kanban/http/response"
	"go-kanban/models"
	repositories "go-kanban/repositories/item"
	"go-kanban/services"
	"go-kanban/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type ItemController interface {
	CreateItem(ctx *gin.Context)
	GetItemList(ctx *gin.Context)
	GetItemById(ctx *gin.Context)
	UpdateItem(ctx *gin.Context)
	DeleteItem(ctx *gin.Context)
}

type ItemControllerImpl struct {
	itemService services.ItemService
	validator   *validator.Validate
}

func NewItemController(db *gorm.DB) ItemController {
	itemRepo := repositories.NewItemImpl(db)
	itemService := services.NewItemService(itemRepo)
	validator := validator.New()

	return &ItemControllerImpl{
		itemService: itemService,
		validator:   validator,
	}
}

// CreateItem implements ItemController.
func (i *ItemControllerImpl) CreateItem(ctx *gin.Context) {
	createRequest := request.ItemRequest{}
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

	err := i.validator.Struct(createRequest)
	if err != nil {
		errMessage := utils.GetErrorMessage(err)
		response := response.ErrorResponse{
			Success: false,
			Code:    http.StatusBadRequest,
			Message: errMessage,
		}
		ctx.JSON(response.Code, response)
		return
	}
	item := &models.Items{
		Todo_id:             createRequest.Todo_id,
		Name:                createRequest.Name,
		Progress_percentage: createRequest.Progress_percentage,
	}
	createItem, err := i.itemService.CreateNewItemService(item)
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
			Message: "Failed to create new item",
		}
		ctx.JSON(response.Code, response)
		return
	}
	webResponse := response.SuccessResponse{
		Success: true,
		Code:    http.StatusOK,
		Message: "New item added successfully",
		Data:    createItem,
	}
	ctx.JSON(webResponse.Code, webResponse)
}

// GetItemList implements ItemController.
func (i *ItemControllerImpl) GetItemList(ctx *gin.Context) {
	idParam := ctx.Param("todoId")
	todoId, _ := strconv.Atoi(idParam)
	itemList, err := i.itemService.GetAllItemService(todoId)
	if err != nil {
		if CustomError, ok := err.(*helper.CustomError); ok {
			response := response.ErrorResponse{
				Success: false,
				Code:    CustomError.Code,
				Message: CustomError.Message,
			}
			ctx.JSON(response.Code, response)
			return
		}

		response := response.ErrorResponse{
			Success: false,
			Code:    http.StatusInternalServerError,
			Message: "Failed to get item list",
		}
		ctx.JSON(response.Code, response)
		return
	}
	if len(itemList) == 0 {
		response := response.SuccessResponse{
			Success: true,
			Code:    http.StatusAccepted,
			Message: "User hasn't created any item in current todo.",
			Data:    itemList,
		}
		ctx.JSON(response.Code, response)
		return
	}
	webResponse := response.SuccessResponse{
		Success: true,
		Code:    http.StatusOK,
		Message: "Item list retrieved successfully",
		Data:    itemList,
	}
	ctx.JSON(webResponse.Code, webResponse)
}

// GetItemById implements ItemController.
func (i *ItemControllerImpl) GetItemById(ctx *gin.Context) {
	idParam := ctx.Param("id")
	itemId, _ := strconv.Atoi(idParam)

	item, err := i.itemService.GetItemByIdService(itemId)
	if err != nil {
		if CustomError, ok := err.(*helper.CustomError); ok {
			response := response.ErrorResponse{
				Success: false,
				Code:    CustomError.Code,
				Message: CustomError.Message,
			}
			ctx.JSON(response.Code, response)
			return
		}

		response := response.ErrorResponse{
			Success: false,
			Code:    http.StatusInternalServerError,
			Message: "Failed to get item",
		}
		ctx.JSON(response.Code, response)
		return
	}
	webResponse := response.SuccessResponse{
		Success: true,
		Code:    http.StatusOK,
		Message: "Item retrieved successfully",
		Data:    item,
	}
	ctx.JSON(webResponse.Code, webResponse)
}

// UpdateItem implements ItemController.
func (i *ItemControllerImpl) UpdateItem(ctx *gin.Context) {
	createRequest := request.ItemRequest{}
	errType := ctx.ShouldBindJSON(&createRequest)
	if errType != nil {
		response := response.ErrorResponse{
			Success: false,
			Code:    http.StatusBadRequest,
			Message: "Payload type invalid",
		}
		ctx.JSON(response.Code, response)
		return
	}

	err := i.validator.Struct(createRequest)
	if err != nil {
		errorMsg := utils.GetErrorMessage(err)
		response := response.ErrorResponse{
			Success: false,
			Code:    http.StatusBadRequest,
			Message: errorMsg,
		}
		ctx.JSON(response.Code, response)
		return
	}

	idParam := ctx.Param("id")
	itemId, _ := strconv.Atoi(idParam)

	item := &models.Items{
		Todo_id:             createRequest.Todo_id,
		Name:                createRequest.Name,
		Progress_percentage: createRequest.Progress_percentage,
	}

	updateItem, err := i.itemService.UpdateItemService(itemId, item)
	if err != nil {
		if CustomError, ok := err.(*helper.CustomError); ok {
			response := response.ErrorResponse{
				Success: false,
				Code:    CustomError.Code,
				Message: CustomError.Message,
			}
			ctx.JSON(response.Code, response)
			return
		}

		response := response.ErrorResponse{
			Success: false,
			Code:    http.StatusInternalServerError,
			Message: "Failed to update item",
		}
		ctx.JSON(response.Code, response)
		return
	}
	webResponse := response.SuccessResponse{
		Success: true,
		Code:    http.StatusOK,
		Message: "Item updated successfully",
		Data:    updateItem,
	}
	ctx.JSON(webResponse.Code, webResponse)
}

// DeleteItem implements ItemController.
func (i *ItemControllerImpl) DeleteItem(ctx *gin.Context) {
	idParam := ctx.Param("id")
	itemId, _ := strconv.Atoi(idParam)
	_, err := i.itemService.DeleteItemService(itemId)
	if err != nil {
		if CustomError, ok := err.(*helper.CustomError); ok {
			response := response.ErrorResponse{
				Success: false,
				Code:    CustomError.Code,
				Message: CustomError.Message,
			}
			ctx.JSON(response.Code, response)
			return
		}

		response := response.ErrorResponse{
			Success: false,
			Code:    http.StatusInternalServerError,
			Message: "Failed to delete item",
		}
		ctx.JSON(response.Code, response)
		return
	}
	webResponse := response.SuccessResponse{
		Success: true,
		Code:    http.StatusOK,
		Message: "Item deleted successfully",
	}
	ctx.JSON(webResponse.Code, webResponse)
}
