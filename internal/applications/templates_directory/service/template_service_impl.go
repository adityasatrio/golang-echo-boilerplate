package service

import (
	"context"
	"myapp/ent"
	"myapp/internal/applications/templates_directory/dto"
	"myapp/internal/applications/templates_directory/repository/db"
)

type TemplateServiceImpl struct {
	repository db.TemplateRepository
}

func NewTemplateService(repository db.TemplateRepository) *TemplateServiceImpl {
	return &TemplateServiceImpl{
		repository: repository,
	}
}

func (s *TemplateServiceImpl) LogicFunction(ctx context.Context, request *dto.ExampleRequest) (*ent.Pet, error) {
	//this service layer and functions will consist of your business logic,
	//logic related with related multiple domain activities and communication
	//if needed, transactional handling mandatory in this service layer

	//call repository and handle the result
	return &ent.Pet{}, nil
}
