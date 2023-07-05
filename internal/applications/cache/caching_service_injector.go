//go:build wireinject
// +build wireinject

package cache

import (
	"github.com/go-redis/redis/v8"
	"github.com/google/wire"
)

var provider = wire.NewSet(
	NewCachingService,
	wire.Bind(new(CachingService), new(*CachingServiceImpl)),
)

func InitializedCachingServiceImpl(redisClient *redis.Client) *CachingServiceImpl {
	wire.Build(provider)
	return nil
}
