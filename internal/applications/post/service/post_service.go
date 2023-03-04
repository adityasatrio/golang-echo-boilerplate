package service

import (
	"context"
	"myapp/ent"
	"myapp/internal/applications/post/dto"
)

type PostService interface {
	Create(ctx context.Context, request dto.PostRequest) (*ent.Post, error)
	Update(ctx context.Context, request dto.PostRequest, id int) (*ent.Post, error)
	SoftDelete(ctx context.Context, id int) (*ent.Post, error)
	Delete(ctx context.Context, id int) (*ent.Post, error)
	GetById(ctx context.Context, id int) (*ent.Post, error)
	GetAll(ctx context.Context) ([]*ent.Post, error)
}
