//go:build wireinject
// +build wireinject

package user

import (
	"github.com/google/wire"
	"myapp/ent"
	roleRepository "myapp/internal/applications/role/repository"
	roleUserRepository "myapp/internal/applications/role_user/repository"
	"myapp/internal/applications/transaction"
	"myapp/internal/applications/user/repository"
	"myapp/internal/applications/user/service"
)

var providerUser = wire.NewSet(
	repository.NewUserRepositoryImpl,
	roleRepository.NewRoleRepositoryImpl,
	roleUserRepository.NewRoleUserRepositoryImpl,
	transaction.NewTrxServiceImpl,
	service.NewUserServiceImpl,

	wire.Bind(new(repository.UserRepository), new(*repository.UserRepositoryImpl)),
	wire.Bind(new(roleRepository.RoleRepository), new(*roleRepository.RoleRepositoryImpl)),
	wire.Bind(new(roleUserRepository.RoleUserRepository), new(*roleUserRepository.RoleUserRepositoryImpl)),
	wire.Bind(new(transaction.TrxService), new(*transaction.TrxServiceImpl)),
	wire.Bind(new(service.UserService), new(*service.UserServiceImpl)),
)

func InitializedUserService(dbClient *ent.Client) *service.UserServiceImpl {
	wire.Build(providerUser)
	return nil
}
