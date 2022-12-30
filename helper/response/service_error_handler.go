package response

import (
	"errors"
	"github.com/labstack/echo/v4"
	"myapp/exceptions"
	"myapp/exceptions/errcode"
	"net/http"
)

func ServiceErrorHandler(ctx echo.Context, result any, err error) error {
	if err != nil {
		if errors.Is(err, exceptions.TargetBusinessLogicError) {
			errorCode := err.(*exceptions.BusinessLogicError).ErrorCode
			errorMessage := errcode.BusinessLogicReason(errorCode)
			return Base(ctx, http.StatusUnprocessableEntity, errorCode, errorMessage, nil, err)

		} else if errors.Is(err, exceptions.TargetDataNotFoundError) {
			return Error(ctx, http.StatusNotFound, err)

		} else {
			return Error(ctx, http.StatusInternalServerError, err)
		}
	}

	//default return handler
	return Success(ctx, result)
}
