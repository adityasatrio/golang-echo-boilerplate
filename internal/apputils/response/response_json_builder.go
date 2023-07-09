package response

import (
	"errors"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

type Body struct {
	ErrorCode  int         `json:"code"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
	Error      string      `json:"error"`
	ServerTime string      `json:"serverTime"`
}

func Base(ctx echo.Context, httpCode int, errorCode int, message string, data interface{}, err error) error {

	date := time.Now().Format(time.RFC1123)
	bodyResponse := Body{
		ErrorCode:  errorCode,
		Message:    message,
		ServerTime: date,
	}

	if data != nil {
		bodyResponse.Data = data
	}

	if err != nil {
		bodyResponse.Error = err.Error()
	}

	//added header for standard response
	//https://developer.mozilla.org/en-US/docs/Glossary/Response_header
	ctx.Response().Header().Add("date", date)

	return ctx.JSON(httpCode, bodyResponse)
}

func Created(ctx echo.Context, data interface{}) error {
	if data == nil {
		panic(errors.New("success response : data on body is mandatory"))
	}

	return Base(ctx, http.StatusCreated, http.StatusCreated, http.StatusText(http.StatusCreated), data, nil)
}

func Success(ctx echo.Context, data interface{}) error {
	if data == nil {
		panic(errors.New("success response : data on body is mandatory"))
	}

	return Base(ctx, http.StatusOK, http.StatusOK, http.StatusText(http.StatusOK), data, nil)
}

//goland:noinspection GoUnusedExportedFunction
func Error(ctx echo.Context, httpCode int, err error) error {
	return Base(ctx, httpCode, httpCode, http.StatusText(httpCode), nil, err)
}
