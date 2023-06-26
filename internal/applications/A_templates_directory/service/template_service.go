package service

import (
	"context"
	"myapp/ent"
	"myapp/internal/applications/A_templates_directory/dto"
)

type (
	TemplateService interface {
		LogicFunction(ctx context.Context, request *dto.ExampleRequest) (*ent.Pet, error)
	}
)
