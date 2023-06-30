package outbound

import (
	"context"
	"myapp/internal/applications/quotes/dto"
)

type QuotesOutbound interface {
	GetQuotes(ctx context.Context, reqBody dto.QuoteApiRequest) (*dto.QuoteApiResponse, error)
	PostQuotes(ctx context.Context, reqBody dto.QuoteApiRequest) (*dto.QuoteApiResponse, error)
}
