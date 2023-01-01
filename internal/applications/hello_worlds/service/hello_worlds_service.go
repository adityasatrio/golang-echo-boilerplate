package service

import (
	"context"
)

type HelloWorldsService interface {
	Hello(ctx context.Context, message string, errorFlag string) (string, error)
}
