package validator

import (
	"github.com/go-playground/validator/v10"
	"regexp"
)

func IsValidPassword(fl validator.FieldLevel) bool {
	password := fl.Field().String()

	//Password Should be of 8 characters long:
	if len(password) < 8 {
		fl.Field().SetString("Should be of 8 characters long")
		return false
	}

	//Password should contain at least one lower case character:
	oneCharLower, errOneCharLower := regexp.MatchString("([a-z])+", password)
	if errOneCharLower != nil || !oneCharLower {
		fl.Field().SetString("Should contain at least one lower case character")
		return false
	}

	//Password should contain at least one upper case character:
	oneCharUpper, errOneCharUpper := regexp.MatchString("([A-Z])+", password)
	if errOneCharUpper != nil || !oneCharUpper {
		fl.Field().SetString("Should contain at least one upper case character")
		return false
	}

	//Password should contain at least one digit:
	oneDigit, errOneDigit := regexp.MatchString("([0-9])+", password)
	if errOneDigit != nil || !oneDigit {
		fl.Field().SetString("Should contain at least one digit")
		return false
	}

	//Password should contain at least one special character
	sChar, errSChar := regexp.MatchString("([!@#$%^&*.?-])+", password)
	if errSChar != nil || !sChar {
		fl.Field().SetString("Should contain at least one special character")
		return false
	}

	return true
}
