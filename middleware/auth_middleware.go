package middleware

import (
	"net/http"

	"github.com/labstack/echo/v4"
	authService "myapp/internal/applications/auth/service"
)

// SessionCookieName is the httpOnly cookie that carries the app-issued JWT.
const SessionCookieName = "session_token"

// AdminRoleID matches the "Admin" role seeded by
// migrations/20260630120005_seed_roles_and_admin_user.go.
const AdminRoleID = uint64(1)

const (
	ContextKeyUserID = "auth_user_id"
	ContextKeyRoleID = "auth_role_id"
	ContextKeyEmail  = "auth_email"
)

func readSessionClaims(c echo.Context, authSvc authService.AuthService) (*authService.Claims, error) {
	cookie, err := c.Cookie(SessionCookieName)
	if err != nil || cookie.Value == "" {
		return nil, echo.NewHTTPError(http.StatusUnauthorized, "missing session")
	}

	claims, err := authSvc.VerifyJWT(cookie.Value)
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusUnauthorized, "invalid or expired session")
	}

	return claims, nil
}

// RequireAuth protects JSON API routes: it validates the session cookie and
// responds with 401 JSON on failure.
func RequireAuth(authSvc authService.AuthService) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			claims, err := readSessionClaims(c, authSvc)
			if err != nil {
				return echo.NewHTTPError(http.StatusUnauthorized, "unauthorized")
			}

			c.Set(ContextKeyUserID, claims.UserID)
			c.Set(ContextKeyRoleID, claims.RoleID)
			c.Set(ContextKeyEmail, claims.Email)

			return next(c)
		}
	}
}

// RequireAuthWeb protects HTML page/partial routes: it validates the session
// cookie and redirects to /login on failure.
func RequireAuthWeb(authSvc authService.AuthService) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			claims, err := readSessionClaims(c, authSvc)
			if err != nil {
				return c.Redirect(http.StatusFound, "/login")
			}

			c.Set(ContextKeyUserID, claims.UserID)
			c.Set(ContextKeyRoleID, claims.RoleID)
			c.Set(ContextKeyEmail, claims.Email)

			return next(c)
		}
	}
}

// RequireAdmin must run after RequireAuth/RequireAuthWeb. It rejects callers
// whose role claim does not match the seeded Admin role.
func RequireAdmin() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			roleID, ok := c.Get(ContextKeyRoleID).(uint64)
			if !ok || roleID != AdminRoleID {
				return echo.NewHTTPError(http.StatusForbidden, "forbidden")
			}

			return next(c)
		}
	}
}
