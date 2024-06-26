// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package health

import (
	"github.com/go-redis/redis/v8"
	"github.com/google/wire"
	"myapp/ent"
	"myapp/internal/applications/cache"
	"myapp/internal/applications/health/repository"
	"myapp/internal/applications/health/service"
)

// Injectors from health_injector.go:

func InitializeHealthService(dbClient *ent.Client, cacheClient *redis.Client) *service.HealthServiceImpl {
	healthRepositoryImpl := repository.NewHealthRepository(dbClient)
	cachingServiceImpl := cache.NewCachingService(cacheClient)
	healthServiceImpl := service.NewHealthService(healthRepositoryImpl, cachingServiceImpl)
	return healthServiceImpl
}

// health_injector.go:

var providerSetHealth = wire.NewSet(repository.NewHealthRepository, service.NewHealthService, cache.NewCachingService, wire.Bind(new(repository.HealthRepository), new(*repository.HealthRepositoryImpl)), wire.Bind(new(service.HealthService), new(*service.HealthServiceImpl)), wire.Bind(new(cache.CachingService), new(*cache.CachingServiceImpl)))
