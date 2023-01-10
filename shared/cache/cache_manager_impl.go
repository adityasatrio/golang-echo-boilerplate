package cache

import (
	"context"
	"github.com/eko/gocache/lib/v4/cache"
	"github.com/eko/gocache/lib/v4/store"
	"log"
	"time"
)

type CacheManagerImpl struct {
	cacheConnection *cache.ChainCache[any]
}

func NewCacheManager(cacheConnection *cache.ChainCache[any]) *CacheManagerImpl {
	return &CacheManagerImpl{
		cacheConnection: cacheConnection,
	}
}

func (c *CacheManagerImpl) Set(ctx context.Context, key string, value *CacheValue, ttlInSecond int) {
	err := c.cacheConnection.Set(ctx, key, value, store.WithExpiration(5*time.Minute))
	if err != nil {
		log.Println("cache set failed", err)
	}
}

func (c *CacheManagerImpl) Get(ctx context.Context, key string) (*CacheValue, error) {
	value, err := c.cacheConnection.Get(ctx, key)
	if err != nil {
		log.Println("cache get failed", err)
		return nil, err
	}

	return value.(*CacheValue), nil
}
