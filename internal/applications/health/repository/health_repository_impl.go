package repository

import (
	"context"
	"github.com/labstack/gommon/log"
	"myapp/ent"
)

type HealthRepositoryImpl struct {
	client *ent.Client
}

func NewHealthRepository(dbConn *ent.Client) *HealthRepositoryImpl {
	return &HealthRepositoryImpl{
		client: dbConn,
	}
}

func (repo *HealthRepositoryImpl) Health(ctx context.Context, message string, queryFlag string) (map[string]string, error) {

	healthCheck := map[string]string{}
	if ctx != nil {
		log.Info("ctx debug", ctx)
		healthCheck["ctx_status"] = "UP"
		healthCheck["ctx_name"] = "echo"
	}

	if repo.client != nil {
		log.Info("client db debug", repo.client.Debug())
		healthCheck["db_status"] = "UP"
		healthCheck["db_name"] = "mysql"
	}

	finalMsg := message + "hello from repository layer " + "hello from query parameter=" + queryFlag
	healthCheck["final_msg"] = finalMsg

	return healthCheck, nil
}
