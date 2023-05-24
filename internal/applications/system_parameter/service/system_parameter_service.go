package service

import (
	"context"
	"myapp/ent"
	"myapp/internal/applications/system_parameter/dto"
)

type (
	SystemParameterService interface {
		Create(ctx context.Context, create *dto.SystemParameterCreateRequest) (*ent.SystemParameter, error)
		Update(ctx context.Context, id int, update *dto.SystemParameterUpdateRequest) (*ent.SystemParameter, error)
		Delete(ctx context.Context, id int) (*ent.SystemParameter, error)
		SoftDelete(ctx context.Context, id int) (*ent.SystemParameter, error)
		GetById(ctx context.Context, id int) (*ent.SystemParameter, error)
		GetAll(ctx context.Context) ([]*ent.SystemParameter, error)
	}
)
