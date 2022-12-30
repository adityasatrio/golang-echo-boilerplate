package response

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

type (
	JsonResponse struct {
		ctx echo.Context
	}

	body struct {
		Code       int         `json:"code"`
		Status     string      `json:"status"`
		Data       interface{} `json:"data"`
		Error      string      `json:"error"`
		ServerTime string      `json:"serverTime"`
	}
)

func NewBaseResponse() Response {
	return &JsonResponse{}
}

func (responseImpl *JsonResponse) Base(ctx echo.Context, code int, status string, data interface{}, err error) error {

	date := time.Now().Format(time.RFC1123)
	bodyResponse := body{
		Code:       code,
		Status:     status,
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
	ctx.Response().Header().Add(HeaderDate, date)
	//TODO : should we implement etag header ?

	return ctx.JSON(bodyResponse.Code, bodyResponse)
}

func (responseImpl *JsonResponse) Created(ctx echo.Context, data interface{}) error {
	return responseImpl.Base(ctx, http.StatusCreated, http.StatusText(http.StatusCreated), data, nil)
}

func (responseImpl *JsonResponse) Success(ctx echo.Context, data interface{}) error {
	return responseImpl.Base(ctx, http.StatusOK, http.StatusText(http.StatusOK), data, nil)
}

func (responseImpl *JsonResponse) Error(ctx echo.Context, code int, err error) error {
	return responseImpl.Base(ctx, code, http.StatusText(code), nil, err)
}
