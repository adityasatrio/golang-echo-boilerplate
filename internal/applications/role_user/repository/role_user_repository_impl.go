package repository

import (
	"context"
	"errors"
	"github.com/labstack/gommon/log"
	"myapp/ent"
	"myapp/ent/roleuser"
	"time"
)

type RoleUserRepositoryImpl struct {
	client *ent.Client
}

func NewRoleUserRepository(client *ent.Client) *RoleUserRepositoryImpl {
	return &RoleUserRepositoryImpl{
		client: client,
	}
}

func (r *RoleUserRepositoryImpl) GetByUserIdAndRoleId(ctx context.Context, userId uint64, roleId uint64) (*ent.RoleUser, error) {
	existingData, err := r.client.RoleUser.Query().
		Where(roleuser.UserID(userId), roleuser.RoleID(roleId)).
		Only(ctx)

	if err != nil {
		return nil, err
	}

	return existingData, nil

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

func (r *RoleUserRepositoryImpl) UpdateTx(ctx context.Context, clientTrx *ent.Client, updateRoleUser *ent.RoleUser) (*ent.RoleUser, error) {
	affected, err := clientTrx.RoleUser.
		Update().Where(roleuser.UserID(updateRoleUser.ID), roleuser.Version(updateRoleUser.Version)).
		SetUserID(updateRoleUser.UserID).
		SetRoleID(updateRoleUser.RoleID).
		SetUpdatedAt(time.Now()).
		SetUpdatedBy("user").
		Save(ctx)

	if err != nil {
		return nil, err
	}

	if affected < 1 {
		log.Errorf("ID %s no records were updated in database", updateRoleUser.ID)
		return nil, errors.New("no records were updated in database")
	}

	updatedData, err := clientTrx.RoleUser.Get(ctx, updateRoleUser.ID)
	if err != nil {
		return nil, err
	}

	return updatedData, nil
}
