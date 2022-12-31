//go:build wireinject
// +build wireinject

package service

import (
	"myapp/config/database"
	"myapp/internal/applications/system_parameter/repository/db"
)
import "github.com/google/wire"

func IntializedUseCase() *systemParameterCase {
	wire.Build(db.NewSystemParameterRepository, database.NewEntClient())
	return nil
}
