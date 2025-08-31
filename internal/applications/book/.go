//go:build wireinject
// +build wireinject

package book

import (
	"github.com/go-redis/redis/v8"
	"github.com/google/wire"
	"myapp/ent"
	"myapp/internal/applications/book/repository/db"
	"myapp/internal/applications/book/service"
	"myapp/internal/component/cache"
)

var providerSetHealth = wire.NewSet(
	db.NewBookRepository,
	service.NewBookService,
	cache.NewCache,
	wire.Bind(new(db.BookRepository), new(*db.BookRepositoryImpl)),
	wire.Bind(new(service.BookService), new(*service.BookServiceImpl)),
	wire.Bind(new(cache.Cache), new(*cache.CacheImpl)),
)

func InitializeHealthService(dbClient *ent.Client, cacheClient *redis.Client) *service.BookServiceImpl {
	wire.Build(providerSetHealth)
	return nil
}
