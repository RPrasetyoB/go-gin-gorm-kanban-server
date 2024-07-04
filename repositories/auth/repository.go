package repositories

import "go-kanban/models"

type AuthRepository interface {
	CreateUser(user *models.Users) (*models.Users, error)
	FindByUsername(username string) (*models.Users, error)
	FindById(userId int) (*models.Users, error)
	Update(user models.Users) (*models.Users, error)
	Delete(userId int) (*models.Users, error)
	FindAll() []models.Users
}
