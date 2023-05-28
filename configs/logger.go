package configs

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func SetupLogger(e *echo.Echo) {
	e.Logger.SetLevel(log.DEBUG)
}
