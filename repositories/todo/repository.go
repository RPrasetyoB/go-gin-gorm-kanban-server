package repositories

import "go-kanban/models"

type TodoRepository interface {
	CreateTodo(userId int, todo *models.Todos) (*models.Todos, error)
	FindTodoById(todoId int) (*models.Todos, error)
	Update(todoId int, updatedTodo *models.Todos) (*models.Todos, error)
	Delete(todoId int) (*models.Todos, error)
	FindAll(userId int) ([]models.Todos, error)
}
