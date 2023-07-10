package service

import (
	"context"
	"myapp/internal/applications/health/repository"
)

type HealthServiceImpl struct {
	repository repository.HealthRepository
}

func NewHealthService(repository repository.HealthRepository) *HealthServiceImpl {
	return &HealthServiceImpl{
		repository: repository,
	}
}

func (s *HealthServiceImpl) Health(ctx context.Context, message string, queryFlag string) (map[string]string, error) {
	messageService := message + "hello from service layer "
	result, err := s.repository.Health(ctx, messageService, queryFlag)
	return result, err
}
