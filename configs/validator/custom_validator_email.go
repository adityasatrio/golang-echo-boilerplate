package validator

import (
	"github.com/go-playground/validator/v10"
	"net/mail"
)

func IsValidEmail(fl validator.FieldLevel) bool {
	email := fl.Field().String()
	_, err := mail.ParseAddress(email)
	return err == nil
}
