package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/viper"
	"net/http"
	"time"
)

func InitMiddlewares(e *echo.Echo) {
	e.Use(CorsConfig())
	e.Use(GlobalRequestTimeout())

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.Gzip())
	//e.Use(middleware.CSRF())
	e.Use(middleware.Logger())
	e.Use(middleware.RequestID())
	e.Use(middleware.Secure())

}

func GlobalRequestTimeout() echo.MiddlewareFunc {
	return middleware.TimeoutWithConfig(middleware.TimeoutConfig{
		Timeout: 5 * time.Minute,
	})
}

func CorsConfig() echo.MiddlewareFunc {
	corsAllowedHost := viper.GetString("application.cors.allowedHost")
	return middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{corsAllowedHost},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
		AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodPatch, http.MethodDelete},
	})
}
