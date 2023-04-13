package repository

import (
	"context"
	"myapp/ent"
	"myapp/ent/roleuser"
	"time"
)

type RoleUserRepositoryImpl struct {
	client *ent.Client
}

func NewRoleUserRepositoryImpl(client *ent.Client) *RoleUserRepositoryImpl {
	return &RoleUserRepositoryImpl{client: client}
}

func (r *RoleUserRepositoryImpl) Create(ctx context.Context, client *ent.Tx, request ent.RoleUser) (*ent.RoleUser, error) {
	response, err := client.RoleUser.Create().
		SetUserID(request.UserID).
		SetRoleID(request.RoleID).
		SetCreatedAt(time.Now()).
		Save(ctx)

	if err != nil {
		return nil, err
	}

	return response, nil
}

func (r *RoleUserRepositoryImpl) Update(ctx context.Context, client *ent.Tx, request ent.RoleUser, id uint64) (*ent.RoleUser, error) {

	//delete existing role user:
	_, err := client.RoleUser.Delete().Where(roleuser.UserID(id)).Exec(ctx)

	if err != nil {
		return nil, err
	}

	//create new role user:
	response, err := client.RoleUser.Create().
		SetUserID(request.UserID).
		SetRoleID(request.RoleID).
		SetCreatedAt(time.Now()).
		SetUpdatedAt(time.Now()).
		Save(ctx)

	if err != nil {
		return nil, err
	}

	return response, nil
}
