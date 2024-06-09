package utils

import (
	"errors"

	"github.com/go-playground/validator"
)

func ValidateRequestStruct(err error) error {
	if err == nil {
		return nil
	}

	if _, ok := err.(validator.ValidationErrors); !ok {
		return err
	}

	for _, err := range err.(validator.ValidationErrors) {
		switch err.Tag() {
		case "required":
			return errors.New("Missing required parameter: " + err.Field())
		default:
			return errors.New("Invalid parameter: " + err.Field())
		}
	}

	return nil
}
