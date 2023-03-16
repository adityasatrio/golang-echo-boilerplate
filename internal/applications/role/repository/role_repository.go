package repository

import (
	"context"
	"myapp/ent"
	"myapp/internal/applications/role/dto"
)

type RoleRepository interface {
	Create(ctx context.Context, request dto.RoleRequest) (*ent.Role, error)
	Update(ctx context.Context, request dto.RoleRequest, id uint64) (*ent.Role, error)
	Delete(ctx context.Context, id uint64) (*ent.Role, error)
	SoftDelete(ctx context.Context, id uint64) (*ent.Role, error)
	GetById(ctx context.Context, id uint64) (*ent.Role, error)
	GetAll(ctx context.Context) ([]*ent.Role, error)
}
