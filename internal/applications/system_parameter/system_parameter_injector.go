//go:build wireinject
// +build wireinject

package system_parameter

import (
	"github.com/eko/gocache/lib/v4/cache"
	"github.com/google/wire"
	"myapp/ent"
	"myapp/internal/applications/system_parameter/repository/db"
	"myapp/internal/applications/system_parameter/service"
)

var providerSetSystemParameter = wire.NewSet(
	db.NewSystemParameterRepository,
	service.NewSystemParameterService,
	wire.Bind(new(db.SystemParameterRepository), new(*db.SystemParameterRepositoryImpl)),
	wire.Bind(new(service.SystemParameterService), new(*service.SystemParameterServiceImpl)),
)

func InitializedSystemParameterService(dbClient *ent.Client, cacheManager *cache.ChainCache[any]) *service.SystemParameterServiceImpl {
	wire.Build(providerSetSystemParameter)
	return nil
}
