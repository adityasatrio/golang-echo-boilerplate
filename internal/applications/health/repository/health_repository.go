package repository

import (
	"context"
)

type HealthRepository interface {
	Health(ctx context.Context, message string) (map[string]string, error)
}
