package controller

import (
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	mockProducer "myapp/mocks/component/rabbitmq/producer"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPublishRabbitMQ_Success(t *testing.T) {
	mockProducer := new(mockProducer.Producer)

	// Set up expectations for the mock producer
	mockProducer.On("SendToDirect", mock.Anything, mock.Anything).Return(true, nil)

	// Create an instance of ExampleController with the mock producer
	exampleController := NewExampleRabbitMQController(mockProducer)

	// Create a new Echo instance
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/example/publish", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Call the controller method
	err := exampleController.PublishRabbitMQ(c)

	// Assertions
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)

	// Assert that expected methods were called on the mock
	mockProducer.AssertExpectations(t)
}

func TestPublishRabbitMQ_Error(t *testing.T) {
	mockProducer := new(mockProducer.Producer)

	// Set up expectations for the mock producer to simulate an error
	mockProducer.On("SendToDirect", mock.Anything, mock.Anything).Return(false, errors.New("producer error"))

	// Create an instance of ExampleController with the mock producer
	exampleController := NewExampleRabbitMQController(mockProducer)

	// Create a new Echo instance
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/example/publish", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Call the controller method
	err := exampleController.PublishRabbitMQ(c)

	// Assertions
	assert.Error(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)

	// Assert that expected methods were called on the mock
	mockProducer.AssertExpectations(t)
}
