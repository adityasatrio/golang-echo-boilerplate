package repository

import (
	"context"
	"myapp/ent"
)

type UserRepository interface {
	Create(ctx context.Context, client *ent.Tx, user ent.User) (*ent.User, error)
	Update(ctx context.Context, client *ent.Tx, user ent.User, id uint64) (*ent.User, error)
	Delete(ctx context.Context, client *ent.Tx, id uint64) (*ent.User, error)

	SoftDelete(ctx context.Context, id uint64) (*ent.User, error)
	GetById(ctx context.Context, id uint64) (*ent.User, error)
	GetAll(ctx context.Context) ([]*ent.User, error)
}
