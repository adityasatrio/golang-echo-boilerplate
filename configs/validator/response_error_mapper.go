package validator

import (
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"myapp/exceptions"
	"net/http"
)

// MapperErrorCode all mapper goes here
func MapperErrorCode(err error) (errHttpCode int, errBusinessCode int, errMessage string, errUnexpected error) {

	var bizErr *exceptions.BusinessLogicError
	if errors.As(err, &bizErr) {
		errorLogic := exceptions.BusinessLogicReason(bizErr.ErrorCode)

		return errorLogic.HttpCode, errorLogic.ErrCode, errorLogic.Message, nil
	}

	errorMessage := err.Error()
	var castedObject validator.ValidationErrors
	if errors.As(err, &castedObject) {
		for _, e := range castedObject {
			switch e.Tag() {
			case "required":
				errorMessage = fmt.Sprintf("%s is required", e.Field())
			case "email":
				errorMessage = fmt.Sprintf("%s is not valid email", e.Field())
			case "gte":
				errorMessage = fmt.Sprintf("%s value must be greater than %s", e.Field(), e.Param())
			case "lte":
				errorMessage = fmt.Sprintf("%s value must be lower than %s", e.Field(), e.Param())
			case "password":
				errorMessage = fmt.Sprintf("%s %s", e.Field(), e.Value())
			}
			break //nolint:staticcheck // only first validation error is used
		}

		return http.StatusBadRequest, http.StatusBadRequest, errorMessage, nil
	}

	var httpErr *echo.HTTPError
	if errors.As(err, &httpErr) {
		message, ok := httpErr.Message.(string)
		if !ok {
			message = http.StatusText(httpErr.Code)
		}
		return httpErr.Code, httpErr.Code, message, nil
	}

	return http.StatusInternalServerError, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), errors.New(errorMessage)

}
