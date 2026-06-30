package middleware

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golang.org/x/time/rate"
	"myapp/configs/credential"
)

func rateLimiterConfig(ratePerSecond int) middleware.RateLimiterConfig {
	if ratePerSecond <= 0 {
		ratePerSecond = 20
	}

	return middleware.RateLimiterConfig{
		Skipper: middleware.DefaultSkipper,
		Store: middleware.NewRateLimiterMemoryStoreWithConfig(
			middleware.RateLimiterMemoryStoreConfig{
				Rate:  rate.Limit(ratePerSecond),
				Burst: ratePerSecond,
			},
		),
		IdentifierExtractor: func(c echo.Context) (string, error) {
			return c.RealIP(), nil
		},
		ErrorHandler: func(c echo.Context, err error) error {
			return c.JSON(http.StatusInternalServerError, echo.Map{"message": "rate limiter error"})
		},
		DenyHandler: func(c echo.Context, identifier string, err error) error {
			return c.JSON(http.StatusTooManyRequests, echo.Map{"message": "too many requests"})
		},
	}
}

// GlobalRateLimiter applies a generous, in-memory, per-IP rate limit to all
// routes via Echo's built-in rate limiter (no Redis dependency).
func GlobalRateLimiter() echo.MiddlewareFunc {
	return middleware.RateLimiterWithConfig(rateLimiterConfig(credential.GetInt("ratelimit.global.rps")))
}

// LoginRateLimiter applies a stricter, in-memory, per-IP rate limit to login
// and Auth0 callback routes to blunt credential-stuffing attempts.
func LoginRateLimiter() echo.MiddlewareFunc {
	return middleware.RateLimiterWithConfig(rateLimiterConfig(credential.GetInt("ratelimit.login.rps")))
}
