package service

import (
	"context"
)

type HelloWorldService interface {
	Hello(ctx context.Context, message string, errorFlag string) (string, error)
}
