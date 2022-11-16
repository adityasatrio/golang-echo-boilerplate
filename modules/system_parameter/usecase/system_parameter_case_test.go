package usecase

import (
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHello(t *testing.T) {
	// Setup
	e := echo.New()
	request := httptest.NewRequest(http.MethodGet, "/", nil)
	ctx := e.NewContext(request, nil)

	//test use case
	useCase := NewUseCase()
	actualResult, err := useCase.Hello(ctx.Request().Context())

	//assert
	expected := "hello from case impl"
	if assert.NoError(t, err) {
		assert.Equal(t, expected, actualResult)
	}
}
