package repository

import (
	"context"
	"fmt"
	"github.com/labstack/gommon/log"
	"myapp/ent"
	"strings"
)

type HelloWorldsRepositoryImpl struct {
	client *ent.Client
}

func NewHelloWorldsRepository(dbConn *ent.Client) *HelloWorldsRepositoryImpl {
	return &HelloWorldsRepositoryImpl{
		client: dbConn,
	}
}

func (repo *HelloWorldsRepositoryImpl) Hello(ctx context.Context, message string, errorFlag string) (string, error) {

	log.Info("ctx debug", ctx)
	log.Info("client db debug", repo.client.Debug())

	if repo.client == nil {
		err := fmt.Errorf("client db is null")
		return "", err
	}

	if strings.EqualFold(errorFlag, "repository") {
		return "", fmt.Errorf("error from repository-impl layer")
	}

	messageRepository := message + " hello from repository-impl layer "
	trimmedMessageRepository := strings.Trim(messageRepository, " ")

	return trimmedMessageRepository, nil
}
