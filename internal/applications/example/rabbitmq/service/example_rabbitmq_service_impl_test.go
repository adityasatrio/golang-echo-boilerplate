package service

import (
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"myapp/ent"
	"myapp/internal/applications/system_parameter/dto"
	mock_db "myapp/mocks/applications/system_parameter/repository/db"
	"testing"
)

func TestGetMessage_DataDoesNotExist(t *testing.T) {
	mockRepo := new(mock_db.SystemParameterRepository)
	svc := NewExampleRabbitMQService(mockRepo)

	ctx := context.Background()
	request := &dto.SystemParameterCreateRequest{
		Key:   "testKey",
		Value: "testValue",
	}

	// Mocking repository behavior for GetByKey to return nil (data doesn't exist)
	mockRepo.On("GetByKey", ctx, request.Key).
		Return(nil, nil)

	// Mocking repository behavior for Create to return a new entry
	mockRepo.On("Create", ctx, mock.Anything).
		Return(&ent.SystemParameter{Key: request.Key, Value: request.Value}, nil)

	result, err := svc.GetMessage(ctx, request)

	// Assert that the result matches the expected output
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, request.Key, result.Key)
	assert.Equal(t, request.Value, result.Value)

	// Verify that the expected repository methods were called
	mockRepo.AssertExpectations(t)
}

func TestGetMessage_CreateError(t *testing.T) {
	mockRepo := new(mock_db.SystemParameterRepository)
	svc := NewExampleRabbitMQService(mockRepo)

	ctx := context.Background()
	request := &dto.SystemParameterCreateRequest{
		Key:   "testKey",
		Value: "testValue",
	}

	expectedErr := errors.New("example error")

	// Mocking repository behavior for GetByKey to return nil (data doesn't exist)
	mockRepo.On("GetByKey", ctx, request.Key).Return(nil, nil)

	// Mocking repository behavior for Create to return an error
	mockRepo.On("Create", ctx, mock.Anything).Return(nil, expectedErr)

	result, err := svc.GetMessage(ctx, request)

	// Assert that the error returned is of type BusinessLogicError and has the expected error code
	assert.Error(t, err)
	assert.Nil(t, result)

	// Verify that the expected repository methods were called
	mockRepo.AssertExpectations(t)
}

func TestGetMessage_DataExists(t *testing.T) {
	mockRepo := new(mock_db.SystemParameterRepository)
	svc := NewExampleRabbitMQService(mockRepo)

	ctx := context.Background()
	request := &dto.SystemParameterCreateRequest{
		Key:   "testKey",
		Value: "testValue",
	}

	// Mocking repository behavior for GetByKey to return existing data
	mockRepo.On("GetByKey", ctx, request.Key).
		Return(&ent.SystemParameter{Key: request.Key, Value: request.Value}, nil)

	result, err := svc.GetMessage(ctx, request)

	// Assert that the error is of type BusinessLogicError and the error code is DataAlreadyExist
	assert.Error(t, err)

	// Assert that the result is nil
	assert.Nil(t, result)

	// Verify that the expected repository methods were called
	mockRepo.AssertExpectations(t)
}
