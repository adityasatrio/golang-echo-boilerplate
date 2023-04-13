package repository

import (
	"context"
	"myapp/ent"
)

type RoleUserRepository interface {
	Create(ctx context.Context, client *ent.Tx, request ent.RoleUser) (*ent.RoleUser, error)
	Update(ctx context.Context, client *ent.Tx, request ent.RoleUser, id uint64) (*ent.RoleUser, error)
}
