package db

import (
	"context"
	"myapp/ent"
)

type TemplateRepositoryImpl struct {
	dbClient *ent.Client
}

func NewTemplateRepository(dbClient *ent.Client) *TemplateRepositoryImpl {
	return &TemplateRepositoryImpl{
		dbClient: dbClient,
	}
}

func (r *TemplateRepositoryImpl) DatabaseAction(ctx context.Context, requestDomain *ent.Pet) (*ent.Pet, error) {
	//this repository layer and functions will consist of your interaction with source of data,
	//like CRUD operation on database. NO logic business in this layer
	//and NO transactional handling in this repository  layer

	return &ent.Pet{}, nil
}