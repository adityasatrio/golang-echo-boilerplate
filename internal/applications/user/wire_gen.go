// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

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

// Injectors from user_injector.go:

func InitializedUserService(dbClient *ent.Client) *service.UserServiceImpl {
	userRepositoryImpl := repository.NewUserRepository(dbClient)
	roleRepositoryImpl := repository2.NewRoleRepository(dbClient)
	roleUserRepositoryImpl := repository3.NewRoleUserRepository(dbClient)
	trxServiceImpl := transaction.NewTrxService(dbClient)
	userServiceImpl := service.NewUserService(userRepositoryImpl, roleRepositoryImpl, roleUserRepositoryImpl, trxServiceImpl)
	return userServiceImpl
}

// user_injector.go:

var providerUser = wire.NewSet(repository.NewUserRepository, repository2.NewRoleRepository, repository3.NewRoleUserRepository, transaction.NewTrxService, service.NewUserService, wire.Bind(new(repository.UserRepository), new(*repository.UserRepositoryImpl)), wire.Bind(new(repository2.RoleRepository), new(*repository2.RoleRepositoryImpl)), wire.Bind(new(repository3.RoleUserRepository), new(*repository3.RoleUserRepositoryImpl)), wire.Bind(new(transaction.TrxService), new(*transaction.TrxServiceImpl)), wire.Bind(new(service.UserService), new(*service.UserServiceImpl)))
