package helper

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

func NewErrValidation(err error) *ErrBadRequest {
	message := FormatValidationErrors(err)
	return &ErrBadRequest{Message: message}
}

func FormatValidationErrors(err error) string {
	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		errorMessages := make([]string, 0)
		for _, e := range validationErrors {
			errorMessages = append(errorMessages, formatValidationError(e))
		}
		return strings.Join(errorMessages, ", ")
	}
	return "Invalid request payload"
}

func formatValidationError(err validator.FieldError) string {
	switch err.Tag() {
	case "required":
		return fmt.Sprintf("%s is required", err.Field())
	case "min":
		return fmt.Sprintf("%s must be at least %s characters long", err.Field(), err.Param())
	case "max":
		return fmt.Sprintf("%s must be at most %s characters long", err.Field(), err.Param())
	case "email":
		return fmt.Sprintf("%s must be a valid email address", err.Field())
	default:
		return fmt.Sprintf("%s is invalid", err.Field())
	}
}
