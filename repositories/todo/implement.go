package repositories

import (
	"go-kanban/models"

	"gorm.io/gorm"
)

// TodoImpl struct implements the TodoRepository interface.
type TodoImpl struct {
	Db *gorm.DB
}

func NewTodoImpl(Db *gorm.DB) TodoRepository {
	return &TodoImpl{Db: Db}
}

// CreateTodo implements TodoRepository.
func (t *TodoImpl) CreateTodo(userId int, todo *models.Todos) (*models.Todos, error) {
	todo.User_id = userId
	result := t.Db.Create(todo)
	if result.Error != nil {
		return nil, result.Error
	}
	return todo, nil
}

// FindAll implements TodoRepository.
func (t *TodoImpl) FindAll(userId int) ([]models.Todos, error) {
	var todos []models.Todos
	result := t.Db.Where("user_id = ?", userId).Find(&todos)
	if result.Error != nil {
		return nil, result.Error
	}
	return todos, nil
}

// FindTodoById implements TodoRepository.
func (t *TodoImpl) FindTodoById(todoId int) (*models.Todos, error) {
	var todo models.Todos
	db := t.Db.Session(&gorm.Session{PrepareStmt: true}).Begin()
	result := db.Where("id = ?", todoId).First(&todo)
	if result.Error != nil {
		return nil, result.Error
	}
	return &todo, nil
}

// Update implements TodoRepository.
func (t *TodoImpl) Update(todoId int, updateTodo *models.Todos) (*models.Todos, error) {
	var todo models.Todos
	session := t.Db.Session(&gorm.Session{})
	err := session.First(&todo, todoId).Error
	if err != nil {
		return nil, err
	}

	todo.Title = updateTodo.Title
	todo.Description = updateTodo.Description
	result := session.Save(&todo)

	if result.Error != nil {
		return nil, result.Error
	}
	return &todo, nil
}

// Delete implements TodoRepository.
func (t *TodoImpl) Delete(todoId int) (*models.Todos, error) {
	var todo models.Todos
	result := t.Db.Unscoped().Delete(&todo, todoId)
	if result.Error != nil {
		return nil, result.Error
	}
	return &todo, nil
}
