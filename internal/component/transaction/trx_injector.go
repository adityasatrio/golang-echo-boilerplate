//go:build wireinject
// +build wireinject

package transaction

import (
	"github.com/google/wire"
	"myapp/ent"
)

var provider = wire.NewSet(
	NewTrx,
	wire.Bind(new(Trx), new(*TrxImpl)),
)

func InitializedTxService(dbClient *ent.Client) *TrxImpl {
	wire.Build(provider)
	return nil
}
