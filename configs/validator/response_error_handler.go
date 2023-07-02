package validator

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"myapp/internal/apputils"
)

func SetupGlobalHttpUnhandleErrors(e *echo.Echo) {
	e.HTTPErrorHandler = GlobalUnHandleErrors()
	log.Infof("initialized GlobalUnHandleErrors : success")
}

func GlobalUnHandleErrors() func(err error, ctx echo.Context) {
	return func(err error, ctx echo.Context) {
		_, ok := err.(*echo.HTTPError)
		if !ok {

			errHttpCode, errBusinessCode, msg, errCode := MapperErrorCode(err)
			_ = apputils.Base(ctx, errHttpCode, errBusinessCode, msg, nil, errCode)
			return
		}

		errHttpCode, errBusinessCode, msg, errCode := MapperErrorCode(err)
		_ = apputils.Base(ctx, errHttpCode, errBusinessCode, msg, nil, errCode)
		return
	}
}
