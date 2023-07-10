package response

import (
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"myapp/internal/helper"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestBase(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Call the Base function with all possible variations
	err := errors.New("some error")
	data := struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}{ID: 1, Name: "John Doe"}

	// Test with nil data and nil error
	err = Base(c, http.StatusOK, 200, "OK", nil, nil)
	if err != nil {
		t.Fatalf("Expected error to be nil, but got: %v", err)
	}

	// Test with data and nil error
	err = Base(c, http.StatusOK, 200, "OK", data, nil)
	if err != nil {
		t.Fatalf("Expected error to be nil, but got: %v", err)
	}

	// Test with data and error
	err = Base(c, http.StatusInternalServerError, 500, "Internal Server Error", data, err)
	if err != nil {
		t.Fatalf("Expected error to be nil, but got: %v", err)
	}

	// Test with nil data and error
	err = Base(c, http.StatusNotFound, 404, "Not Found", nil, err)
	if err != nil {
		t.Fatalf("Expected error to be nil, but got: %v", err)
	}
}

func TestCreated(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Call the Created function with data
	data := struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}{
		ID:   1,
		Name: "John Doe",
	}

	err := Created(c, data)
	if err != nil {
		t.Fatalf("Expected error to be nil, but got: %v", err)
	}

	// Check response code
	if rec.Code != http.StatusCreated {
		t.Errorf("Expected status code %d, but got %d", http.StatusCreated, rec.Code)
	}

	codeJson, _ := helper.GetFieldBytes(rec.Body.Bytes(), "code")
	assert.Equal(t, float64(201), codeJson)

	messageJson, _ := helper.GetFieldBytes(rec.Body.Bytes(), "message")
	assert.Equal(t, "Created", messageJson)

	dataIdJson, _ := helper.GetFieldBytes(rec.Body.Bytes(), "data.id")
	assert.Equal(t, float64(1), dataIdJson)

	dataNameJson, _ := helper.GetFieldBytes(rec.Body.Bytes(), "data.name")
	assert.Equal(t, "John Doe", dataNameJson)

}

func TestSuccess(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Call the Success function with data
	data := struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}{ID: 1, Name: "John Doe"}
	err := Success(c, data)
	if err != nil {
		t.Fatalf("Expected error to be nil, but got: %v", err)
	}

	// Check response code
	if rec.Code != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, rec.Code)
	}

	codeJson, _ := helper.GetFieldBytes(rec.Body.Bytes(), "code")
	assert.Equal(t, float64(200), codeJson)

	messageJson, _ := helper.GetFieldBytes(rec.Body.Bytes(), "message")
	assert.Equal(t, "OK", messageJson)

	dataIdJson, _ := helper.GetFieldBytes(rec.Body.Bytes(), "data.id")
	assert.Equal(t, float64(1), dataIdJson)

	dataNameJson, _ := helper.GetFieldBytes(rec.Body.Bytes(), "data.name")
	assert.Equal(t, "John Doe", dataNameJson)
}

func TestError(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Call the Error function with an error
	err := errors.New("some error occurred")
	httpCode := http.StatusInternalServerError
	err = Error(c, httpCode, err)
	if err != nil {
		t.Fatalf("Expected error to be nil, but got: %v", err)
	}

	// Check response code
	if rec.Code != httpCode {
		t.Errorf("Expected status code %d, but got %d", httpCode, rec.Code)
	}

	codeJson, _ := helper.GetFieldBytes(rec.Body.Bytes(), "code")
	assert.Equal(t, float64(500), codeJson)

	messageJson, _ := helper.GetFieldBytes(rec.Body.Bytes(), "message")
	assert.Equal(t, "Internal Server Error", messageJson)

	errorJson, _ := helper.GetFieldBytes(rec.Body.Bytes(), "error")
	assert.Equal(t, "some error occurred", errorJson)

}
