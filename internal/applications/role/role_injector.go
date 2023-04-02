//go:build wireinject
// +build wireinject

package role

import (
	"github.com/google/wire"
	"myapp/ent"
	"myapp/internal/applications/role/repository"
	"myapp/internal/applications/role/service"
	"myapp/internal/applications/transaction"
)

var provider = wire.NewSet(
	repository.NewRoleRepositoryImpl,
	service.NewRoleServiceImpl,
	transaction.NewTrxServiceImpl,
	wire.Bind(new(transaction.TrxService), new(*transaction.TrxServiceImpl)),
	wire.Bind(new(repository.RoleRepository), new(*repository.RoleRepositoryImpl)),
	wire.Bind(new(service.RoleService), new(*service.RoleServiceImpl)),
)

func InitializedRoleService(dbClient *ent.Client) *service.RoleServiceImpl {
	wire.Build(provider)
	return nil
}
