package service

import (
	"context"
	"myapp/ent"
	"myapp/internal/applications/system_parameter/dto"
)

type (
	SystemParameterService interface {
		Create(ctx context.Context, create *dto.SystemParameterCreateRequest) (*ent.System_parameter, error)
		Update(ctx context.Context, id int, update *dto.SystemParameterUpdateRequest) (*ent.System_parameter, error)
		Delete(ctx context.Context, id int) (*ent.System_parameter, error)
		GetById(ctx context.Context, id int) (*ent.System_parameter, error)
		GetAll(ctx context.Context) ([]*ent.System_parameter, error)
	}
)
