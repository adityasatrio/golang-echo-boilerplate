package service

import (
	"context"
	"myapp/internal/applications/quotes/dto"
)

type QuotesService interface {
	GetQuotes(ctx context.Context, queryParameter map[string]string) (*dto.QuoteApiResponse, error)
	PostQuotes(ctx context.Context, reqBody *dto.QuoteRequest) (*dto.QuoteApiResponse, error)
}
