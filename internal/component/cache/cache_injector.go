//go:build wireinject
// +build wireinject

package cache

import (
	"github.com/go-redis/redis/v8"
	"github.com/google/wire"
)

var provider = wire.NewSet(
	NewCache,
	wire.Bind(new(Cache), new(*CacheImpl)),
)

func InitializedCachingImpl(redisClient *redis.Client) *CacheImpl {
	wire.Build(provider)
	return nil
}
