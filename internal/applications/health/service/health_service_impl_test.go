package service

import (
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	mock_cache "myapp/mocks/applications/cache"
	mock_repository "myapp/mocks/applications/health/repository"

	"testing"
)

func TestNewNewHealthService(t *testing.T) {
	// CreateTx a mock service object.
	mockRepository := &mock_repository.HealthRepository{}
	mockCache := new(mock_cache.CachingService)

	// Call the function being tested.
	service := NewHealthService(mockRepository, mockCache)

	// Check that the service field of the controller has the expected value.
	if service.repository != mockRepository {
		t.Errorf("Expected service to be %v, but got %v", mockRepository, service.repository)
	}
}

func TestHealth_Up(t *testing.T) {

	ctx := context.Background()

	healthCheckResult := map[string]string{}
	healthCheckResult["final_msg"] = "hello from controller layer hello from service layer hello from repository layer "

	mockRepository := &mock_repository.HealthRepository{}
	mockRepository.On("Health", ctx, "hello from controller layer hello from service layer ").
		Return(healthCheckResult, nil)

	mockCache := new(mock_cache.CachingService)
	mockCache.On("Ping", ctx).Return(nil)

	service := NewHealthService(mockRepository, mockCache)
	result, err := service.Health(ctx, "hello from controller layer ")

	assert.Equal(t, nil, err)
	assert.Equal(t, "hello from controller layer hello from service layer hello from repository layer ", result["final_msg"])
}

func TestHealth_Down(t *testing.T) {

	ctx := context.Background()

	healthCheckResult := map[string]string{}
	healthCheckResult["final_msg"] = "hello from controller layer hello from service layer hello from repository layer "

	mockRepository := &mock_repository.HealthRepository{}
	mockRepository.On("Health", ctx, "hello from controller layer hello from service layer ").
		Return(healthCheckResult, nil)

	mockCache := new(mock_cache.CachingService)
	mockCache.On("Ping", ctx).Return(errors.New("random err"))

	service := NewHealthService(mockRepository, mockCache)
	result, err := service.Health(ctx, "hello from controller layer ")

	assert.NotNil(t, errors.New("random err"), err)
	assert.Equal(t, "hello from controller layer hello from service layer hello from repository layer ", result["final_msg"])
}
