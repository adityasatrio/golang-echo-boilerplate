package service

import (
	"context"
	"github.com/labstack/gommon/log"
	"myapp/internal/applications/cache"
	"myapp/internal/applications/health/repository"
)

type HealthServiceImpl struct {
	repository repository.HealthRepository
	cache      cache.CachingService
}

func NewHealthService(repository repository.HealthRepository, cache cache.CachingService) *HealthServiceImpl {
	return &HealthServiceImpl{
		repository: repository,
		cache:      cache,
	}
}

func (s *HealthServiceImpl) Health(ctx context.Context, message string) (map[string]string, error) {
	messageService := message + "hello from service layer "
	result, errRepo := s.repository.Health(ctx, messageService)

	errCache := s.cache.Ping(ctx)
	if errCache != nil {
		result["cache_status"] = "DOWN"
		result["cache_name"] = "redis"
	} else {
		result["cache_status"] = "UP"
		result["cache_name"] = "redis"
	}

	//database will have first err priority
	if errRepo != nil {
		return result, errRepo
	}

	if errCache != nil {
		return result, errCache
	}

	return result, nil
}

func (s *HealthServiceImpl) HealthDatabase(ctx context.Context, message string) (map[string]string, error) {
	messageService := message + "hello from service layer "
	result, errRepo := s.repository.Health(ctx, messageService)
	//database will have first err priority
	if errRepo != nil {
		return result, errRepo
	}

	return result, nil
}

func (s *HealthServiceImpl) HealthCache(ctx context.Context, message string) (map[string]string, error) {

	healthCheck := map[string]string{}
	messageService := message + "hello from service layer "
	healthCheck["final_msg"] = messageService

	if ctx != nil {
		log.Info("ctx debug", ctx)
		healthCheck["ctx_status"] = "UP"
		healthCheck["ctx_name"] = "echo"
	}

	errCache := s.cache.Ping(ctx)
	if errCache != nil {
		healthCheck["cache_status"] = "DOWN"
		healthCheck["cache_name"] = "redis"
	} else {
		healthCheck["cache_status"] = "UP"
		healthCheck["cache_name"] = "redis"
	}

	if errCache != nil {
		return healthCheck, errCache
	}

	return healthCheck, nil
}
