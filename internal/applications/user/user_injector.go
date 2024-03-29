//go:build wireinject
// +build wireinject

package user

import (
	"github.com/go-redis/redis/v8"
	"github.com/google/wire"
	"myapp/ent"
	roleRepository "myapp/internal/applications/role/repository"
	roleUserRepository "myapp/internal/applications/role_user/repository"
	"myapp/internal/applications/user/repository"
	"myapp/internal/applications/user/service"
	"myapp/internal/component/cache"
	"myapp/internal/component/transaction"
)

var providerUser = wire.NewSet(
	repository.NewUserRepository,
	roleRepository.NewRoleRepository,
	roleUserRepository.NewRoleUserRepository,
	transaction.NewTrx,
	service.NewUserService,
	cache.NewCache,

	wire.Bind(new(repository.UserRepository), new(*repository.UserRepositoryImpl)),
	wire.Bind(new(roleRepository.RoleRepository), new(*roleRepository.RoleRepositoryImpl)),
	wire.Bind(new(roleUserRepository.RoleUserRepository), new(*roleUserRepository.RoleUserRepositoryImpl)),
	wire.Bind(new(transaction.Trx), new(*transaction.TrxImpl)),
	wire.Bind(new(cache.Cache), new(*cache.CacheImpl)),
	wire.Bind(new(service.UserService), new(*service.UserServiceImpl)),
)

func InitializedUserService(dbClient *ent.Client, redisClient *redis.Client) *service.UserServiceImpl {
	wire.Build(providerUser)
	return nil
}
