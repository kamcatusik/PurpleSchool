package reques

import (
	"github.com/go-playground/validator"
)

func Valid[T any](emailLoad T) error {
	validate := validator.New()
	err := validate.Struct(emailLoad)
	if err != nil {
		return err

	}
	return nil
}
