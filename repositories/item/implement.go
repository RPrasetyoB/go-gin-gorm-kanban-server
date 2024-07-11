package repositories

import (
	"go-kanban/models"

	"gorm.io/gorm"
)

type ItemImpl struct {
	Db *gorm.DB
}

func NewItemImpl(Db *gorm.DB) ItemRepository {
	return &ItemImpl{Db: Db}
}

// CreateItem implements ItemRepository.
func (i *ItemImpl) CreateItem(item *models.Items) (*models.Items, error) {
	result := i.Db.Create(item)
	if result.Error != nil {
		return nil, result.Error
	}
	return item, nil
}

// FindAll implements ItemRepository.
func (i *ItemImpl) FindAll(todoId int) ([]models.Items, error) {
	var items []models.Items
	result := i.Db.Where("todo_id = ?", todoId).Find(&items)
	if result.Error != nil {
		return nil, result.Error
	}
	return items, nil
}

// GetItemById implements ItemRepository.
func (i *ItemImpl) GetItemById(itemId int) (*models.Items, error) {
	var item models.Items
	db := i.Db.Session(&gorm.Session{PrepareStmt: false}).Begin()
	result := db.Where("id = ?", itemId).First(&item)
	if result.Error != nil {
		return nil, result.Error
	}
	return &item, nil
}

func (i *ItemImpl) UpdateItem(itemId int, updateItem *models.Items) (*models.Items, error) {
	var item models.Items

	session := i.Db.Session(&gorm.Session{})
	err := session.First(&item, itemId).Error
	if err != nil {
		return nil, err
	}

	item.Todo_id = updateItem.Todo_id
	item.Name = updateItem.Name
	item.Progress_percentage = updateItem.Progress_percentage

	result := session.Save(&item)
	if result.Error != nil {
		return nil, result.Error
	}

	return &item, nil
}

// Delete implements ItemRepository.
func (i *ItemImpl) Delete(itemId int) (*models.Items, error) {
	var item models.Items
	result := i.Db.Unscoped().Delete(&item, itemId)
	if result.Error != nil {
		return nil, result.Error
	}
	return &item, nil
}
