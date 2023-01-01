package service

import (
	"context"
	"errors"
	"github.com/labstack/gommon/log"
	"myapp/exceptions"
	"myapp/internal/applications/hello_worlds/repository"
	"strings"
)

type HelloWorldsServiceImpl struct {
	repository repository.HelloWorldsRepository
}

func NewHelloWorldsService(repository repository.HelloWorldsRepository) *HelloWorldsServiceImpl {
	return &HelloWorldsServiceImpl{
		repository: repository,
	}
}

func (s *HelloWorldsServiceImpl) Hello(ctx context.Context, message string, errorFlag string) (string, error) {
	log.Info("ctx debug", ctx)

	messageService := message + " hello from s-impl layer -"
	if strings.EqualFold(errorFlag, "service") {
		return "", exceptions.NewBusinessLogicError(exceptions.EBL10007, errors.New(messageService))
	}
	result, err := s.repository.Hello(ctx, messageService, errorFlag)
	return result, err
}
