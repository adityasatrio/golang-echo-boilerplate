package response

import "github.com/labstack/echo/v4"

type Response struct {
	Code   int         `json:"code"`
	Status string      `json:"status"`
	Error  string      `json:"error"`
	Data   interface{} `json:"data"`
}

func Return(c echo.Context, code int, status string, err error, data interface{}) error {
	response := Response{
		Code:   code,
		Status: status,
		Error:  "",
		Data:   data,
	}

	if err != nil {
		response.Error = err.Error()
	}

	return c.JSON(code, data)
}
