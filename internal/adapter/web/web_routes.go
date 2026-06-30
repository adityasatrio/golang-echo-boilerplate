package web

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"myapp/internal/builder"
	"myapp/middleware"
)

// SetupWebRouteHandler registers the HTMX + Alpine.js UI routes: the login
// page, Auth0 (password + Google) flows, and the authenticated User /
// System Parameter pages. It is a sibling of rest.SetupRouteHandler and is
// wired in alongside it from cmd/main.go.
func SetupWebRouteHandler(e *echo.Echo, container *builder.Container) {
	authSvc := container.BuildAuthService()
	userSvc := container.BuildUserService()
	roleSvc := container.BuildRoleService()
	sysParamSvc := container.BuildSystemParameterService()

	h := NewWebHandler(authSvc, userSvc, roleSvc, sysParamSvc)

	requireAuthWeb := middleware.RequireAuthWeb(authSvc)
	requireAdmin := middleware.RequireAdmin()
	loginRateLimiter := middleware.LoginRateLimiter()

	e.GET("/", func(c echo.Context) error {
		return c.Redirect(http.StatusFound, "/system-parameters")
	})

	e.GET("/login", h.LoginPage, loginRateLimiter)
	e.POST("/auth/login", h.PostLogin, loginRateLimiter)
	e.GET("/auth/google", h.GoogleLoginRedirect, loginRateLimiter)
	e.GET("/auth/callback", h.GoogleCallback, loginRateLimiter)
	e.POST("/auth/logout", h.Logout)

	users := e.Group("/users", requireAuthWeb, requireAdmin)
	users.GET("", h.UsersPage)
	users.GET("/new", h.UserFormNew)
	users.GET("/:id/edit", h.UserFormEdit)
	users.POST("", h.UserCreate)
	users.PUT("/:id", h.UserUpdate)
	users.DELETE("/:id", h.UserDelete)

	systemParameters := e.Group("/system-parameters", requireAuthWeb)
	systemParameters.GET("", h.SystemParametersPage)
	systemParameters.GET("/new", h.SystemParameterFormNew, requireAdmin)
	systemParameters.GET("/:id/edit", h.SystemParameterFormEdit, requireAdmin)
	systemParameters.POST("", h.SystemParameterCreate, requireAdmin)
	systemParameters.PUT("/:id", h.SystemParameterUpdate, requireAdmin)
	systemParameters.DELETE("/:id", h.SystemParameterDelete, requireAdmin)
}
