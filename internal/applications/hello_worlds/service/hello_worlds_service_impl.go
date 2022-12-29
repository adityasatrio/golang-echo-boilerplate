package service

import (
	"context"
	"fmt"
	"github.com/labstack/gommon/log"
	"myapp/internal/applications/hello_worlds/repository"
	"strings"
)

type HelloWorldsServiceImpl struct {
	repository repository.HelloWorldsRepository
}

func NewHelloWorldsService(repository repository.HelloWorldsRepository) HelloWorldService {
	return &HelloWorldsServiceImpl{
		repository: repository,
	}
}

func (service *HelloWorldsServiceImpl) Hello(ctx context.Context, message string, errorFlag string) (string, error) {
	log.Info("ctx debug", ctx)

	if strings.EqualFold(errorFlag, "service") {
		return "", fmt.Errorf("error from service-impl layer")
	}

	messageService := message + " hello from service-impl layer -"
	result, err := service.repository.Hello(ctx, messageService, errorFlag)
	return result, err
}
