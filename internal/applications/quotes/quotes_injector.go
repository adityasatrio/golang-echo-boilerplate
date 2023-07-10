//go:build wireinject
// +build wireinject

package system_parameter

import (
	"github.com/google/wire"
	"myapp/internal/applications/quotes/repository/outbound"
	"myapp/internal/applications/quotes/service"
)

var providerSetQuotes = wire.NewSet(
	outbound.NewQuoteOutbound,
	service.NewQuotesService,
	wire.Bind(new(outbound.QuotesOutbound), new(*outbound.QuoteOutboundImpl)),
	wire.Bind(new(service.QuotesService), new(*service.QuotesServiceImpl)),
)

func InitializedQuotesService() *service.QuotesServiceImpl {
	wire.Build(providerSetQuotes)
	return nil
}
