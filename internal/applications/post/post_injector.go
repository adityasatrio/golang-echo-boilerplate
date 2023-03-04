//go:build wireinject
// +build wireinject

package post

import (
	"github.com/google/wire"
	"myapp/ent"
	"myapp/internal/applications/post/repository/db"
	"myapp/internal/applications/post/service"
	"myapp/internal/applications/transaction"
)

var providerPost = wire.NewSet(
	db.NewPostRepositoryImpl,
	service.NewPostServiceImpl,
	transaction.NewTxService,
	wire.Bind(new(db.PostRepository), new(*db.PostRepositoryImpl)),
	wire.Bind(new(service.PostService), new(*service.PostServiceImpl)),
)

func InitializedPostInjector(dbClient *ent.Client) *service.PostServiceImpl {
	wire.Build(providerPost)
	return nil
}
