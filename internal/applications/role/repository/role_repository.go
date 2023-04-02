package repository

import (
	"context"
	"myapp/ent"
)

type RoleRepository interface {
	Create(ctx context.Context, role ent.Role) (*ent.Role, error)
	Update(ctx context.Context, role ent.Role, id uint64) (*ent.Role, error)
	Delete(ctx context.Context, id uint64) (*ent.Role, error)
	SoftDelete(ctx context.Context, id uint64) (*ent.Role, error)
	GetById(ctx context.Context, id uint64) (*ent.Role, error)
	GetAll(ctx context.Context) ([]*ent.Role, error)
}
