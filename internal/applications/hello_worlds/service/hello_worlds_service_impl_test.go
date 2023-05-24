package service

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"myapp/exceptions"
	mock_repository "myapp/mocks/hello_worlds/repository"
	"myapp/test/test_helper"
	"testing"
)

func TestNewHelloWorldsService(t *testing.T) {
	// CreateTx a mock service object.
	mockRepository := &mock_repository.HelloWorldsRepository{}

	// Call the function being tested.
	service := NewHelloWorldsService(mockRepository)

	// Check that the service field of the controller has the expected value.
	if service.repository != mockRepository {
		t.Errorf("Expected service to be %v, but got %v", mockRepository, service.repository)
	}
}

func TestHello(t *testing.T) {

	ctx := test_helper.NewServiceCtx()

	mockRepository := &mock_repository.HelloWorldsRepository{}
	mockRepository.On("Hello", ctx, "hello from controller -> hello from s-impl layer ->", "").
		Return("hello from controller -> hello from s-impl layer -> hello from repository-impl layer ", nil)

	service := NewHelloWorldsService(mockRepository)
	result, err := service.Hello(ctx, "hello from controller ->", "")

	assert.Equal(t, nil, err)
	assert.Equal(t, "hello from controller -> hello from s-impl layer -> hello from repository-impl layer ", result)
}

func TestHelloErr(t *testing.T) {

	ctx := test_helper.NewServiceCtx()

	mockRepository := &mock_repository.HelloWorldsRepository{}
	messageService := "hello from controller -> hello from s-impl layer ->"
	mockErr := exceptions.NewBusinessLogicError(exceptions.EBL10007, errors.New(messageService))

	service := NewHelloWorldsService(mockRepository)
	result, err := service.Hello(ctx, "hello from controller ->", "service")

	assert.Equal(t, mockErr, err)
	assert.Equal(t, "", result)
}
