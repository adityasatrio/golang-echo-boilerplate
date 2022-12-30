package response

import (
	"errors"
	"github.com/labstack/echo/v4"
	"myapp/exceptions"
	"net/http"
)

func ServiceErrorHandler(ctx echo.Context, result any, err error) error {
	if err != nil {
		if errors.Is(err, exceptions.TargetBusinessLogicError) {
			return Error(ctx, http.StatusUnprocessableEntity, err)

		} else if errors.Is(err, exceptions.TargetDataNotFoundError) {
			return Base(ctx, http.StatusNotFound, result, err)

		} else {
			return Error(ctx, http.StatusInternalServerError, err)
		}
	}

	//default return handler
	return Success(ctx, result)
}
