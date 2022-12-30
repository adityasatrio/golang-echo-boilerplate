package response

import (
	"errors"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

type body struct {
	Code       int         `json:"code"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
	Error      string      `json:"error"`
	ServerTime string      `json:"serverTime"`
}

func Base(ctx echo.Context, code int, data interface{}, err error) error {

	date := time.Now().Format(time.RFC1123)
	bodyResponse := body{
		Code:       code,
		Message:    http.StatusText(code),
		ServerTime: date,
	}

	if data != nil {
		bodyResponse.Data = data
	}

	if err != nil {
		bodyResponse.Error = err.Error()
	}

	//TODO : using context as injection ?
	//return responseImpl.ctx.JSON(bodyResponse.Code, bodyResponse)

	//added header for standard
	//https://developer.mozilla.org/en-US/docs/Glossary/Response_header
	ctx.Response().Header().Add("date", date)
	//TODO : should we implement etag header ?

	return ctx.JSON(bodyResponse.Code, bodyResponse)
}

func Created(ctx echo.Context, data interface{}) error {
	if data == nil {
		panic(errors.New("success response : data on body is mandatory"))
	}

	return Base(ctx, http.StatusCreated, data, nil)
}

func Success(ctx echo.Context, data interface{}) error {
	if data == nil {
		panic(errors.New("success response : data on body is mandatory"))
	}

	return Base(ctx, http.StatusOK, data, nil)
}

func Error(ctx echo.Context, code int, err error) error {
	return Base(ctx, code, nil, err)
}
