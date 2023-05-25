package validator

import (
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"log"
	"myapp/exceptions"
	"myapp/internal/apputils"
	"net/http"
)

func SetupGlobalHttpUnhandleErrors(e *echo.Echo) {
	e.HTTPErrorHandler = GlobalUnHandleErrors()
	log.Default().Println("initialized GlobalUnHandleErrors : success")
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

// MapperErrorCode all mapper goes here
func MapperErrorCode(err error) (errHttpCode int, errBusinessCode int, errMessage string, errUnexpected error) {

	if errors.Is(err, exceptions.TargetBusinessLogicError) {
		errorCode := err.(*exceptions.BusinessLogicError).ErrorCode
		errorLogic := exceptions.BusinessLogicReason(errorCode)

		return errorLogic.HttpCode, errorLogic.ErrBusinessCode, errorLogic.Message, nil
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
			case "password":
				errorMessage = fmt.Sprintf("%s %s", err.Field(), err.Value())
			}
			break
		}

		return http.StatusBadRequest, http.StatusBadRequest, errorMessage, nil
	}

	return http.StatusInternalServerError, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), errors.New(errorMessage)

}
