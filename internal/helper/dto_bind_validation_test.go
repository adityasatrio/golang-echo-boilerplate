package helper

import (
	validator2 "github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"myapp/configs/validator"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type TestStruct struct {
	Name string `json:"name" validate:"required"`
}

func TestBindAndValidate(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(echo.POST, "/", strings.NewReader(`{"name":"John Doe"}`))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	validator.SetupValidator(e)

	data := new(TestStruct)

	err := BindAndValidate(c, data)
	assert.NoError(t, err)

	assert.Equal(t, "John Doe", data.Name)
}

func TestBindAndValidate_BindError(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(echo.POST, "/", strings.NewReader(`{"name":"John Doe"`)) // Malformed JSON
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	validator.SetupValidator(e)

	data := new(TestStruct)

	err := BindAndValidate(c, data)
	assert.Error(t, err)

	// Check if the error is of type echo.HTTPError
	if httpErr, ok := err.(*echo.HTTPError); ok {
		// Assert specific properties of the HTTP error
		assert.Equal(t, http.StatusBadRequest, httpErr.Code)
		assert.Equal(t, "unexpected EOF", httpErr.Message)
		assert.Equal(t, "unexpected EOF", httpErr.Internal.Error())
	} else {
		t.Errorf("Expected error of type *echo.HTTPError, but got %T", err)
	}
}

func TestBindAndValidate_ValidationError(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(echo.POST, "/", strings.NewReader(`{"name":""}`)) // Empty name field
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	validator.SetupValidator(e)

	data := new(TestStruct)

	err := BindAndValidate(c, data)
	assert.Error(t, err)

	// Check if the error is of type echo.HTTPError
	if vErr, ok := err.(validator2.ValidationErrors); ok {
		// Assert specific properties of the HTTP error
		assert.Equal(t, "required", vErr[0].Tag())
		assert.Equal(t, "Name", vErr[0].Field())
	} else {
		t.Errorf("Expected error of type *echo.HTTPError, but got %T", err)
	}
}
