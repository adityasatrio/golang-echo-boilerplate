//go:build wireinject
// +build wireinject

package user

import (
	"github.com/go-redis/redis/v8"
	"github.com/google/wire"
	"myapp/ent"
	"myapp/internal/applications/cache"
	roleRepository "myapp/internal/applications/role/repository"
	roleUserRepository "myapp/internal/applications/role_user/repository"
	"myapp/internal/applications/transaction"
	"myapp/internal/applications/user/repository"
	"myapp/internal/applications/user/service"
)

var providerUser = wire.NewSet(
	repository.NewUserRepository,
	roleRepository.NewRoleRepository,
	roleUserRepository.NewRoleUserRepository,
	transaction.NewTrxService,
	service.NewUserService,
	cache.NewCachingService,

	wire.Bind(new(repository.UserRepository), new(*repository.UserRepositoryImpl)),
	wire.Bind(new(roleRepository.RoleRepository), new(*roleRepository.RoleRepositoryImpl)),
	wire.Bind(new(roleUserRepository.RoleUserRepository), new(*roleUserRepository.RoleUserRepositoryImpl)),
	wire.Bind(new(transaction.TrxService), new(*transaction.TrxServiceImpl)),
	wire.Bind(new(cache.CachingService), new(*cache.CachingServiceImpl)),
	wire.Bind(new(service.UserService), new(*service.UserServiceImpl)),
)

func InitializedUserService(dbClient *ent.Client, redisClient *redis.Client) *service.UserServiceImpl {
	wire.Build(providerUser)
	return nil
}
