package controllers

import (
	"go-kanban/helper"
	"go-kanban/http/request"
	"go-kanban/http/response"
	"go-kanban/models"
	repositories "go-kanban/repositories/auth"
	"go-kanban/services"
	"go-kanban/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type AuthController interface {
	CreateUser(ctx *gin.Context)
	LoginUser(ctx *gin.Context)
}

type AuthControllerImpl struct {
	authService services.AuthService
	validator   *validator.Validate
}

func NewAuthController(db *gorm.DB) AuthController {
	authRepo := repositories.NewAuthImpl(db)
	authService := services.NewAuthService(authRepo)
	validator := validator.New()

	return &AuthControllerImpl{
		authService: authService,
		validator:   validator,
	}
}

func (r *AuthControllerImpl) CreateUser(ctx *gin.Context) {
	createRequest := request.CreateRequest{}
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

	err := r.validator.Struct(createRequest)
	if err != nil {
		errorMsg := utils.GetErrorMessage(err)
		response := response.ErrorResponse{
			Success: false,
			Code:    http.StatusBadRequest,
			Message: errorMsg,
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	user := &models.Users{
		Username: createRequest.Username,
		Password: createRequest.Password,
	}

	createdUser, err := r.authService.Register(user)
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
			Message: "Failed to register user",
		}
		ctx.JSON(http.StatusInternalServerError, response)
		return
	}

	webResponse := response.SuccessResponse{
		Success: true,
		Code:    http.StatusOK,
		Message: "User registered successfully",
		Data:    createdUser,
	}
	ctx.JSON(http.StatusOK, webResponse)
}

func (l *AuthControllerImpl) LoginUser(ctx *gin.Context) {
	loginUser := request.LoginRequest{}
	errType := ctx.ShouldBindJSON(&loginUser)
	if errType != nil {
		response := response.ErrorResponse{
			Success: false,
			Code:    http.StatusBadRequest,
			Message: "Payload type invalid",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	err := l.validator.Struct(loginUser)
	if err != nil {
		errorMsg := utils.GetErrorMessage(err)
		response := response.ErrorResponse{
			Success: false,
			Code:    http.StatusBadRequest,
			Message: errorMsg,
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	token, err := l.authService.Login(loginUser.Username, loginUser.Password)
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
			Message: "Failed to login",
		}
		ctx.JSON(http.StatusInternalServerError, response)
		return
	}

	webResponse := response.SuccessResponse{
		Success: true,
		Code:    http.StatusOK,
		Message: "User logged in successfully",
		Data: map[string]interface{}{
			"token": token,
		},
	}
	ctx.JSON(http.StatusOK, webResponse)
}
