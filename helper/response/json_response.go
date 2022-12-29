package response

import (
	"github.com/labstack/echo/v4"
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
		ServerTime int64       `json:"serverTime"`
	}
)

func NewBaseResponse() Response {
	return &JsonResponse{}
}

func (responseImpl *JsonResponse) BaseResponse(ctx echo.Context, code int, status string, data interface{}, err error) error {
	bodyResponse := body{
		Code:       code,
		Status:     status,
		ServerTime: time.Now().UnixMilli(),
	}

	if data != nil {
		bodyResponse.Data = data
	}

	if err != nil {
		bodyResponse.Error = err.Error()
	}

	//return responseImpl.ctx.JSON(bodyResponse.Code, bodyResponse)
	return ctx.JSON(bodyResponse.Code, bodyResponse)
}
