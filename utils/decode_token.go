package utils

import (
	"fmt"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
)

var jwtSecret string

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Printf("Error loading .env file: %v\n", err)
	}

	jwtSecret = os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		fmt.Println("JWT_SECRET not found in .env file")
	}
}

func GetToken(c *gin.Context) (jwt.MapClaims, error) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		return nil, nil
	}

	tokenString := strings.Split(authHeader, " ")[1]

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtSecret), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}

func LoggedUser(decodedToken jwt.MapClaims) map[string]interface{} {
	if decodedToken == nil {
		return nil
	}
	var userID interface{}
	if v, ok := decodedToken["user_id"].(float64); ok {
		userID = int(v)
	} else {
		userID = decodedToken["user_id"]
	}

	return map[string]interface{}{
		"ID": userID,
	}
}
