//go:build wireinject
// +build wireinject

package user

import (
	"github.com/google/wire"
	"myapp/ent"
	repository2 "myapp/internal/applications/role/repository"
	repository3 "myapp/internal/applications/role_user/repository"
	"myapp/internal/applications/transaction"
	"myapp/internal/applications/user/repository"
	"myapp/internal/applications/user/service"
)

var provider = wire.NewSet(
	repository.NewUserRepositoryImpl,
	repository2.NewRoleRepositoryImpl,
	repository3.NewRoleUserRepositoryImpl,
	service.NewUserServiceImpl,
	transaction.NewTxService,

	wire.Bind(new(repository.UserRepository), new(*repository.UserRepositoryImpl)),
	wire.Bind(new(repository2.RoleRepository), new(*repository2.RoleRepositoryImpl)),
	wire.Bind(new(repository3.RoleUserRepository), new(*repository3.RoleUserRepositoryImpl)),
	wire.Bind(new(service.UserService), new(*service.UserServiceImpl)),
)

func InitializedUserService(dbClient *ent.Client) *service.UserServiceImpl {
	wire.Build(provider)
	return nil
}
