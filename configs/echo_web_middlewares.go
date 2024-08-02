//references :
//	https://dasarpemrogramangolang.novalagung.com/C-advanced-middleware-and-logging.html
//	https://echo.labstack.com/middleware

package configs

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	customMiddleware "myapp/middleware"
)

func SetupMiddlewares(e *echo.Echo) {

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.Gzip())
	//e.Use(middleware.CSRF())

	e.Use(middleware.Secure())

	e.Use(customMiddleware.CorsConfig())
	e.Use(customMiddleware.GlobalRequestTimeout())

	//disable default built-in middleware RequestID, because we using the custom that we own
	//e.Use(middleware.RequestID())

	//use custom middleware for expose request id in log
	e.Use(customMiddleware.LoggerTraceRequestID())

	log.Info("initialized middleware : success")

}
