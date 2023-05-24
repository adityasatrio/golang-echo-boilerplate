package repository

import (
	"context"
	"myapp/ent"
	"time"
)

type RoleUserRepositoryImpl struct {
}

func NewRoleUserRepositoryImpl() *RoleUserRepositoryImpl {
	return &RoleUserRepositoryImpl{}
}

func (r *RoleUserRepositoryImpl) CreateTx(ctx context.Context, clientTrx *ent.Client, request ent.RoleUser) (*ent.RoleUser, error) {
	response, err := clientTrx.RoleUser.Create().
		SetUserID(request.UserID).
		SetRoleID(request.RoleID).
		SetCreatedBy("user").
		SetCreatedAt(time.Now()).
		Save(ctx)

	if err != nil {
		return nil, err
	}

	return response, nil
}

func (r *RoleUserRepositoryImpl) Update(ctx context.Context, clientTrx *ent.Client, request ent.RoleUser, id uint64) (*ent.RoleUser, error) {
	response, err := clientTrx.RoleUser.UpdateOneID(request.ID).
		SetUserID(request.UserID).
		SetRoleID(request.RoleID).
		SetUpdatedAt(time.Now()).
		SetUpdatedBy("user").
		Save(ctx)

	if err != nil {
		return nil, err
	}

	return response, nil
}
