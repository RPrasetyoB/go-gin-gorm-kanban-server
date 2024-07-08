package repositories

import (
	"errors"
	"go-kanban/models"

	"gorm.io/gorm"
)

type AuthImpl struct {
	Db *gorm.DB
}

func NewAuthImpl(Db *gorm.DB) AuthRepository {
	return &AuthImpl{Db: Db}
}

// CreateUser implements AuthRepository.
func (a *AuthImpl) CreateUser(user *models.Users) (*models.Users, error) {
	result := a.Db.Create(user)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

// FindByUsername implements AuthRepository.
func (a *AuthImpl) FindByUsername(username string) (*models.Users, error) {
	var user models.Users
	db := a.Db.Session(&gorm.Session{PrepareStmt: true}).Begin()

	result := db.Where("username = ?", username).First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			db.Rollback()
			return nil, errors.New("user not found")
		}
		db.Rollback()
		return nil, result.Error
	}

	if err := db.Commit().Error; err != nil {
		return nil, err
	}

	return &user, nil
}

// Delete implements AuthRepository.
func (a *AuthImpl) Delete(userId int) (*models.Users, error) {
	panic("unimplemented")
}

// FindAll implements AuthRepository.
func (a *AuthImpl) FindAll() []models.Users {
	panic("unimplemented")
}

// FindById implements AuthRepository.
func (a *AuthImpl) FindById(userId int) (*models.Users, error) {
	panic("unimplemented")
}

// Update implements AuthRepository.
func (a *AuthImpl) Update(user models.Users) (*models.Users, error) {
	panic("unimplemented")
}
