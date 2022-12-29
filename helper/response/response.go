package response

import "github.com/labstack/echo/v4"

type (
	Response interface {
		BaseResponse(ctx echo.Context, code int, status string, data interface{}, err error) error
	}
)

const (
	SUCCESS = "success"
	FAILED  = "failed"
)
