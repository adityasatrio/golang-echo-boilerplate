package service

import (
	"context"
	"github.com/stretchr/testify/assert"
	mock_repository "myapp/mocks/health/repository"
	"testing"
)

func TestNewHelloWorldsService(t *testing.T) {
	// CreateTx a mock service object.
	mockRepository := &mock_repository.HealthRepository{}

	// Call the function being tested.
	service := NewHealthService(mockRepository)

	// Check that the service field of the controller has the expected value.
	if service.repository != mockRepository {
		t.Errorf("Expected service to be %v, but got %v", mockRepository, service.repository)
	}
}

func TestHello(t *testing.T) {

	ctx := context.Background()

	healthCheckResult := map[string]string{}
	healthCheckResult["final_msg"] = "hello from controller layer hello from service layer hello from repository layer "

	mockRepository := &mock_repository.HealthRepository{}
	mockRepository.On("Health", ctx, "hello from controller layer hello from service layer ", "default").
		Return(healthCheckResult, nil)

	service := NewHealthService(mockRepository)
	result, err := service.Health(ctx, "hello from controller layer ", "default")

	assert.Equal(t, nil, err)
	assert.Equal(t, "hello from controller layer hello from service layer hello from repository layer ", result["final_msg"])
}
