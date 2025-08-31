//go:build wireinject
// +build wireinject

package post

import (
	"github.com/go-redis/redis/v8"
	"github.com/google/wire"
	"myapp/ent"
	"myapp/internal/applications/post/repository/db"
	"myapp/internal/applications/post/service"
	"myapp/internal/component/cache"
)

var providerSetHealth = wire.NewSet(
	db.NewPostRepository,
	service.NewPostService,
	cache.NewCache,
	wire.Bind(new(db.PostRepository), new(*db.PostRepositoryImpl)),
	wire.Bind(new(service.PostService), new(*service.PostServiceImpl)),
	wire.Bind(new(cache.Cache), new(*cache.CacheImpl)),
)

func InitializeHealthService(dbClient *ent.Client, cacheClient *redis.Client) *service.PostServiceImpl {
	wire.Build(providerSetHealth)
	return nil
}
