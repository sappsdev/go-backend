package utils

import (
	"errors"

	"gopkg.in/go-playground/validator.v9"
)

func NewValidator() *validator.Validate {
	validate := validator.New()
	return validate
}

func ValidatorErrors(error error) error {

	for range error.(validator.ValidationErrors) {
		return errors.New(error.Error())
	}
	return errors.New("User error.")
}
