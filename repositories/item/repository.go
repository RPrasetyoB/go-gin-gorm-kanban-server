package repositories

import "go-kanban/models"

type ItemRepository interface {
	CreateItem(item *models.Items) (*models.Items, error)
	FindAll(todoId int) ([]models.Items, error)
	GetItemById(itemId int) (*models.Items, error)
	UpdateItem(itemId int, item *models.Items) (*models.Items, error)
	Delete(itemId int) (*models.Items, error)
}
