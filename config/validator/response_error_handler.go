package validator

import (
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"myapp/exceptions"
	"myapp/helper/response"
	"net/http"
)

func SetupHttpErrorHandler(e *echo.Echo) {
	e.HTTPErrorHandler = NewHttpErrorHandler()
}

func NewHttpErrorHandler() func(err error, ctx echo.Context) {
	return func(err error, ctx echo.Context) {
		_, ok := err.(*echo.HTTPError)
		if !ok {
			if errors.Is(err, exceptions.TargetBusinessLogicError) {
				errorCode := err.(*exceptions.BusinessLogicError).ErrorCode
				errorMessage := exceptions.BusinessLogicReason(errorCode)

				response.Base(ctx, http.StatusUnprocessableEntity, errorCode, errorMessage, nil, err)
				return

			} else if errors.Is(err, exceptions.TargetDataNotFoundError) {
				response.Error(ctx, http.StatusNotFound, err)
				return

			} else if generalTypeError(err) {
				response.Error(ctx, http.StatusInternalServerError, err)
				return

			}

		}

		errorMessage := err.Error()
		if castedObject, ok := err.(validator.ValidationErrors); ok {
			for _, err := range castedObject {
				switch err.Tag() {
				case "required":
					errorMessage = fmt.Sprintf("%s is required", err.Field())
				case "email":
					errorMessage = fmt.Sprintf("%s is not valid email", err.Field())
				case "gte":
					errorMessage = fmt.Sprintf("%s value must be greater than %s", err.Field(), err.Param())
				case "lte":
					errorMessage = fmt.Sprintf("%s value must be lower than %s", err.Field(), err.Param())
				}
				break
			}
		}

		response.Base(ctx, http.StatusBadRequest, http.StatusBadRequest, http.StatusText(http.StatusBadRequest), nil, errors.New(errorMessage))
		return
	}
}

func generalTypeError(err error) bool {
	return errors.Is(err, exceptions.TargetDataCreateError) || errors.Is(err, exceptions.TargetDataDeleteError) ||
		errors.Is(err, exceptions.TargetDataGetError) || errors.Is(err, exceptions.TargetDataUpdateError)
}
