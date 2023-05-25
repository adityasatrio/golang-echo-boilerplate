package repository

import (
	"context"
	"myapp/ent"
)

type UserRepository interface {
	CreateTx(ctx context.Context, txClient *ent.Client, newUser ent.User) (*ent.User, error)
	UpdateTx(ctx context.Context, txClient *ent.Client, updateUser ent.User, id uint64) (*ent.User, error)

	Delete(ctx context.Context, id uint64) (*ent.User, error)
	SoftDelete(ctx context.Context, id uint64) (*ent.User, error)
	GetById(ctx context.Context, id uint64) (*ent.User, error)
	GetAll(ctx context.Context) ([]*ent.User, error)
}
