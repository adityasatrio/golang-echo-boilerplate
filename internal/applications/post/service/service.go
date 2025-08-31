package service

import (
	"context"
	"myapp/ent"
	"myapp/internal/applications/post/dto"
)

type (
	PostService interface {
		LogicFunction(ctx context.Context, request *dto.ExampleRequest) (*ent.Pet, error)
	}
)
