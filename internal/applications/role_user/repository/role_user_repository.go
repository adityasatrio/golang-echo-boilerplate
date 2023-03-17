package repository

import (
	"context"
	"myapp/ent"
)

type RoleUserRepository interface {
	Create(ctx context.Context, client *ent.Client, request ent.RoleUser) (*ent.RoleUser, error)
	Update(ctx context.Context, client *ent.Client, request ent.RoleUser, id uint64) (*ent.RoleUser, error)
	DeleteByUserId(ctx context.Context, client *ent.Client, id uint64) (int, error)
}
