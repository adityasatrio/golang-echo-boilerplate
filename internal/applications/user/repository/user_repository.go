package repository

import (
	"context"
	"myapp/ent"
	"myapp/internal/applications/user/dto"
)

type UserRepository interface {
	Create(ctx context.Context, request dto.UserRequest) (*ent.User, error)
	Update(ctx context.Context, request dto.UserRequest, id uint64) (*ent.User, error)
	Delete(ctx context.Context, id uint64) (*ent.User, error)
	SoftDelete(ctx context.Context, id uint64) (*ent.User, error)
	GetById(ctx context.Context, id uint64) (*ent.User, error)
	GetAll(ctx context.Context) ([]*ent.User, error)
}
