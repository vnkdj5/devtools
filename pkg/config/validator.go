package config

import "github.com/go-playground/validator/v10"

type AppValidator struct {
	validator *validator.Validate
}

func (cv *AppValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		return err
	}
	return nil
}
