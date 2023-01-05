// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package system_parameter

import (
	"github.com/eko/gocache/lib/v4/cache"
	"github.com/google/wire"
	"myapp/ent"
	"myapp/internal/applications/system_parameter/repository/db"
	"myapp/internal/applications/system_parameter/service"
)

// Injectors from system_parameter_injector.go:

func InitializedSystemParameterService(dbClient *ent.Client, cacheManager *cache.ChainCache[any]) *service.SystemParameterServiceImpl {
	systemParameterRepositoryImpl := db.NewSystemParameterRepository(dbClient)
	systemParameterServiceImpl := service.NewSystemParameterService(systemParameterRepositoryImpl, cacheManager)
	return systemParameterServiceImpl
}

// system_parameter_injector.go:

var providerSetSystemParameter = wire.NewSet(db.NewSystemParameterRepository, service.NewSystemParameterService, wire.Bind(new(db.SystemParameterRepository), new(*db.SystemParameterRepositoryImpl)), wire.Bind(new(service.SystemParameterService), new(*service.SystemParameterServiceImpl)))
