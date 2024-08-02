package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/viper"
	"net/http"
)

func CorsConfig() echo.MiddlewareFunc {
	corsAllowedHost := viper.GetString("application.cors.allowedHost")
	return middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{corsAllowedHost},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
		AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodPatch, http.MethodDelete},
	})
}
