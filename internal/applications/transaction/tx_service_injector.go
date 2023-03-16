//go:build wireinject
// +build wireinject

package transaction

import (
	"github.com/google/wire"
	"myapp/ent"
)

func InitializedTxService(dbClient *ent.Client) *TxService {
	wire.Build(NewTxService)
	return nil
}
