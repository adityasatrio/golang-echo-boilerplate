//go:build wireinject
// +build wireinject

package user

import (
	"github.com/google/wire"
	"myapp/ent"
	"myapp/internal/applications/transaction"
	"myapp/internal/applications/user/repository"
	"myapp/internal/applications/user/service"
)

var provider = wire.NewSet(
	repository.NewUserRepositoryImpl,
	service.NewUserServiceImpl,
	transaction.NewTxService,
	wire.Bind(new(repository.UserRepository), new(*repository.UserRepositoryImpl)),
	wire.Bind(new(service.UserService), new(*service.UserServiceImpl)),
)

func InitializedRoleService(dbClient *ent.Client) *service.UserServiceImpl {
	wire.Build(provider)
	return nil
}
