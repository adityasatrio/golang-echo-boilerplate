package test

import (
	"github.com/labstack/echo/v4"
	"myapp/configs/validator"
	"testing"
)

func InitEchoTest(*testing.T) *echo.Echo {
	e := echo.New()
	validator.SetupValidator(e)
	validator.SetupGlobalHttpUnhandleErrors(e)

	return e
}
