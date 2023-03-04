package db

import (
	"context"
	"myapp/ent"
	"myapp/ent/post"
	"myapp/internal/applications/post/dto"
	"time"
)

type PostRepositoryImpl struct {
	client *ent.Client
}

func NewPostRepositoryImpl(client *ent.Client) *PostRepositoryImpl {
	return &PostRepositoryImpl{client: client}
}

func (r *PostRepositoryImpl) Create(ctx context.Context, request dto.PostRequest) (*ent.Post, error) {
	response, err := r.client.Post.Create().
		SetTitle(request.Title).
		SetContent(request.Content).
		SetSlug(request.Slug).
		SetStatus(request.Status).
		Save(ctx)

	if err != nil {
		return nil, err
	}

	return response, nil
}

func (r *PostRepositoryImpl) Update(ctx context.Context, request dto.PostRequest, id int) (*ent.Post, error) {
	saved, err := r.client.Post.
		UpdateOneID(id).
		SetTitle(request.Title).
		SetContent(request.Content).
		SetSlug(request.Slug).
		Save(ctx)

	if err != nil {
		return nil, err
	}

	return saved, nil
}

func (r *PostRepositoryImpl) SoftDelete(ctx context.Context, id int, client *ent.Client) (*ent.Post, error) {
	deleted, err := client.Post.
		UpdateOneID(id).
		SetDeletedAt(time.Now()).
		Save(ctx)

	if err != nil {
		return nil, err
	}

	return deleted, nil
}

func (r *PostRepositoryImpl) Delete(ctx context.Context, id int) (*ent.Post, error) {
	err := r.client.Post.DeleteOneID(id).Exec(ctx)

	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (r *PostRepositoryImpl) GetById(ctx context.Context, id int) (*ent.Post, error) {
	data, err := r.client.Post.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	return data, err
}

func (r *PostRepositoryImpl) GetAll(ctx context.Context) ([]*ent.Post, error) {
	data, err := r.client.Post.Query().
		Where(post.DeletedAtIsNil()).
		All(ctx)
	if err != nil {
		return nil, err
	}

	return data, nil
}
