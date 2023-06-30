//go:build wireinject
// +build wireinject

package system_parameter

import (
	"github.com/go-redis/redis/v8"
	"github.com/google/wire"
	"myapp/ent"
	"myapp/internal/applications/cache"
	"myapp/internal/applications/system_parameter/repository/db"
	"myapp/internal/applications/system_parameter/service"
)

var providerSetSystemParameter = wire.NewSet(
	db.NewSystemParameterRepository,
	service.NewSystemParameterService,
	cache.NewCachingServiceImpl,
	wire.Bind(new(db.SystemParameterRepository), new(*db.SystemParameterRepositoryImpl)),
	wire.Bind(new(cache.CachingService), new(*cache.CachingServiceImpl)),
	wire.Bind(new(service.SystemParameterService), new(*service.SystemParameterServiceImpl)),
)

func InitializedSystemParameterService(dbClient *ent.Client, redisClient *redis.Client) *service.SystemParameterServiceImpl {
	wire.Build(providerSetSystemParameter)
	return nil
}
