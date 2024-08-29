package validator_format

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

type ValidationError struct {
	Field   string `json:"fieldRequired"`
	Message string `json:"message"`
}

func FormatValidator(err error) []ValidationError {
	var ValidationErr []ValidationError

	for _, err := range err.(validator.ValidationErrors) {
		ValidationError := ValidationError{
			Field:   err.Field(),
			Message: fmt.Sprintf("%s is %s !", err.Field(), err.Tag()),
		}
		ValidationErr = append(ValidationErr, ValidationError)
	}
	return ValidationErr
}
