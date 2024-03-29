//go:build wireinject
// +build wireinject

package health

import (
	"github.com/go-redis/redis/v8"
	"github.com/google/wire"
	"myapp/ent"
	"myapp/internal/applications/A_templates_directory/repository/db"
	"myapp/internal/applications/A_templates_directory/service"
	"myapp/internal/component/cache"
)

var providerSetHealth = wire.NewSet(
	db.NewTemplateRepository,
	service.NewTemplateService,
	cache.NewCache,
	wire.Bind(new(db.TemplateRepository), new(*db.TemplateRepositoryImpl)),
	wire.Bind(new(service.TemplateService), new(*service.TemplateServiceImpl)),
	wire.Bind(new(cache.Cache), new(*cache.CacheImpl)),
)

func InitializeHealthService(dbClient *ent.Client, cacheClient *redis.Client) *service.TemplateServiceImpl {
	wire.Build(providerSetHealth)
	return nil
}
