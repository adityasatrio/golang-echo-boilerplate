package cache

import (
	"context"
)

type CacheManager interface {
	Set(ctx context.Context, key string, value any, ttlInSecond int)
	Get(ctx context.Context, key string) (any, error)
}
