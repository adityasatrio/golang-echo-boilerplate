package test_helper

import (
	"github.com/labstack/echo/v4"
	"myapp/configs/validator"
	"testing"
)

func InitEchoTest(*testing.T) *echo.Echo {
	e := echo.New()
	validator.SetupValidator(e)

	return e
}
