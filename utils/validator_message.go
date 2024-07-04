package utils

import (
	"reflect"

	"github.com/go-playground/validator/v10"
)

func GetErrorMessage(err error) string {
	validationErrors, ok := err.(validator.ValidationErrors)
	if !ok {
		return "Validation error: " + err.Error()
	}

	errorMsg := ""
	for _, fieldErr := range validationErrors {
		errorMsg += GetMessage(fieldErr)
	}

	return errorMsg
}

func GetMessage(fieldErr validator.FieldError) string {
	switch fieldErr.Tag() {
	case "required":
		return fieldErr.Field() + " is required"
	case "min":
		return fieldErr.Field() + " must be at least " + fieldErr.Param() + " characters long"
	case "max":
		return fieldErr.Field() + " must be at most " + fieldErr.Param() + " characters long"
	case "numeric":
		return fieldErr.Field() + " must be a valid number"
	case "email":
		return fieldErr.Field() + " must be a valid email address"
	case "string":
		if fieldErr.Kind() == reflect.String {
			return fieldErr.Field() + " must be a string"
		}
	case "integer":
		if fieldErr.Kind() == reflect.Int {
			return fieldErr.Field() + " must be an integer"
		}
	// Add more cases for other custom tags as needed
	default:
		// Handle other errors
		return fieldErr.Error()
	}

	// Default fallback
	return fieldErr.Error()
}
