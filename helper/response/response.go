package response

import (
	"github.com/labstack/echo/v4"
)

type (
	Response interface {
		Base(ctx echo.Context, code int, status string, data interface{}, err error) error
		Created(ctx echo.Context, data interface{}) error
		Success(ctx echo.Context, data interface{}) error
		Error(ctx echo.Context, code int, err error) error
	}
)

const (
	// StatusSuccess status for success on response
	StatusSuccess = "success"

	// StatusFailed status for failed on response
	StatusFailed = "failed"

	HeaderDate = "Date"
)
