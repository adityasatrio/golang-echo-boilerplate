package db

import (
	"context"
	"myapp/ent"
)

type PostRepositoryImpl struct {
	dbClient *ent.Client
}

func NewPostRepository(dbClient *ent.Client) *PostRepositoryImpl {
	return &PostRepositoryImpl{
		dbClient: dbClient,
	}
}

func (r *PostRepositoryImpl) DatabaseAction(ctx context.Context, requestDomain *ent.Pet) (*ent.Pet, error) {
	//TLDR; this repository layer and functions will consist of your interaction with source of domain data, such as CRUD operation on database.
	//NO logic business in this layer and NO transactional handling in this repository  layer

	return &ent.Pet{}, nil
}
