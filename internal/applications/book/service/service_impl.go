package service

import (
	"context"
	"myapp/ent"
	"myapp/internal/applications/book/dto"
	"myapp/internal/applications/book/repository/db"
)

type BookServiceImpl struct {
	repository db.BookRepository
}

func NewBookService(repository db.BookRepository) *BookServiceImpl {
	return &BookServiceImpl{
		repository: repository,
	}
}

func (s *BookServiceImpl) LogicFunction(ctx context.Context, request *dto.ExampleRequest) (*ent.Pet, error) {
	//TLDR; this service layer and functions will consist of your business logic.
	//The intention is to isolate all related logic with multiple domain activities and communication in service layer
	//if needed, transactional handling mandatory in this service layer

	//call repository and handle the result
	return &ent.Pet{}, nil
}
