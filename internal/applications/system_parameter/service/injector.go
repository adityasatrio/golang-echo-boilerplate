//go:build wireinject
// +build wireinject

package service

import (
	"myapp/config/database"
	"myapp/internal/applications/system_parameter/repository"
)
import "github.com/google/wire"

func IntializedUseCase() *systemParameterCase {
	wire.Build(repository.NewSystemParameterRepository, database.NewEntClient())
	return nil
}
