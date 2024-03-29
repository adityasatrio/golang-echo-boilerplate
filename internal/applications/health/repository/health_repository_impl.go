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

func (r *HealthRepositoryImpl) Health(ctx context.Context, message string) (map[string]string, error) {

	healthCheck := map[string]string{}
	if ctx != nil {
		log.Info("ctx debug", ctx)
		healthCheck["ctx_status"] = "UP"
		healthCheck["ctx_name"] = "echo"
	}

	finalMsg := message + "hello from repository layer " + "hello from query parameter"
	healthCheck["final_msg"] = finalMsg

	_, err := r.client.ExecContext(ctx, "SELECT 1")
	log.Info("client db debug", r.client.Debug())

	if err != nil {
		healthCheck["db_status"] = "DOWN"
		healthCheck["db_name"] = "mysql"
	} else {
		healthCheck["db_status"] = "UP"
		healthCheck["db_name"] = "mysql"
	}

	return healthCheck, err
}
