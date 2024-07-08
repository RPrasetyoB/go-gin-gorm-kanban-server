package services

import (
	"fmt"
	"go-kanban/helper"
	"go-kanban/models"
	repositories "go-kanban/repositories/auth"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	Register(user *models.Users) (*models.Users, error)
	Login(username, password string) (string, error)
}

type AuthServiceImpl struct {
	repo      repositories.AuthRepository
	jwtSecret []byte
}

func NewAuthService(repo repositories.AuthRepository) AuthService {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Printf("Error loading .env file: %v\n", err)
	}

	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		fmt.Println("JWT_SECRET not found in .env file")
	}

	return &AuthServiceImpl{
		repo:      repo,
		jwtSecret: []byte(jwtSecret),
	}
}

// Register a new user
func (s *AuthServiceImpl) Register(user *models.Users) (*models.Users, error) {
	// Check if username already exists
	existingUser, _ := s.repo.FindByUsername(user.Username)
	if existingUser != nil {
		return nil, &helper.CustomError{
			Code:    http.StatusConflict,
			Message: "Username already exists",
		}
	}
	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, &helper.CustomError{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	}
	user.Password = string(hashedPassword)

	// Create user
	newUser, err := s.repo.CreateUser(user)
	if err != nil {
		return nil, &helper.CustomError{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	return newUser, nil
}

// Login user
func (s *AuthServiceImpl) Login(username, password string) (string, error) {
	user, err := s.repo.FindByUsername(username)
	if err != nil {
		return "", &helper.CustomError{
			Code:    http.StatusBadRequest,
			Message: "invalid username or password",
		}
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", &helper.CustomError{
			Code:    http.StatusBadRequest,
			Message: "invalid username or password",
		}
	}

	tokenString, err := s.generateToken(user)
	if err != nil {
		return "", &helper.CustomError{
			Code:    http.StatusInternalServerError,
			Message: "failed to generat tooken",
		}
	}

	return tokenString, nil
}

// function to generate JWT token
func (s *AuthServiceImpl) generateToken(user *models.Users) (string, error) {
	claims := jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(time.Hour * 72).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(s.jwtSecret)
	if err != nil {
		return "", fmt.Errorf("failed to sign JWT token: %w", err)
	}
	return tokenString, nil
}
