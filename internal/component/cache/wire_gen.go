// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package cache

import (
	"github.com/go-redis/redis/v8"
	"github.com/google/wire"
)

// Injectors from cache_injector.go:

func InitializedCachingImpl(redisClient *redis.Client) *CacheImpl {
	cachingImpl := NewCache(redisClient)
	return cachingImpl
}

// cache_injector.go:

var provider = wire.NewSet(
	NewCache, wire.Bind(new(Cache), new(*CacheImpl)),
)
