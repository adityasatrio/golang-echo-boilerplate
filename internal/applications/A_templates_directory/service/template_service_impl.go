package service

import (
	"context"
	"myapp/ent"
	"myapp/internal/applications/A_templates_directory/dto"
	"myapp/internal/applications/A_templates_directory/repository/db"
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
	//TLDR; this service layer and functions will consist of your business logic.
	//The intention is to isolate all related logic with multiple domain activities and communication in service layer
	//if needed, transactional handling mandatory in this service layer

	//call repository and handle the result
	return &ent.Pet{}, nil
}
