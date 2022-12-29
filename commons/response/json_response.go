package response

import "github.com/labstack/echo/v4"

type (
	JsonResponse struct {
		ctx echo.Context
	}

	body struct {
		Code   int         `json:"code"`
		Status string      `json:"status"`
		Data   interface{} `json:"data"`
		Error  string      `json:"error"`
	}
)

func NewBaseResponse() Response {
	return &JsonResponse{}
}

func (responseImpl *JsonResponse) BaseResponse(ctx echo.Context, code int, status string, data interface{}, err error) error {
	bodyResponse := body{
		Code:   code,
		Status: status,
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
