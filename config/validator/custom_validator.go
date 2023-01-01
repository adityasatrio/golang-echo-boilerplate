package validator

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"log"
)

type CustomValidator struct {
	Validator *validator.Validate
}

func NewValidator() *validator.Validate {
	return validator.New()
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.Validator.Struct(i)
}

func SetupValidator(e *echo.Echo) {
	e.Validator = &CustomValidator{Validator: NewValidator()}
	log.Default().Println("initialized CustomValidator : success")
}
