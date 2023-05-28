package validator

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type CustomValidator struct {
	Validator *validator.Validate
}

var (
	validate    *validator.Validate
	passwordTag = "password"
	emailTag    = "email"
)

func NewValidator() *validator.Validate {

	validate = validator.New()

	//register password:
	err := validate.RegisterValidation(passwordTag, IsValidPassword)
	if err != nil {
		return nil
	}

	//register email:
	err = validate.RegisterValidation(emailTag, IsValidEmail)
	if err != nil {
		return nil
	}

	return validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.Validator.Struct(i)
}

func SetupValidator(e *echo.Echo) {
	e.Validator = &CustomValidator{Validator: NewValidator()}
	log.Info("initialized CustomValidator : success")
}
