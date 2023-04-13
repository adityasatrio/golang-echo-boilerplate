//go:build wireinject
// +build wireinject

package transaction

import (
	"github.com/google/wire"
	"myapp/ent"
)

var provider = wire.NewSet(
	NewTrxServiceImpl,
	wire.Bind(new(TrxService), new(*TrxServiceImpl)),
)

func InitializedTxService(dbClient *ent.Client) *TrxServiceImpl {
	wire.Build(provider)
	return nil
}
