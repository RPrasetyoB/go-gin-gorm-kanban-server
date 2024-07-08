package services

import (
	"errors"
	"go-kanban/helper"
	"go-kanban/models"
	repositories "go-kanban/repositories/todo"
	"net/http"

	"gorm.io/gorm"
)

type TodoService interface {
	CreateNewTodoService(userId int, todo *models.Todos) (*models.Todos, error)
	GetTodoByIdService(todoId int, userId int) (*models.Todos, error)
	UpdateTodoService(todoId int, userId int, updateTodo *models.Todos) (*models.Todos, error)
	DeleteTodoService(todoId int, userId int) (*models.Todos, error)
	FindUserTodoService(userId int) ([]models.Todos, error)
}

type TodoServiceImpl struct {
	repo repositories.TodoRepository
}

func NewTodoService(repo repositories.TodoRepository) TodoService {
	return &TodoServiceImpl{
		repo: repo,
	}
}

// CreateNewTodo implements TodoService.
func (t *TodoServiceImpl) CreateNewTodoService(userId int, todo *models.Todos) (*models.Todos, error) {
	newTodo, err := t.repo.CreateTodo(userId, todo)
	if err != nil {
		return nil, &helper.CustomError{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	}
	return newTodo, nil
}

// FindUserTodo implements TodoService.
func (t *TodoServiceImpl) FindUserTodoService(userId int) ([]models.Todos, error) {
	todolist, err := t.repo.FindAll(userId)
	if err != nil {
		return nil, &helper.CustomError{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	}
	return todolist, nil
}

// getTodoById implements TodoService.
func (t *TodoServiceImpl) GetTodoByIdService(todoId int, userId int) (*models.Todos, error) {
	getTodo, err := t.repo.FindTodoById(todoId)
	if err != nil {
		var statusCode int
		if errors.Is(err, gorm.ErrRecordNotFound) {
			statusCode = http.StatusNotFound
		} else {
			statusCode = http.StatusInternalServerError
		}

		return nil, &helper.CustomError{
			Code:    statusCode,
			Message: err.Error(),
		}
	}
	if userId != getTodo.User_id {
		return nil, &helper.CustomError{
			Code:    http.StatusForbidden,
			Message: "Forbidden to retrieve other user's todo",
		}
	}
	todo, err := t.repo.FindTodoById(todoId)
	if err != nil {
		var statusCode int
		if errors.Is(err, gorm.ErrRecordNotFound) {
			statusCode = http.StatusNotFound
		} else {
			statusCode = http.StatusInternalServerError
		}

		return nil, &helper.CustomError{
			Code:    statusCode,
			Message: err.Error(),
		}
	}

	return todo, nil
}

// UpdateTodo implements TodoService.
func (t *TodoServiceImpl) UpdateTodoService(todoId int, userId int, updateTodo *models.Todos) (*models.Todos, error) {
	getTodo, err := t.repo.FindTodoById(todoId)
	if err != nil {
		var statusCode int
		if errors.Is(err, gorm.ErrRecordNotFound) {
			statusCode = http.StatusNotFound
		} else {
			statusCode = http.StatusInternalServerError
		}

		return nil, &helper.CustomError{
			Code:    statusCode,
			Message: err.Error(),
		}
	}
	if userId != getTodo.User_id {
		return nil, &helper.CustomError{
			Code:    http.StatusForbidden,
			Message: "Forbidden to update other user's todo",
		}
	}
	todo, err := t.repo.Update(todoId, updateTodo)
	if err != nil {
		var statusCode int
		if errors.Is(err, gorm.ErrRecordNotFound) {
			statusCode = http.StatusNotFound
		} else {
			statusCode = http.StatusInternalServerError
		}

		return nil, &helper.CustomError{
			Code:    statusCode,
			Message: err.Error(),
		}
	}

	return todo, nil
}

// Delete implements TodoService.
func (t *TodoServiceImpl) DeleteTodoService(todoId int, userId int) (*models.Todos, error) {
	getTodo, err := t.repo.FindTodoById(todoId)
	if err != nil {
		var statusCode int
		var messageErr string
		if errors.Is(err, gorm.ErrRecordNotFound) {
			statusCode = http.StatusNotFound
			messageErr = "todo not found"
		} else {
			statusCode = http.StatusInternalServerError
			messageErr = err.Error()
		}

		return nil, &helper.CustomError{
			Code:    statusCode,
			Message: messageErr,
		}
	}
	if userId != getTodo.User_id {
		return nil, &helper.CustomError{
			Code:    http.StatusForbidden,
			Message: "Forbidden to delete other user's todo",
		}
	}
	todo, err := t.repo.Delete(todoId)
	if err != nil {
		return nil, &helper.CustomError{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	return todo, nil
}
