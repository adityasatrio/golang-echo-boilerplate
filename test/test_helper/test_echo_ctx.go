package test_helper

import (
	"context"
	"github.com/labstack/echo/v4"
	"net/http"
	"net/http/httptest"
)

func NewServiceCtx() context.Context {
	e := echo.New()
	rec := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodGet, "/hello", nil)
	c := e.NewContext(request, rec)
	ctx := c.Request().Context()

	return ctx
}
