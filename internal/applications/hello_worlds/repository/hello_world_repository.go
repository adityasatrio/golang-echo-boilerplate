package repository

import (
	"context"
)

type HelloWorldsRepository interface {
	Hello(ctx context.Context, message string, errorFlag string) (string, error)
}
