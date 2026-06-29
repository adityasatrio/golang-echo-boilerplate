package validator

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"myapp/internal/helper/response"
)

func SetupGlobalHttpUnhandleErrors(e *echo.Echo) {
	e.HTTPErrorHandler = GlobalUnHandleErrors()
	log.Infof("initialized GlobalUnHandleErrors : success")
}

func GlobalUnHandleErrors() func(err error, ctx echo.Context) {
	return func(err error, ctx echo.Context) {
		errHttpCode, errBusinessCode, msg, errCode := MapperErrorCode(err)
		_ = response.Base(ctx, errHttpCode, errBusinessCode, msg, nil, errCode)
	}
}
