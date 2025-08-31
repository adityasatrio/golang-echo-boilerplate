package service

import (
	"context"
	"myapp/ent"
	"myapp/internal/applications/book/dto"
)

type (
	BookService interface {
		LogicFunction(ctx context.Context, request *dto.ExampleRequest) (*ent.Pet, error)
	}
)
