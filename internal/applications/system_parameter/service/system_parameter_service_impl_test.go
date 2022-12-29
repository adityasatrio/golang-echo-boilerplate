package service

import (
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHello(t *testing.T) {
	// SetupRouteHandler
	e := echo.New()
	request := httptest.NewRequest(http.MethodGet, "/", nil)
	ctx := e.NewContext(request, nil)

	//test use case
	useCase := NewSystemParameterService(nil)
	actualResult, err := useCase.Hello(ctx.Request().Context())

	//assert
	expected := "hello from case impl"
	if assert.NoError(t, err) {
		assert.Equal(t, expected, actualResult)
	}
}
