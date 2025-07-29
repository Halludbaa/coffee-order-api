package apperrors

import (
	"errors"
	"fmt"

	"github.com/go-playground/validator/v10"
)

func getFieldMessage(fieldError validator.FieldError) string {
	switch fieldError.Tag() {
	case "required":
		return "This field is required"
	case "email":
		return "Invalid email"
	case "min":
		return fmt.Sprintf("Should have at least %v characters", fieldError.Param())
	case "max":
		return fmt.Sprintf("Should be less than or equal to %v characters", fieldError.Param())
	}
	return fieldError.Error()
}


func GetValidateMessage(err error) []APIError {
	var validationError validator.ValidationErrors
	if errors.As(err, &validationError) {
		out := make([]APIError, len(validationError))
		for index, fieldError := range validationError {
			out[index] = APIError{fieldError.Field(), getFieldMessage(fieldError)}
		}
		return out
	}
	return []APIError{}
}