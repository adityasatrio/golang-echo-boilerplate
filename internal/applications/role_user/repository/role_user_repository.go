package repository

import (
	"context"
	"myapp/ent"
)

type RoleUserRepository interface {
	GetByUserIdAndRoleId(ctx context.Context, userId uint64, roleId uint64) (*ent.RoleUser, error)
	CreateTx(ctx context.Context, client *ent.Client, request ent.RoleUser) (*ent.RoleUser, error)
	UpdateTx(ctx context.Context, client *ent.Client, request *ent.RoleUser) (*ent.RoleUser, error)
}
