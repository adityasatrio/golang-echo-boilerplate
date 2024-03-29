//go:build wireinject
// +build wireinject

package system_parameter

import (
	"github.com/go-redis/redis/v8"
	"github.com/google/wire"
	"myapp/ent"
	"myapp/internal/applications/system_parameter/repository/db"
	"myapp/internal/applications/system_parameter/service"
	"myapp/internal/component/cache"
)

var providerSetSystemParameter = wire.NewSet(
	db.NewSystemParameterRepository,
	service.NewSystemParameterService,
	cache.NewCache,
	wire.Bind(new(db.SystemParameterRepository), new(*db.SystemParameterRepositoryImpl)),
	wire.Bind(new(cache.Cache), new(*cache.CacheImpl)),
	wire.Bind(new(service.SystemParameterService), new(*service.SystemParameterServiceImpl)),
)

func InitializedSystemParameterService(dbClient *ent.Client, redisClient *redis.Client) *service.SystemParameterServiceImpl {
	wire.Build(providerSetSystemParameter)
	return nil
}
