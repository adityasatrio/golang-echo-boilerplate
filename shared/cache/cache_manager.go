package cache

import (
	"context"
)

type CacheManager interface {
	Set(ctx context.Context, key string, value *CacheValue, ttlInSecond int)
	Get(ctx context.Context, key string) (*CacheValue, error)
}
