package services

import (
	"errors"
	"go-kanban/helper"
	"go-kanban/models"
	repositories "go-kanban/repositories/item"
	"net/http"

	"gorm.io/gorm"
)

type ItemService interface {
	CreateNewItemService(item *models.Items) (*models.Items, error)
	GetItemByIdService(itemId int) (*models.Items, error)
	GetAllItemService(todoId int) ([]models.Items, error)
	UpdateItemService(itemId int, updateItem *models.Items) (*models.Items, error)
	DeleteItemService(itemId int) (*models.Items, error)
}

type ItemServiceImpl struct {
	repo repositories.ItemRepository
}

func NewItemService(repo repositories.ItemRepository) ItemService {
	return &ItemServiceImpl{
		repo: repo,
	}
}

// CreateNewItemService implements ItemService.
func (i *ItemServiceImpl) CreateNewItemService(item *models.Items) (*models.Items, error) {
	newItem, err := i.repo.CreateItem(item)
	if err != nil {
		return nil, &helper.CustomError{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	}
	return newItem, nil
}

// GetAllItemService implements ItemService.
func (i *ItemServiceImpl) GetAllItemService(todoId int) ([]models.Items, error) {
	items, err := i.repo.FindAll(todoId)
	if err != nil {
		return nil, &helper.CustomError{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	}
	return items, nil
}

// GetItemByIdService implements ItemService.
func (i *ItemServiceImpl) GetItemByIdService(itemId int) (*models.Items, error) {
	getItem, err := i.repo.GetItemById(itemId)
	if err != nil {
		var statusCode int
		var messageErr string
		if errors.Is(err, gorm.ErrRecordNotFound) {
			statusCode = http.StatusNotFound
			messageErr = "item not found"
		} else {
			statusCode = http.StatusInternalServerError
			messageErr = err.Error()
		}

		return nil, &helper.CustomError{
			Code:    statusCode,
			Message: messageErr,
		}
	}

	return getItem, nil
}

// UpdateItemService implements ItemService.
func (i *ItemServiceImpl) UpdateItemService(itemId int, updateItem *models.Items) (*models.Items, error) {
	item, err := i.repo.UpdateItem(itemId, updateItem)
	if err != nil {
		var statusCode int
		var messageErr string
		if errors.Is(err, gorm.ErrRecordNotFound) {
			statusCode = http.StatusNotFound
			messageErr = "item not found"
		} else {
			statusCode = http.StatusInternalServerError
			messageErr = err.Error()
		}

		return nil, &helper.CustomError{
			Code:    statusCode,
			Message: messageErr,
		}
	}

	return item, nil
}

// DeleteItemService implements ItemService.
func (i *ItemServiceImpl) DeleteItemService(itemId int) (*models.Items, error) {
	_, err := i.repo.GetItemById(itemId)
	if err != nil {
		var statusCode int
		var messageErr string
		if errors.Is(err, gorm.ErrRecordNotFound) {
			statusCode = http.StatusNotFound
			messageErr = "item not found"
		} else {
			statusCode = http.StatusInternalServerError
			messageErr = err.Error()
		}

		return nil, &helper.CustomError{
			Code:    statusCode,
			Message: messageErr,
		}
	}
	item, err := i.repo.Delete(itemId)
	if err != nil {
		return nil, &helper.CustomError{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	}
	return item, nil
}
