package service

import "context"

type HealthService interface {
	Health(ctx context.Context, message string) (map[string]string, error)
	HealthDatabase(ctx context.Context, message string) (map[string]string, error)
	HealthCache(ctx context.Context, message string) (map[string]string, error)
}
