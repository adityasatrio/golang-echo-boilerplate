package cache

import (
	"context"
	"time"
)

type CachingService interface {
	Ping(ctx context.Context) error
	Create(ctx context.Context, key string, data interface{}, expiration time.Duration) (bool, error)
	Get(ctx context.Context, key string, data interface{}) (interface{}, error)
	Delete(ctx context.Context, key string) (bool, error)
}
